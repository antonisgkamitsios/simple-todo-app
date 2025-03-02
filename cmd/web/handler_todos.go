package main

import (
	"net/http"

	"github.com/antonisgkamitsios/simple-todo-app/view"
)

func (app *application) handleTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.models.Todos.GetAll()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, view.Todos(todos))
}
