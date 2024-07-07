package main

import (
	"log"
	"net/http"

	"aqua-nav/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/control", handlers.ControlBoat).Methods("POST")
	r.HandleFunc("/api/log", handlers.LogData).Methods("POST")

	http.Handle("/", r)
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
