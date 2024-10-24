package main

import (
	"fmt"
	"log"
	"net/http"
)

var webPort = ":8080"

func hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello Welcome to this SERVER!\n"))
}

func hellotwo(w http.ResponseWriter, r *http.Request) {

	headers := r.Header
	headersstring := "The Headers are:\n"

	for key, values := range headers {
		for _, value := range values {
			headersstring += fmt.Sprintf("%s: %s\n", key, value)
		}
	}

	_, _ = w.Write([]byte(headersstring))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", hello)
	mux.HandleFunc("GET /another", hellotwo)

	log.Printf("Server is running on port%s", webPort)
	err := http.ListenAndServe(webPort, mux)
	if err != nil {
		log.Println("Error launching server", err)
	}

}
