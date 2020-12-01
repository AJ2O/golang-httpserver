package main

// 	Basic key jobs of an HTTP server:
// 	1. Process Dynamic Requests
//		- incoming user reqeusts to browse sire, log into accounts, post images, etc.
//	2. Serve Static Assets
//		- Serve JS, CSS, and images to browsers to create a dynamix user experience
//	3. Accept Connections
//		- listen to a port and be able to accept connections from the Internet

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Process dynamic requests
	// This block displays "Welcome to my website!", when a person browses to
	// the root of the website.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")

		// r contains all the information about the request and it's parameters
		// GET parameters can be read with r.URL.Query().Get("variable")
		// POST parameters can be read with r.FormValue("variable")
	})

	// 2. Serve Static Assets
	// To serve static assets (JS, CSS, images, etc.), use the inbuilt
	// http.FileServer and point it to a url path. This will act as the
	// file server.
	fs := http.FileServer(http.Dir("static/"))

	// Now that the file server is in place, we can point a URL path to it,
	// similar to how we handled dynamic requests.
	// Note: To serve files from a server correctly, we need to just strip away
	// the name of the directory the files live in.
	// Note: Run this server using this folder as the working directory for the
	// file server to be viewable. Otherwise, there will be 404 error.
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Note: Example image obtained from https://golang.org/

	// 3. Accept Connections
	// Set the program to listen on a specific port to accept Internet connections.
	// Once started, it is viewable in the web browser.
	http.ListenAndServe(":80", nil)
}
