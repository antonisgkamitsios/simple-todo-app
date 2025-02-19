package main

import (
	"net/http"

	"github.com/antonisgkamitsios/simple-todo-app/view"
)

func (app *application) handleHome(w http.ResponseWriter, r *http.Request) {
	// use the new app.render function
	app.render(w, r, http.StatusOK, view.Index())
}
