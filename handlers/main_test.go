package handlers

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

func Test_all(t *testing.T) {
	test_GetTags(t)

	DB.DB.Exec("DELETE from comment")
	DB.DB.Exec("DELETE from tag")
}
