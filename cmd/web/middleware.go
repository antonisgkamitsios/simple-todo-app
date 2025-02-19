package main

import (
	"net/http"

	"github.com/antonisgkamitsios/simple-todo-app/internal/htmx"
)

func (app *application) comesFromHTMX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if the request has "Hx-History-Restore-Request" is means it had a cache miss
		// and it wants back the entire html
		isHTMXRequest := r.Header.Get("Hx-Request") != "" && r.Header.Get("Hx-History-Restore-Request") == ""

		next.ServeHTTP(w, htmx.ContextSetHtmxRequest(r, isHTMXRequest))
	})
}
