package handlers

import (
	"encoding/json"
	"github.com/wurkhappy/WH-Comments/models"
	"net/http"
)

func GetTags(params map[string]interface{}, body []byte) ([]byte, error, int) {
	id := params["agreementID"].(string)

	tags, _ := models.FindTagsByAgreementID(id)

	t, _ := json.Marshal(tags)
	return t, nil, http.StatusOK
}
