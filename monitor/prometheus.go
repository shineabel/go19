package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
)

func main() {

	http.Handle("/metrics",promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080",nil))
}
