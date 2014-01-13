package models

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	rbtmq "github.com/wurkhappy/Rabbitmq-go-wrapper"
	"github.com/wurkhappy/WH-Comments/DB"
	"github.com/wurkhappy/WH-Config"
	"log"
	"time"
)

type Comment struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"userID"`
	AgreementID        string    `json:"agreementID"`
	AgreementVersionID string    `json:"agreementVersionID"`
	DateCreated        time.Time `json:"dateCreated"`
	Text               string    `json:"text"`
	Tags               []*Tag    `json:"tags"`
}

func NewComment() *Comment {
	id, _ := uuid.NewV4()
	return &Comment{
		ID:          id.String(),
		DateCreated: time.Now(),
	}
}

func (c *Comment) Save() (err error) {
	jsonByte, _ := json.Marshal(c)
	_, err = DB.UpsertComment.Query(c.ID, string(jsonByte))
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func FindCommentsByAgreementID(id string) (comments []*Comment, err error) {
	r, err := DB.FindCommentsByAgreementID.Query(id)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var s string
		err = r.Scan(&s)
		if err != nil {
			return nil, err
		}
		var c *Comment
		json.Unmarshal([]byte(s), &c)
		comments = append(comments, c)
	}
	return comments, nil
}

func FindCommentsByVersionID(id string) (comments []*Comment, err error) {
	r, err := DB.FindCommentsByVersionID.Query(id)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var s string
		err = r.Scan(&s)
		if err != nil {
			return nil, err
		}
		var c *Comment
		json.Unmarshal([]byte(s), &c)
		comments = append(comments, c)
	}
	return comments, nil
}

func SendCommentEmail(c *Comment) {
	payload := map[string]interface{}{
		"Body": map[string]interface{}{
			"comment": c,
		},
	}

	body, _ := json.Marshal(payload)
	publisher, err := rbtmq.NewPublisher(connection, config.EmailExchange, "direct", config.EmailQueue, "/comment")
	if err != nil {
		dialRMQ()
		publisher, _ = rbtmq.NewPublisher(connection, config.EmailExchange, "direct", config.EmailQueue, "/comment")
	}
	publisher.Publish(body, true)
}

func (c *Comment) CreateNewTags() {
	for _, tag := range c.Tags {
		if tag.ID == "" {
			t := NewTag()
			tag.ID = t.ID
			tag.AgreementID = c.AgreementID
			tag.Save()
		}
	}
}
