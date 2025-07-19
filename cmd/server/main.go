package server

import (
	"log"

	dataaccess "github.com/socialsalt/budget-app/internal/data_access"
)

func main() {
	DB, err := dataaccess.ConnectDatabase("budget.db")
	if err != nil {
		log.Fatalf("Encountered an error while connecting to database %#v", err)
	}
	defer DB.Close()

}
