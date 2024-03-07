// Package statistics provides handlers for statistics-related endpoints.
package statistics

import (
	"encoding/json"
	"fmt"
	"gutendex_api/internal/utils/constants"
	"gutendex_api/internal/utils/constants/External"
	"gutendex_api/internal/utils/structs"
	"log"
	"net/http"
	"time"
)

// getEndpointStatus returns the HTTP status code of the provided endpoint.
func getEndpointStatus(endpointURL string) string {
	// Create new request
	r, err := http.NewRequest(http.MethodGet, endpointURL, nil)
	if err != nil {
		// Log and handle the error if request creation fails.
		err := fmt.Errorf("error in creating request: %s", err.Error())
		log.Println(err)
	}

	// Set content type header
	r.Header.Add("content-type", "application/json")

	// Create HTTP client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	if err != nil {
		// Log and handle the error if request execution fails.
		err := fmt.Errorf("error in response: %s", err.Error())
		log.Println(err)
	}

	return res.Status
}

// StatusHandler handles requests to retrieve the status of Endpoints.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleStatusGetRequest(w)
	default:
		// If method is not supported, return an error response.
		http.Error(w, "REST Method: "+r.Method+" not supported. Only supported method for this endpoint is: "+http.MethodGet, http.StatusNotImplemented)
		return
	}
}

// handleStatusGetRequest handles GET requests to retrieve the status of Endpoints.
func handleStatusGetRequest(w http.ResponseWriter) {
	// Initialize a status response.
	status := structs.StatusResponse{
		GutendexAPI:  getEndpointStatus(External.GutendexAPI),
		LanguageAPI:  getEndpointStatus(External.LanguageAPI + "no"),
		CountriesAPI: getEndpointStatus(External.CountriesAPI + "all"),
		Version:      constants.APIVersion,
		// Calculate uptime since the last restart of the service.
		UptimeInSeconds: fmt.Sprintf("%f Seconds", time.Since(startTime).Seconds()),
	}

	// Set content type header
	w.Header().Add("content-type", "application/json")

	// Encode status as JSON and send the response.
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		// If encoding fails, return an error response.
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// startTime keeps track of the service start time.
var startTime = time.Now()
