package main

import (
	"log"
	"net/http"

	"aqua-nav/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/control", handlers.ControlBoat).Methods("POST")
	r.HandleFunc("/api/log", handlers.LogData).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
