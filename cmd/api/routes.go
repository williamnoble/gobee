package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.responseNotFound)
	router.MethodNotAllowed = http.HandlerFunc(app.responseMethodNotAllowed)
	return router

}
