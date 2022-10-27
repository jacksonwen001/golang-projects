package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.jackson.net/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippet  *models.SnippetModel
}

func main() {
	db, err := openDB("root:root@/snippetbox?parseTime=true") // "root:root@tcp(localhost:3306)/snippetbox?parseTime=true"

	// 路由的作用

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERRPR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
		snippet:  &models.SnippetModel{DB: db},
	}

	// 启动服务器
	srv := &http.Server{
		Addr:     "localhost:4000",
		ErrorLog: errLog,
		Handler:  app.routes(),
	}
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
