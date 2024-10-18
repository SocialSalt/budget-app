package server

import (
	"log"
)

func main() {
	DB, err := ConnectDatabase("budget.db")
	if err != nil {
		log.Fatalf("Encountered an error while connecting to database %#v", err)
	}
	defer DB.Close()

}
