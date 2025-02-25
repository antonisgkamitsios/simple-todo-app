package main

import (
	"net/http"

	"github.com/antonisgkamitsios/simple-todo-app/view" // new import
)

func (app *application) handleTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.models.Todos.GetAll()

	if err != nil {
		// todo: handle error
		return
	}

	app.render(w, r, http.StatusOK, view.Todos(todos))
}
