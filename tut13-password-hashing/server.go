package main

import (
	"html/template"
	"net/http"

	// This library will allow us to hash values using bcrypt
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the hash of the given string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash returns true if the given string hashes to the given hash value
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	hash, _ := HashPassword(password)
	tmpl := template.Must(template.ParseFiles("home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		enteredPassword := r.FormValue("password")

		// no entered password means no value checking
		if len(enteredPassword) == 0 {
			tmpl.Execute(w, nil)
			return
		}

		// validate the password
		if CheckPasswordHash(enteredPassword, hash) {
			tmpl := template.Must(template.ParseFiles("secret.html"))
			tmpl.Execute(w, struct {
				Password string
				Hash     string
			}{enteredPassword, hash})
		} else {
			tmpl.Execute(w, struct{ Fail bool }{true})
		}
	})

	http.ListenAndServe(":80", nil)
}
