package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string `json:"port"`
	DatabaseURL string `json:"databaseUrl"`
}

func LoadConfig() (*Config, error) {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	// Load configuration file
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal configuration file
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}