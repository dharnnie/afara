package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dharnnie/tai/site/handlers"
	"github.com/gorilla/mux"
)

func main() {
	serve()
}

func serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/assets/", handlers.ServeResource)

	myMux := mux.NewRouter()
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.Handle("/", myMux)
	err := (http.ListenAndServe(":"+port, nil))
	if err != nil {
		log.Fatal("Server error", err)
	}
}
