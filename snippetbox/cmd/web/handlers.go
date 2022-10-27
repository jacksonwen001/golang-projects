package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// 处理函数，用于接收请求和发送结果的函数
func (app *application) home(w http.ResponseWriter, req *http.Request) {
	if (req.URL.Path != "/") {
		http.NotFound(w, req)
		return 
	}
	files := []string {"./ui/html/base.tmpl", "./ui/html/partials/nav.tmpl", "./ui/html/pages/home.tmpl"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	// w.Write([]byte("view"))
	fmt.Fprintf(w, "view %d....\n", id)
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {
	// http 的常量 ： MethodPost 
	if (r.Method != http.MethodPost) {
		w.Header().Set("Allow", "POST")
		// 可以直接返回错误
		http.Error(w, "Method Not Allow", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("create"))
}
