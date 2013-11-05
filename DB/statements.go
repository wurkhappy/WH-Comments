package DB

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	// "log"
)

var UpsertComment *sql.Stmt
var FindCommentsByAgreementID *sql.Stmt
var FindCommentsByVersionID *sql.Stmt

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
}
