package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// 路由的作用
	mux := http.NewServeMux()
	fileServce := http.FileServer(http.Dir("./ui/static"))
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERRPR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errLog,
		infoLog: infoLog,
	}
	// 注册路由
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.view)
	mux.HandleFunc("/snippet/create", app.create)
	mux.Handle("/static", fileServce)

	infoLog.Print("Starting server in 4000....")

	// 启动服务器
	srv := &http.Server{
		Addr:     "localhost:4000",
		ErrorLog: errLog,
		Handler:  mux,
	}
	err := srv.ListenAndServe()
	if err != nil {
		// 这个函数的作用是输出错误，并且直接结束程序退出
		errLog.Fatal(err)
	}
}
