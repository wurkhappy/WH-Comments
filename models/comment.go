package models

import (
	"github.com/nu7hatch/gouuid"
	"github.com/wurkhappy/WH-Comments/DB"
	"labix.org/v2/mgo/bson"
	"time"
)

type Comment struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userID"`
	AgreementID string    `json:"agreementID"`
	DateCreated time.Time `json:"dateCreated"`
	Text        string    `json:"text"`
	MilestoneID string    `json:"milestoneID"`
	StatusID    string    `json:"statusID"`
}

func FindCommentsByAgreementID(id string, ctx *DB.Context) (comments []*Comment, err error) {
	err = ctx.Database.C("comments").Find(bson.M{"agreementid": id}).Sort("-lastmodified").All(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *Comment) SaveWithCtx(ctx *DB.Context) (err error) {
	if c.ID == "" {
		commentid, _ := uuid.NewV4()
		c.ID = commentid.String()
	}
	c.DateCreated = time.Now()

	coll := ctx.Database.C("comments")
	if _, err := coll.UpsertId(c.ID, &c); err != nil {
		return err
	}
	return nil
}
