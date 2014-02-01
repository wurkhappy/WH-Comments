package models

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"github.com/wurkhappy/WH-Comments/DB"
)

type Tag struct {
	ID          string `json:"id"`
	AgreementID string `json:"agreementID"`
	Name        string `json:"name"`
}

func NewTag() *Tag {
	id, _ := uuid.NewV4()
	return &Tag{
		ID: id.String(),
	}
}

func (t *Tag) Save() (err error) {
	jsonByte, _ := json.Marshal(t)
	r, err := DB.SaveTag.Query(t.ID, string(jsonByte))
	r.Close()
	if err != nil {
		return err
	}
	return nil
}

func FindTagsByAgreementID(id string) (tags []*Tag, err error) {
	r, err := DB.FindTagsByAgreementID.Query(id)
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
		var t *Tag
		json.Unmarshal([]byte(s), &t)
		tags = append(tags, t)
	}
	return tags, nil
}
