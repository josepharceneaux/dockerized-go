// Simple web service

package main

import (
	"fmt"
	"log"
	"net/http"
)

const TITLE = "<h2>Automation for the People</h2>"

// Simple handler to print a title
func getHomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, TITLE)
	})
}

// Log our requests
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// Register our simple handler and start listening for requests
func main() {
	http.Handle("/", getHomeHandler())
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
}
