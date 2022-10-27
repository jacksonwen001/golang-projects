package main

import (
	"log"
	"net/http"
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
	w.Write([]byte("view"))
}

func create(w http.ResponseWriter, r *http.Request) {
	// http 的常量 ： MethodPost 
	if (r.Method != http.MethodPost) {
		w.Header().Set("Allow", "POST")
		// 可以直接返回错误
		http.Error(w, "Method Not Allow", 405)
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
