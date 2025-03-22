package api

import (
	"encoding/json"
	"net/http"

	"example.com/go/webserver/data"
)

// POST handles HTTP POST requests to save an exhibition.
// It expects a JSON payload in the request body and responds with appropriate status codes.
func POST(w http.ResponseWriter, r *http.Request) {
	// Check if the HTTP method is POST
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition

		// Ensure the request body is closed after processing to avoid resource leaks
		defer r.Body.Close()

		// Decode the JSON payload into the exhibition struct
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			// Respond with a 400 Bad Request if the JSON is invalid
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the exhibition to the data store
		// Ensure that `AddExhibition` handles any potential errors internally
		data.AddExhibiton(exhibition)

		// Set the response status to 201 Created
		w.WriteHeader(http.StatusCreated)

		// Set the Content-Type header to indicate the response format
		w.Header().Set("Content-Type", "Content-Type/json")

		// Write a success message to the response body
		w.Write([]byte("Exhibition saved successfully"))
	} else {
		// Respond with a 405 Method Not Allowed for non-POST requests
		http.Error(w, "Invalid verb", http.StatusMethodNotAllowed)
	}
}
