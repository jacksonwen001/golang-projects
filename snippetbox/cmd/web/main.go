package main

import (
	"log"
	"net/http"
)


func main() {
	// 路由的作用
	mux := http.NewServeMux()
	fileServce := http.FileServer(http.Dir("./ui/static"))

	// 注册路由 
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", view)
	mux.HandleFunc("/snippet/create", create)
	mux.Handle("/static", fileServce)

	log.Print("Starting server in 4000....")
	// 启动服务器
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		// 这个函数的作用是输出错误，并且直接结束程序退出
		log.Fatal(err)
	}
}
