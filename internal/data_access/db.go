package dataaccess

import (
	"database/sql"
)

func ConnectDatabase(dbPath string) (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return &sql.DB{}, err
	}

	return DB, nil
}
