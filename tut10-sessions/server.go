package main

/*
This example will only allow authenticated users to view the "/secret" page.

They must first log in on the "/login" page with the correct credentials to get
a valid session cookie.

They can then visit "/logout" to release the cookie.
*/

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	// this popular package can be used to store data in session cookies
)

type User struct {
	Username      string
	Authenticated bool
}

var store *sessions.CookieStore
var userList = map[string]string{
	"user": "pass",
	"a":    "abc",
	"go":   "lang",
	"z":    "z",
}

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		// 15 minute cookie sessions
		MaxAge: 60 * 15,
	}

	gob.Register(User{})
}

// getUser returns a user from session s
// on error returns an empty user
func getUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...
	auth := true

	// only allow post requests
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// parse credentials
	username := r.FormValue("username")
	password := r.FormValue("password")

	// check if user exists
	realPassword, ok := userList[username]
	if auth && !ok {
		auth = false
	}

	// check if password matches
	if auth && realPassword != password {
		auth = false
	}

	// Log user in if successful
	if auth {
		user := &User{
			Username:      username,
			Authenticated: true,
		}
		session.Values["user"] = user
		session.Save(r, w)
		fmt.Fprintf(w, "Logged in! Welcome %s!", username)
	} else {
		fmt.Fprintln(w, "Your login credentials are incorrect!")
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Log user out
	session.Values["user"] = User{}
	session.Options.MaxAge = -1
	session.Save(r, w)

	fmt.Fprintln(w, "Logged out! See ya!")
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "cookie-name")

	user := getUser(session)
	fmt.Fprintf(w, "Checking... %s", user.Username)

	// Check if the user is not authenticated
	if auth := user.Authenticated; !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false
	}
	return true
}

func secret(w http.ResponseWriter, r *http.Request) {
	// only display if the user is authorized
	if isAuthenticated(w, r) {
		// Display secret
		fmt.Fprintln(w, "You found the secret!")
	}
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":80", nil)
}

/* Example usage

$ curl -s http://localhost/secret
Forbidden

$ curl -X POST -F 'username=wrong' -F 'password=creds' http://localhost/login
Your login credentials are incorrect!

// use -i option to get the response headers
$ curl -i -X POST -F 'username=user' -F 'password=pass' http://localhost/login
HTTP/1.1 200 OK
Set-Cookie: cookie-name=MTYwNzk3NTkw...; Path=/; Expires=Wed, 13 Jan 2021 19:58:23 GMT; Max-Age=2592000
Date: Mon, 14 Dec 2020 19:58:23 GMT
Content-Length: 23
Content-Type: text/plain; charset=utf-8

$ curl -s --cookie "cookie-name=MTYwNzk3NTkw..." http://localhost/secret
You found the secret!

$ curl -s http://localhost/logout
Logged out! See ya!


*/
