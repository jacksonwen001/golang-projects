package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServce := http.FileServer(http.Dir("./ui/static"))
	// 注册路由
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.view)
	mux.HandleFunc("/snippet/create", app.create)
	mux.Handle("/static", fileServce)
	return mux
}