package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	types "github.com/harshshekhar15/monitoring/pkg/types"
)

var Servers []types.Server

var Hosts []string

var Report []types.Usage

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homepage)
	router.HandleFunc("/metrics", ingestMetrics).Methods("POST")
	router.HandleFunc("/report", generateReport)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Metrics server running on http://localhost:8080")
	// Servers stores metrics of clients
	Servers = []types.Server{}
	// Hosts stores IP address of clients
	Hosts = []string{}
	handleRequests()
}
