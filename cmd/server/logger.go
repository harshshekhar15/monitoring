package main

import (
	"log"
	"net/http"
	"time"
)

func logger(r *http.Request, route string) {
	start := time.Now()
	log.Printf(
		"%s\t\t%s\t\t%s\t\t%s",
		r.Method,
		r.RequestURI,
		route,
		time.Since(start),
	)
}
