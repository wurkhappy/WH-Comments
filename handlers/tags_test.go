package handlers

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"github.com/wurkhappy/WH-Comments/models"
	"testing"
)

func createTag(tagName, agreementID string) *models.Tag {
	tag := models.NewTag()
	tag.Name = tagName
	tag.AgreementID = agreementID
	tag.Save()
	return tag
}

func test_GetTags(t *testing.T) {
	id, _ := uuid.NewV4()
	tagName1 := "new tag1"
	agreementID := id.String()
	_ = createTag(tagName1, agreementID)

	tagName2 := "new tag2"
	_ = createTag(tagName2, agreementID)

	params := map[string]interface{}{
		"agreementID": agreementID,
	}
	body, err, statusCode := GetTags(params, nil)
	if err != nil {
		t.Error("error incorrectly returned")
	}
	if statusCode >= 400 {
		t.Error("wrong status code returned")
	}
	var tags []map[string]interface{}
	json.Unmarshal(body, &tags)
	if len(tags) != 2 {
		t.Error("wrong number of tags returned")
	}
}
