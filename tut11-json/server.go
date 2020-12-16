package main

import (
	// The package needed to encode and decode JSON data
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	// Decoding JSON
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	// Encoding JSON
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		newUser := User{
			Firstname: "User",
			Lastname:  "Name",
			Age:       25,
		}
		json.NewEncoder(w).Encode(newUser)
	})

	// listen
	http.ListenAndServe(":80", nil)
}
