package dataaccess_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitTestDB(t *testing.T) *sql.DB {
	os.Remove("./test_db.db")
	db, err := sql.Open("sqlite3", "./test_db.db")
	assert.NoError(t, err)

	schema, err := os.ReadFile("schema.sql")
	assert.NoError(t, err)

	// fmt.Println(string(schema))

	_, err = db.Exec(string(schema))
	assert.NoError(t, err)

	return db
}
