package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func NewDB(dbConn *sqlx.DB) *DB {
	return &DB{dbConn}
}

func (d *DB) GetURL(shortCode string) (string, error) {
	var url string
	err := d.Get(&url, "SELECT url FROM urls WHERE short_code = $1", shortCode)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (d *DB) CreateURL(url, shortCode string) error {
	_, err := d.Exec("INSERT INTO urls (url, short_code) VALUES ($1, $2)", url, shortCode)
	if err != nil {
		return err
	}
	return nil
}