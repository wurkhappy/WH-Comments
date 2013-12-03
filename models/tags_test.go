package models

import (
	"github.com/nu7hatch/gouuid"
	"testing"
)

func test_NewTag(t *testing.T) {
	tag := NewTag()
	if tag.ID == "" {
		t.Error("tag ID was not created")
	}
}

func test_SaveTag(t *testing.T) {
	tag := NewTag()
	err := tag.Save()
	if err != nil {
		t.Errorf("%s--- error is:%v", "test_SaveTag", err)
	}
}

func test_FindTagsByAgreementID(t *testing.T) {
	tag1 := NewTag()
	tag2 := NewTag()

	agreementID, _ := uuid.NewV4()
	tag1.AgreementID = agreementID.String()
	tag2.AgreementID = agreementID.String()

	tag1.Save()
	tag2.Save()

	tags, _ := FindTagsByAgreementID(tag1.AgreementID)
	if len(tags) != 2 {
		t.Errorf("%s--- all comments were not found", "test_FindTagsByAgreementID")
	}
}
