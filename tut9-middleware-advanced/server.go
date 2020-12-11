package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Middleware Template
func createMiddleware() Middleware {
	// Create a new Middleware
	middleware := func(next http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc which is called by the server eventually
		handler := func(w http.ResponseWriter, r *http.Request) {

			// ... functions here

			// Call the next middleware in the chain
			next(w, r)
		}

		// Return the newly created handler for the middleware
		return handler
	}

	// Return the newly created middleware
	return middleware
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Example Middleware-Level functions

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Example Application-Level functions

// Hello prints "hello world" to the console
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

// HelloPerson prints hello to the specified person
func HelloPerson(w http.ResponseWriter, r *http.Request) {
	person := r.FormValue("person")
	fmt.Fprintf(w, "hello %s", person)
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.HandleFunc("/hello", Chain(HelloPerson, Method("POST"), Logging()))

	http.ListenAndServe(":8080", nil)
}
