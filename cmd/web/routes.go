package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	router.HandlerFunc(http.MethodGet, "/", app.handleHome)

	router.HandlerFunc(http.MethodGet, "/todos", app.handleTodos)

	return app.comesFromHTMX(router)
}
