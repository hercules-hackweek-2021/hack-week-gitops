package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		fmt.Fprint(w, `{"status":"ok"}`)
	})
	log.Fatalf("error: %s", http.ListenAndServe(":8080", nil))
}
