package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// 处理函数，用于接收请求和发送结果的函数
func home(w http.ResponseWriter, req *http.Request) {
	if (req.URL.Path != "/") {
		http.NotFound(w, req)
		return 
	}
	w.Write([]byte("Hello"))
}

func view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	// w.Write([]byte("view"))
	fmt.Fprintf(w, "view %d....\n", id)
}

func create(w http.ResponseWriter, r *http.Request) {
	// http 的常量 ： MethodPost 
	if (r.Method != http.MethodPost) {
		w.Header().Set("Allow", "POST")
		// 可以直接返回错误
		http.Error(w, "Method Not Allow", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("create"))
}
