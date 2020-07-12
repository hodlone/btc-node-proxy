package server

import (
	"fmt"
	"log"
	"net/http"
)

// Start runs the server which will receive block information
func Start(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Printf("Proxy healthcheck online on port:  %v", port)
	port = fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
