package models

import (
	"github.com/nu7hatch/gouuid"
	"testing"
)

func test_SaveComment(t *testing.T) {
	com := NewComment()
	err := com.Save()

	if err != nil {
		t.Errorf("%s--- error is:%v", "test_SaveComment", err)
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
		t.Errorf("%s--- all comments were not found", "test_FindCommentsByVersionID")
	}
}

func test_createNewTags(t *testing.T) {
	com1 := NewComment()

	agreementID, _ := uuid.NewV4()
	com1.AgreementID = agreementID.String()

	tag1 := new(Tag)
	tag1.Name = "test tag"
	tag2 := new(Tag)
	tag2.Name = "test tag"

	tags := []*Tag{tag1, tag2}

	com1.Tags = tags
	com1.CreateNewTags()

	for _, tag := range com1.Tags {
		if tag.ID == "" {
			t.Error("tag id wasn't set")
		}
		if tag.AgreementID != com1.AgreementID {
			t.Error("tag agreement id wasn't set")
		}
	}
}
