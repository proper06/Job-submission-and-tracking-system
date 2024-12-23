package main

import (
	"log"
	"net/http"

	"github/proper06/handLers"

	"github/proper06/config"

	"github.com/gorilla/mux"
)

func main() {
    cfg := config.LoadConfig()
    
    r := mux.NewRouter()
    
    // Routes
    r.HandleFunc("/api/submit", handlers.SubmitJobHandler).Methods("POST")
    r.HandleFunc("/api/status", handlers.GetJobStatusHandler).Methods("GET")
    
    log.Printf("Server starting on port %s", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
