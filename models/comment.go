package models

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"github.com/streadway/amqp"
	rbtmq "github.com/wurkhappy/Rabbitmq-go-wrapper"
	"github.com/wurkhappy/WH-Comments/DB"
	"time"
	"log"
)

type Comment struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"userID"`
	AgreementID        string    `json:"agreementID"`
	AgreementVersionID string    `json:"agreementVersionID"`
	DateCreated        time.Time `json:"dateCreated"`
	Text               string    `json:"text"`
	MilestoneID        string    `json:"milestoneID"`
	StatusID           string    `json:"statusID"`
}

var connection *amqp.Connection

func init() {
	uri := "amqp://guest:guest@localhost:5672/"
	cn, err := amqp.Dial(uri)
	if err != nil {
		panic(err)
	}
	connection = cn
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
	publisher, _ := rbtmq.NewPublisher(connection, "email", "direct", "email", "/comment")
	publisher.Publish(body, true)
}
