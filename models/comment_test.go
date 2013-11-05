package models

import (
	// "encoding/json"
	"github.com/wurkhappy/WH-Comments/DB"
	// "labix.org/v2/mgo"
	"github.com/nu7hatch/gouuid"
	"testing"
	// "log"
)

func init() {
	DB.Name = "testdb"
	DB.Setup()
	DB.CreateStatements()
}

func TestIntegrationTests(t *testing.T) {
	if !testing.Short() {
		test_SaveComment(t)
		test_FindCommentsByAgreementID(t)
		test_FindCommentsByVersionID(t)

		DB.DB.Exec("DELETE from comment")
	}
}

func test_SaveComment(t *testing.T) {
	com := NewComment()
	err := com.Save()

	if err != nil {
		t.Errorf("%s--- error is:%v", "testSaveUser", err)
	}
}

func test_FindCommentsByAgreementID(t *testing.T) {
	com1 := NewComment()
	com2 := NewComment()

	agreementID, _ := uuid.NewV4()
	com1.AgreementID = agreementID.String()
	com2.AgreementID = agreementID.String()

	com1.Save()
	com2.Save()

	comments, _ := FindCommentsByAgreementID(com1.AgreementID)
	if len(comments) != 2 {
		t.Errorf("%s--- all comments were not found", "test_FindCommentsByAgreementID")
	}
}

func test_FindCommentsByVersionID(t *testing.T) {
	com1 := NewComment()
	com2 := NewComment()

	versionID, _ := uuid.NewV4()
	com1.AgreementVersionID = versionID.String()
	com2.AgreementVersionID = versionID.String()

	com1.Save()
	com2.Save()

	comments, _ := FindCommentsByVersionID(com1.AgreementVersionID)
	if len(comments) != 2 {
		t.Errorf("%s--- all comments were not found", "test_FindCommentsByAgreementID")
	}
}
