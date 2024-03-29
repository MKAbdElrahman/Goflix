package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func stream(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "streaming...")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/video/{$}", stream)

	log.Info("starting server")
	http.ListenAndServe(":3000", mux)
}
