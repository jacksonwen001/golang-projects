package main

import (
	"fmt"
	"log"
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

func main() {
	// 路由的作用
	mux := http.NewServeMux()
	// 注册路由 
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", view)
	mux.HandleFunc("/snippet/create", create)

	log.Print("Starting server in 4000....")
	// 启动服务器
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		// 这个函数的作用是输出错误，并且直接结束程序退出
		log.Fatal(err)
	}
}
