package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// redefine to healthServerPort and HEALTH_SERVER_PORT
var (
	serverPort  = os.Getenv("HTTP_SERVER_PORT")
	serviceName = os.Getenv("SERVICE_NAME")
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
		Status:    fmt.Sprintf("%v is  Healthy!", serviceName),
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

// StartHealthServer runs the server which will receive block information
func StartHealthServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", healthCheck)
	log.Printf("%v healthcheck online on port: %v", serviceName, serverPort)
	p := fmt.Sprintf(":%v", serverPort)
	log.Fatal(http.ListenAndServe(p, mux))
}
