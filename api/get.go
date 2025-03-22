package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/go/webserver/data"
)

func GET(w http.ResponseWriter, r *http.Request) {

	// set header
	w.Header().Set("Content-Type", "application/json")

	// Get the "id" query parameter
	id := r.URL.Query().Get("id")

	if id != "" { // Check if "id" is provided

		// Convert "id" from string to int
		intID, err := strconv.Atoi(id)

		// Validate the "id"
		if err == nil && intID >= 0 && intID < len(data.GetAll()) {

			// Send the specific item in the response
			json.NewEncoder(w).Encode(data.GetAll()[intID])

		} else {
			// Return error for invalid "id"
			http.Error(w, "Invalid Query Parameter", http.StatusBadRequest)
		}
	} else { // Return all items if "id" is not provided
		jsonEncoder := json.NewEncoder(w)
		jsonEncoder.Encode(data.GetAll())
	}

}
