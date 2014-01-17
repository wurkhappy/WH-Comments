package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/wurkhappy/WH-Comments/models"
	"html/template"
	"net/http"
)

func CreateComment(params map[string]interface{}, body []byte) ([]byte, error, int) {
	agreementID := params["agreementID"].(string)

	comment := models.NewComment()
	date := comment.DateCreated

	json.Unmarshal(body, &comment)

	comment.AgreementID = agreementID
	comment.DateCreated = date
	comment.CreateNewTags()
	comment.Text = template.HTMLEscapeString(comment.Text)

	err := comment.Save()
	if err != nil {
		return nil, fmt.Errorf("%s %s", "Error saving: ", err.Error()), http.StatusBadRequest
	}
	sendEmail := true
	if createEmail, ok := params["sendEmail"]; ok {
		if createEmail.([]string)[0] == "false" {
			sendEmail = false
		}
	}
	if sendEmail {
		go models.SendCommentEmail(comment)
	}

	c, _ := json.Marshal(comment)
	return c, nil, http.StatusOK
}

func GetComments(params map[string]interface{}, body []byte) ([]byte, error, int) {
	id := params["agreementID"].(string)

	var version string
	if versions, ok := params["version"]; ok {
		version = versions.([]string)[0]
	}
	var comments []*models.Comment
	if version != "" {
		comments, _ = models.FindCommentsByVersionID(version)
	} else {
		comments, _ = models.FindCommentsByAgreementID(id)
	}

	c, _ := json.Marshal(comments)
	return c, nil, http.StatusOK
}
