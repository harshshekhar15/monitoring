package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/gddo/httputil/header"
	types "github.com/harshshekhar15/monitoring/pkg/types"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Metrics server running!!")
	logger(r, "Homepage")
}

// ingestMetrics stores metrics of client
func ingestMetrics(w http.ResponseWriter, r *http.Request) {
	logger(r, "IngestMetrics")
	var newServer types.Server
	// throw error in case Content-Type is not set
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	}
	dec := json.NewDecoder(r.Body)
	// disallow unknown fields in payload
	dec.DisallowUnknownFields()
	err := dec.Decode(&newServer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get IP address of the client
	IP, err := getIP(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newServer.IP = IP
	Servers = append(Servers, newServer)
	// add client's IP to the Host array
	addHost(IP)
}

// generateReport generates a report of max CPU and memory usage of clients
func generateReport(w http.ResponseWriter, r *http.Request) {
	logger(r, "Report")
	// throw error in case Content-Type is not set
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	}
	Report = []types.Usage{}
	// generate report if client's metrics is ingested
	if len(Servers) != 0 {
		for _, host := range Hosts {
			var newUsage types.Usage
			newUsage.IP = host
			newUsage.MaxCPU = getMaxCPU(host)
			newUsage.MaxMemory = getMaxMemory(host)
			Report = append(Report, newUsage)
		}
	}
	_ = json.NewEncoder(w).Encode(Report)
}
