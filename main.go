package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.URL.Host = "www." + r.Host
	http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
}
