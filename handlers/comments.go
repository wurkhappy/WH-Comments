package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wurkhappy/WH-Comments/models"
	"net/http"
)

func CreateComment(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	agreementID := vars["agreementID"]

	comment := models.NewComment()
	date := comment.DateCreated

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	reqBytes := buf.Bytes()
	json.Unmarshal(reqBytes, &comment)
	comment.AgreementID = agreementID
	comment.DateCreated = date
	_ = comment.Save()
	go models.SendCommentEmail(comment)

	a, _ := json.Marshal(comment)
	w.Write(a)
}

func GetComments(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["agreementID"]

	params := req.URL.Query()
	var version string
	version = params.Get("version")
	var comments []*models.Comment
	if version != "" {
		comments, _ = models.FindCommentsByVersionID(version)
	} else {
		comments, _ = models.FindCommentsByAgreementID(id)
	}

	u, _ := json.Marshal(comments)
	w.Write(u)

}
