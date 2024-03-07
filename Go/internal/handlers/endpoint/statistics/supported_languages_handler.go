// Package statistics provides handlers for statistics-related endpoints.
package statistics

import (
	"encoding/json"
	_func "gutendex_api/internal/func"
	"net/http"
	"sort"
)

// SupportedLanguagesHandler handles HTTP GET requests to retrieve supported languages.
func SupportedLanguagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleSupportedLanguagesGetRequest(w)
	default:
		http.Error(w, "REST Method: "+r.Method+" not supported. Only supported method for this endpoint is: "+http.MethodGet, http.StatusNotImplemented)
		return
	}
}

// handleSupportedLanguagesGetRequest handles GET requests to retrieve supported languages.
func handleSupportedLanguagesGetRequest(w http.ResponseWriter) {
	// Get the supported languages from self defined function
	languages, err := _func.GetSupportedLanguages()
	if err != nil {
		http.Error(w, "Error whilst retrieving languages: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Sort the language entries alphabetically by name
	sort.Slice(languages, func(i, j int) bool {
		return languages[i].Name < languages[j].Name
	})

	// Write content type header
	w.Header().Set("Content-Type", "application/json")

	// Encode languages as JSON and send the response
	if err := json.NewEncoder(w).Encode(languages); err != nil {
		http.Error(w, "Error encoding JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
