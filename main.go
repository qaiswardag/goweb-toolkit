package main

import (
	"fmt"
	"net/http"

	"github.com/qaiswardag/goweb-toolkit/toolkit"
)

// Define the structure of the expected JSON
type RequestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoggingHandler struct{}

func (m LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:7777")
	w.Header().Set("Cache-Control", "no-cache, no-store")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept-Version")

	//
	//
	//
	//
	//
	//
	//
	//
	myTool := toolkit.Tools{}

	// Check if the request method is not POST or GET
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path == "/" {
		fmt.Println("Request URL:", r.URL.String())

		// Create an instance of the struct to hold the parsed JSON
		requestData := RequestData{}

		if r.Method == http.MethodPost {
			// Use ReadJSON to parse the JSON into the struct
			err := myTool.ReadJSON(w, r, &requestData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		// Process the parsed JSON data
		fmt.Printf("Parsed JSON: %+v\n", requestData)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JSON received successfully\n"))
		return
	}

	if r.URL.Path != "/" {
		fmt.Println("Route not found:", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not allowed to enter this page"))
	}
}

func main() {
	loggingHandler := LoggingHandler{}

	server := http.Server{
		Addr:    ":9999",
		Handler: loggingHandler,
	}

	server.ListenAndServe()
}
