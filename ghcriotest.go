package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Started GCRIOtest")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	setCommonHeaders(w)
	_, err := fmt.Fprint(w, "<html><head><title>GHCRIOtest</title></head><body><h1>Success</h1><p>You are successfully running the GHCRIO test</p></body></html>")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
}
