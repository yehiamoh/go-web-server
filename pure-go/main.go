package main

import (
	"fmt"
	"net/http"
	"text/template"

	"example.com/go/webserver/api"
	"example.com/go/webserver/data"
)

// handleHelloHandler writes a simple hello message to the response
func handleHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello my first go app !!!!!"))
}
func handelTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error: unable to parse template"))
		fmt.Println("Error parsing template:", err)
		return
	}

	// Get the data
	data := data.GetAll()
	fmt.Println("Data passed to template:", data) // Log the data for debugging

	// Execute the template with the data
	err = html.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error executing template:", err)
		return
	}
}
func main() {
	// Create a new ServeMux to handle incoming HTTP requests
	server := http.NewServeMux()

	// Create a file server to serve static files from the "./public" directory
	fs := http.FileServer(http.Dir("./public"))

	// Register the handleHelloHandler function to handle requests to the "/home" endpoint
	server.HandleFunc("/home", handleHelloHandler)
	server.HandleFunc("/template", handelTemplate)
	server.HandleFunc("/api/Exhibitions", api.GET)
	server.HandleFunc("/api/Exhibitions/create", api.POST)
	// Register the file server to handle requests to the root ("/") endpoint
	server.Handle("/", fs)

	// Start the HTTP server on port 3030 and use the ServeMux to handle requests
	err := http.ListenAndServe(":3030", server)
	if err != nil {
		// Print an error message if the server fails to start
		fmt.Printf("Error in initializing the server: %v", err)
	}
}
