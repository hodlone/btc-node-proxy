package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// HealthStatus ...
type HealthStatus struct {
	Status    string
	TimeStamp string
}

// healthCheck ...
func healthCheck(res http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()

	healthStatus := HealthStatus{
		Status:    "btc-node-proxy is Healthy!",
		TimeStamp: currentTime.Format("2006-01-02 15:04:05"),
	}

	hj, err := json.Marshal(healthStatus)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", hj)
}

// Start runs the server which will receive block information
func Start(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", healthCheck)
	log.Printf("Proxy healthcheck online on port: %v", port)
	port = fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
