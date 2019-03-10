package main

import (
	"log"
	"net/http"
)

// AuthenticationMiddleware will do stuff
type AuthenticationMiddleware struct {
	tokenUsers map[string]string
}

// Populate will seed the app with users
func (amw *AuthenticationMiddleware) populate() {
	amw.tokenUsers["000000"] = "User 1"
	amw.tokenUsers["111111"] = "User 2"
	amw.tokenUsers["222222"] = "User 3"
	amw.tokenUsers["333333"] = "User 4"
}

// Middleware will compare the Authentication header against the list of known users
func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authentication")

		if user, found := amw.tokenUsers[token]; found {
			log.Printf("Found user: %s\n", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// MakeSampleAMW will return an AuthenticationMiddleware struct with a pre-populated user list
func MakeSampleAMW() AuthenticationMiddleware {
	amw := AuthenticationMiddleware{}
	amw.tokenUsers = make(map[string]string)
	amw.populate()

	return amw
}
