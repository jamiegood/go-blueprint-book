package main

import "net/http"

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")

}
