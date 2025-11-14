package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"os"
	"fmt"
)

type Response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type StatusResponse struct {
	Status  string    `json:"status"`
	Version string    `json:"version"`
	Uptime  string    `json:"uptime"`
}

var startTime time.Time
var name string

func init() {
	startTime = time.Now()
	name = os.Getenv("SERVICE_NAME")
	if name == "" {
		name = "default-service"
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := Response{
		Message:   fmt.Sprintf("Hello from %s!", name),
		Timestamp: time.Now(),
	}
	
	json.NewEncoder(w).Encode(response)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	uptime := time.Since(startTime).Round(time.Second).String()
	
	response := StatusResponse{
		Status:  "ok",
		Version: "1.0.0",
		Uptime:  uptime,
	}
	
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/greetings", helloHandler)
	http.HandleFunc("/status", statusHandler)
	
	port := fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))
	log.Printf("Starting server on port %s", port)
	log.Printf("Endpoints available:")
	log.Printf("  - GET %s/hello", port)
	log.Printf("  - GET %s/status", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
