package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func (app *application) render(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	component templ.Component) {

	w.WriteHeader(status)
	err := component.Render(r.Context(), w)
	if err != nil {
		app.logError(r, err)
		http.Error(w, http.StatusText(status), status)
	}

}
