package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/socialsalt/budget-app/cmd/server/handlers"
	"github.com/socialsalt/budget-app/cmd/server/models"
)

func main() {
	models.ConnectDatabase()

	mux := http.NewServeMux()
	mux.HandleFunc("/balance", handlers.GetBalance)
	mux.HandleFunc("/budget", handlers.GetBudget)
	mux.HandleFunc("/transaction", handlers.Transaction)
	mux.HandleFunc("/upload/transaction", handlers.UploadTransaction)
	mux.HandleFunc("/login/chase", handlers.LoginChase)
	mux.HandleFunc("/", serveMain)

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Starting server and listening on localhost:8080")
	http.ListenAndServe("localhost:8080", mux)

}

func serveMain(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
