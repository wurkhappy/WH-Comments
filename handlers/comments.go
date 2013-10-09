package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wurkhappy/WH-Comments/DB"
	"github.com/wurkhappy/WH-Comments/models"
	// "log"
	"net/http"
)

func CreateComment(w http.ResponseWriter, req *http.Request, ctx *DB.Context) {
	vars := mux.Vars(req)
	agreementID := vars["agreementID"]

	comment := new(models.Comment)

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	reqBytes := buf.Bytes()
	json.Unmarshal(reqBytes, &comment)
	comment.AgreementID = agreementID
	_ = comment.SaveWithCtx(ctx)

	a, _ := json.Marshal(comment)
	w.Write(a)
}

func GetComments(w http.ResponseWriter, req *http.Request, ctx *DB.Context) {
	vars := mux.Vars(req)
	id := vars["agreementID"]
	comments, _ := models.FindCommentsByAgreementID(id, ctx)

	u, _ := json.Marshal(comments)
	w.Write(u)

}
