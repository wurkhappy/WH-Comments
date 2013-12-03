package DB

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	// "log"
)

var UpsertComment *sql.Stmt
var FindCommentsByAgreementID *sql.Stmt
var FindCommentsByVersionID *sql.Stmt

var SaveTag *sql.Stmt
var FindTagsByAgreementID *sql.Stmt

func CreateStatements() {
	var err error
	UpsertComment, err = DB.Prepare("SELECT upsert_comment($1, $2)")
	if err != nil {
		panic(err)
	}

	FindCommentsByAgreementID, err = DB.Prepare("SELECT data FROM comment WHERE data->>'agreementID' = $1")
	if err != nil {
		panic(err)
	}

	FindCommentsByVersionID, err = DB.Prepare("SELECT data FROM comment WHERE data->>'agreementVersionID' = $1")
	if err != nil {
		panic(err)
	}

	SaveTag, err = DB.Prepare("INSERT INTO tag (id, data) VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}

	FindTagsByAgreementID, err = DB.Prepare("SELECT data FROM tag WHERE data->>'agreementID' = $1")
	if err != nil {
		panic(err)
	}
}
