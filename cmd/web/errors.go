package main

import (
	"net/http"

	"github.com/antonisgkamitsios/simple-todo-app/view"
)

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *application) renderError(w http.ResponseWriter, r *http.Request, status int, title, message string, replacePage bool) {

	w.Header().Add("Vary", "HX-Retarget")

	if replacePage {
		w.Header().Add("HX-Retarget", "#contents")
	} else {
		w.Header().Add("HX-Retarget", "#errors")
		w.Header().Add("HX-Reswap", "innerHTML ")
	}

	app.render(w, r, status, view.Error(title, message, replacePage))
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error, replacePage bool) {
	app.logError(r, err)

	app.renderError(
		w,
		r,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"Something went wrong",
		replacePage,
	)
}

func (app *application) clientError(w http.ResponseWriter, r *http.Request, status int, message string, replacePage bool) {
	app.renderError(
		w,
		r,
		status,
		http.StatusText(status),
		message,
		replacePage,
	)
}
