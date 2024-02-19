package main

import (
	"budget-app/go-backend/src/models"
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	models.ConnectDatabase()

	fmt.Println("Starting server and listening on localhost:8080")
	http.ListenAndServe("localhost:8080", mux)

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
