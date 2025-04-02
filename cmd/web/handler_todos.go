package main

import (
	"net/http"
	"strconv"

	"github.com/antonisgkamitsios/simple-todo-app/internal/data"
	"github.com/antonisgkamitsios/simple-todo-app/internal/flash"
	"github.com/antonisgkamitsios/simple-todo-app/view"
	"github.com/antonisgkamitsios/simple-todo-app/view/components"
	"github.com/julienschmidt/httprouter"
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

	if err != nil {
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

	r = flash.ContextSetFlash(r, flash.FlashTypeSuccess, "The todo was successfully added!")

	app.render(w, r, http.StatusCreated, view.TodoFormWithTodo(data.TodoForm{}, *todo))
}

func (app *application) handleTodosDelete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		app.clientError(w, r, http.StatusNotFound, "The todo could not be found", false)
		return
	}

	err = app.models.Todos.Delete(id)
	if err != nil {
		app.clientError(w, r, http.StatusNotFound, "The todo could not be found", false)
		return
	}

	r = flash.ContextSetFlash(r, flash.FlashTypeSuccess, "Todo was successfully deleted!")
	app.render(w, r, http.StatusOK, components.Flash())
}

func (app *application) handleTodosEdit(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		app.clientError(w, r, http.StatusNotFound, "The todo could not be found", false)
		return
	}

	todo, err := app.models.Todos.GetByID(id)
	if err != nil {
		app.clientError(w, r, http.StatusNotFound, "The todo could not be found", false)
		return
	}

	todoForm := data.TodoForm{
		Title: todo.Title,
	}

	app.render(w, r, http.StatusOK, view.TodoFormEdit(todoForm, todo.ID))

}

func (app *application) handleTodosUpdate(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		app.clientError(w, r, http.StatusNotFound, "The todo could not be found", false)
		return
	}

	todo, err := app.models.Todos.GetByID(id)
	if err != nil {
		app.clientError(w, r, http.StatusNotFound, "The todo could not be found", false)
		return
	}

	var input data.TodoForm

	err = app.decodePostForm(r, &input)
	if err != nil {
		app.clientError(w, r, http.StatusBadRequest, "Something went wrong with processing the form, check the data and try again", false)
		return
	}

	todo.Title = input.Title

	if data.ValidateTodo(todo, &input.Validator); !input.Validator.Valid() {
		app.render(w, r, http.StatusUnprocessableEntity, view.TodoFormEdit(input, todo.ID))
		return
	}

	err = app.models.Todos.Update(todo)
	if err != nil {
		app.serverError(w, r, err, false)
		return
	}

	app.render(w, r, http.StatusOK, view.TodoCard(*todo, false))
}

func (app *application) handleGetTodo(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		app.clientError(w, r, http.StatusNotFound, "The todo could not be found", false)
		return
	}

	todo, err := app.models.Todos.GetByID(id)
	if err != nil {
		app.clientError(w, r, http.StatusNotFound, "The todo could not be found", false)
		return
	}

	app.render(w, r, http.StatusOK, view.TodoCard(*todo, false))
}
