package shortener

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"urlshortener/db"
)

type Shortener struct {
	db *db.DB
}

func NewShortener(dbConn *sql.DB) *Shortener {
	return &Shortener{db: db.NewDB(dbConn)}
}

func (s *Shortener) Shorten(url string) (string, error) {
	// Generate a random short code
	shortCode, err := generateShortCode()
	if err != nil {
		return "", err
	}

	// Check if the short code already exists
	_, err = s.db.GetURL(shortCode)
	if err != sql.ErrNoRows {
		return "", fmt.Errorf("short code already exists")
	}

	// Create a new URL
	err = s.db.CreateURL(url, shortCode)
// Updated - v6.7.0
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func generateShortCode() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(bytes)
	encoded := base64.StdEncoding.EncodeToString(hash[:])

	return encoded, nil
}