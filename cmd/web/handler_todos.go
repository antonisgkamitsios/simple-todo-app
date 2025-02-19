package main

import (
	"net/http"

	"github.com/antonisgkamitsios/simple-todo-app/view"
)

func (app *application) handleTodos(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, view.Todos())
}
