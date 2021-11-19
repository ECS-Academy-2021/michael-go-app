// Go imports are really simple, just a string per line
package main

import (
	"fmt"
	"log"
	"net/http"
)

// It's important to have a main method
  func main() {
	// For the route "/", call an anonymous function which returns "Hello, Michael" AND logs that response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	  response := "Hello, Michael"
	  fmt.Fprint(w, response)
	  log.Println("GET '/':", response)
	})
	// Avoid odd duplicate calls caused by browsers automatically requesting the "favicon" endpoint ¬_¬
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request){})
	// Listen on port 8080
	http.ListenAndServe(":8080", nil)
  }