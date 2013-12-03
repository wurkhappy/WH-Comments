package models

import (
	"github.com/wurkhappy/WH-Comments/DB"
	// "github.com/nu7hatch/gouuid"
	"testing"
	// "log"
)

func init() {
	DB.Name = "testdb"
	DB.Setup(false)
	DB.CreateStatements()
}

func TestIntegrationTests(t *testing.T) {
	if !testing.Short() {
		test_SaveComment(t)
		test_FindCommentsByAgreementID(t)
		test_FindCommentsByVersionID(t)
		test_createNewTags(t)

		test_SaveTag(t)
		test_FindTagsByAgreementID(t)

		DB.DB.Exec("DELETE from comment")
		DB.DB.Exec("DELETE from tag")
	}
}

func TestUnitTests(t *testing.T) {
	test_NewTag(t)
}
