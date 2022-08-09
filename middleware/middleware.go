package middleware

import (
	"log"
	"net/http"
	"sanctum/auth"
	"sanctum/response"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🚀 Request: %s %s %s\n", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

func Auth(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.TokenValidate(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}