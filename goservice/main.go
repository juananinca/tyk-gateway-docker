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

func submitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Field1 float64 `json:"field1"`
		Field2 string  `json:"field2"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := map[string]interface{}{
		"received":  req,
		"message":   fmt.Sprintf("Received field1=%s and field2=%s", req.Field1, req.Field2),
		"timestamp": time.Now(),
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/submit", submitHandler)
	
	port := fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT"))
	log.Printf("Starting server on port %s", port)
	log.Printf("Endpoints available:")
	log.Printf("  - GET %s/hello", port)
	log.Printf("  - GET %s/status", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
