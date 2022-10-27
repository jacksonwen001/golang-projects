package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// 路由的作用

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERRPR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
	}

	// 启动服务器
	srv := &http.Server{
		Addr:     "localhost:4000",
		ErrorLog: errLog,
		Handler:  app.routes(),
	}
	db, err := openDB("root:root@/snippetbox?parseTime=true") // "root:root@tcp(localhost:3306)/snippetbox?parseTime=true"
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	infoLog.Print("Starting server in 4000....")

	err = srv.ListenAndServe()
	if err != nil {
		// 这个函数的作用是输出错误，并且直接结束程序退出
		errLog.Fatal(err)
	}
}

func openDB(s string) (*sql.DB, error) {
	db, err := sql.Open("mysql", s)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

