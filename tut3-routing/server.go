package main

// net/http doesn't do complex routing very well, such as segmenting a URL into
// single parameters. A popular Go package that can solve such problems is
// gorilla/mux.
// We can use this package to create routes iwith named parameters, GET/POST handlers,
// and domain restrictions.

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// This is the main router for the web application and will later be passed
	// as a parameter to the server. It receives all HTTP connections and passes
	// them on to the request handlers that are registered to it.
	r := mux.NewRouter()

	// The only difference is that instead of calling http.HandleFunc(...),
	// it is instead r.HandleFunc(...).

	// This example uses gorilla/mux to extract segments from the request URL.
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// We can use mux.Vars(r) to map the data from these segments
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: '%s' on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)

	// It's possible to restrict request handlers for specific HTTP methods
	//r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	//r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	//r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	//r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	// It's possible to restrict the request handler for specific hostnames
	// or subdomains.
	//r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	// It's possible to restrict the handler for http/https
	//r.HandleFunc("/secure", SecureHandler).Schemes("https")
	//r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	// Restrict the handler to specific path prefixes
	//bookrouter := r.PathPrefix("/books").Subrouter()
	//bookrouter.HandleFunc("/", AllBooks)
	//bookrouter.HandleFunc("/{title}", GetBook)
}
