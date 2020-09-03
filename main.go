package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.URL.Host = "www." + r.Host
	http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
}
