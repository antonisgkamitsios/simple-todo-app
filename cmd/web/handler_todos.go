package main

import (
	"net/http"

	"github.com/antonisgkamitsios/simple-todo-app/internal/data"
	"github.com/antonisgkamitsios/simple-todo-app/view"
)

func (app *application) handleTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.models.Todos.GetAll()

	if err != nil {
		app.serverError(w, r, err, true)
		return
	}

	app.render(w, r, http.StatusOK, view.Todos(todos))
}

func (app *application) handleTodosNew(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, view.TodoForm(data.TodoForm{}))
}

func (app *application) handleTodosCreate(w http.ResponseWriter, r *http.Request) {
	var input data.TodoForm

	err := app.decodePostForm(r, &input)

	if err != nil || true {
		app.clientError(w, r, http.StatusBadRequest, "Something went wrong with processing the form, check the data and try again", false)
		return
	}

	todo := &data.Todo{Title: input.Title}
	if data.ValidateTodo(todo, &input.Validator); !input.Validator.Valid() {
		app.render(w, r, http.StatusUnprocessableEntity, view.TodoForm(input))
		return
	}

	err = app.models.Todos.Insert(todo)
	if err != nil {
		app.serverError(w, r, err, false)
		return
	}

	app.render(w, r, http.StatusCreated, view.TodoFormWithTodo(data.TodoForm{}, *todo))
}
