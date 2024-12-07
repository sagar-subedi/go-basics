package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Define a struct for the API response
type ApiResponse struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

// Handler struct to organize methods
type Server struct {
	mu sync.Mutex // For concurrency-safe operations
}

// Health check handler
func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := ApiResponse{
		Message: "Server is healthy",
		Time:    time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// Echo handler: echoes back user input
func (s *Server) Echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	log.Printf("Received data: %v", input)
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(input); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// NotFound handler for unmatched routes
func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Endpoint not found", http.StatusNotFound)
}

// Main function
func main() {
	server := &Server{}

	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/health", server.HealthCheck)
	mux.HandleFunc("/echo", server.Echo)

	// Use NotFoundHandler for unmatched routes
	mux.HandleFunc("/", server.NotFound)

	// Create a server with timeouts
	srv := &http.Server{
		Addr:         ":8070",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on :8070")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
