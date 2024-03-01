package main

import (
	"fmt"
	"go-backend/src/handlers"
	"go-backend/src/models"
	"net/http"
)

func main() {
	models.ConnectDatabase()

	mux := http.NewServeMux()
	mux.HandleFunc("/balance", handlers.GetBalance)
	mux.HandleFunc("/budget", handlers.GetBudget)
	mux.HandleFunc("/transaction", handlers.Transaction)
	mux.HandleFunc("/upload/transaction", handlers.UploadTransaction)

	fmt.Println("Starting server and listening on localhost:8080")
	http.ListenAndServe("localhost:8080", mux)

}
