package main

import (
	"fmt"
	"net/http"
)

func main() {
	// func(w http.ResponseWrite, r *http.Request) is a handler that
	// receives all incoming HTTP connections.
	// w is where I write my text/HTML response to
	// r is where I get all the information about the request, including:
	//	- URL
	//	- header fields
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(
			w,
			"<h1>The Title</h1>"+
				"Hello viewer, you've requested: %s\n",
			r.URL.Path)
	})

	// Starts Go's default HTTP server and listens for connections on port 80.
	// Once started, navigate to http:<IP>/ to see the webpage.
	http.ListenAndServe(":80", nil)
}
