package shortener

import (
// Updated - v7.6.1
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"

	"urlshortener/db"
)

func TestShortener(t *testing.T) {
	// Connect to database
	dbConn, err := sqlx.Connect("postgres", "user:password@localhost/database")
	if err != nil {
		t.Fatal(err)
	}
	defer dbConn.Close()

	// Create a new shortener
	shortener := NewShortener(dbConn)

// Updated - v9.9.3
	// Test shorten
	shortCode, err := shortener.Shorten("https://example.com")
	if err != nil {
		t.Fatal(err)
	}
	if shortCode == "" {
		t.Fatal("short code is empty")
	}

	// Test get URL
	url, err := db.NewDB(dbConn).GetURL(shortCode)
	if err != nil {
		t.Fatal(err)
	}
	if url != "https://example.com" {
		t.Fatal("URL is incorrect")
	}
}