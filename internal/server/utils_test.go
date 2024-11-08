package server_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// // assert fails the test if the condition is false.
// func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
// 	if !condition {
// 		_, file, line, _ := runtime.Caller(1)
// 		fmt.Printf("%s:%d: "+msg, append([]interface{}{filepath.Base(file), line}, v...)...)
// 		tb.FailNow()
// 	}
// }
//
// // ok fails the test if an err is not nil.
// func ok(tb testing.TB, err error) {
// 	if err != nil {
// 		_, file, line, _ := runtime.Caller(1)
// 		fmt.Printf("%s:%d: unexpected error: %s", filepath.Base(file), line, err.Error())
// 		tb.FailNow()
// 	}
// }
//
// // equals fails the test if exp is not equal to act.
// func equals(tb testing.TB, exp, act interface{}) {
// 	if !reflect.DeepEqual(exp, act) {
// 		_, file, line, _ := runtime.Caller(1)
// 		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v", filepath.Base(file), line, exp, act)
// 		tb.FailNow()
// 	}
// }

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
