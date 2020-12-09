package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Middleware is software that's between an operating system and the applications running on it.
Enables communication for fistributed applications.
Ex.
	- database
	- application server
	- message-oriented
	- transaction-processing monitors

Use cases for data transmitted:
	- security authentication
	- transaction management
	- message queues
	- directories
*/

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Entry: %s", r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "print foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "print bar")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":80", nil)
}
