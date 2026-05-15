package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"urlshortener/config"
	"urlshortener/db"
	"urlshortener/handler"
	"urlshortener/shortener"
	"urlshortener/util"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database
	dbConn, err := sqlx.Connect("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Create a new shortener
	shortener := shortener.NewShortener(dbConn)

	// Create a new handler
	h := handler.NewHandler(shortener)

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/shorten", h.Shorten).Methods("POST")
	r.HandleFunc("/{shortCode}", h.Redirect).Methods("GET")

	// Start server
	log.Println("Server listening on port", config.Port)
	http.ListenAndServe(":"+config.Port, r)
}