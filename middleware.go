package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func sessionload(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// prevent csrf Attack
func Nosurf(next http.Handler) http.Handler {
	csfhandler := nosurf.New(next)
	csfhandler.SetBaseCookie(
		http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Secure:   App.AppSecureCookie,
			SameSite: http.SameSiteLaxMode,
		},
	)
	return csfhandler
}
