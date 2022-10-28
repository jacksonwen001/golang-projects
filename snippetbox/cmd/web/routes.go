package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServce := http.FileServer(http.Dir("./ui/static"))
	// 注册路由
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.view)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.create)
	router.Handler(http.MethodGet, "/static/*filepath", fileServce)
	return alice.New(app.recoverPanic, app.logRequest, secureHeaders).Then(router)
}
