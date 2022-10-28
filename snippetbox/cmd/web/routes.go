package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServce := http.FileServer(http.Dir("./ui/static"))
	dynamic := alice.New(app.session.LoadAndSave)
	// 注册路由
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.view))
	router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.create))
	router.Handler(http.MethodGet, "/static/*filepath", fileServce)
	return alice.New(app.recoverPanic, app.logRequest, secureHeaders).Then(router)
}
