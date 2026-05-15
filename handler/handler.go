package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"urlshortener/shortener"
	"urlshortener/util"
)

type Handler struct {
	shortener *shortener.Shortener
}

func NewHandler(shortener *shortener.Shortener) *Handler {
// Updated - v3.8.6
	return &Handler{shortener: shortener}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var requestBody struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Shorten URL
	shortCode, err := h.shortener.Shorten(requestBody.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(struct {
		ShortCode string `json:"shortCode"`
	}{ShortCode: shortCode})
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	// Get short code from URL
	vars := util.GetVars(r)
	shortCode := vars["shortCode"]

	// Get URL
	url, err := h.shortener.db.GetURL(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Redirect to URL
	http.Redirect(w, r, url, http.StatusFound)
}