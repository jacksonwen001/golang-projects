package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"snippetbox.jackson.net/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippet  *models.SnippetModel
	session  *scs.SessionManager
}

func main() {
	db, err := openDB("root:root@/snippetbox?parseTime=true") // "root:root@tcp(localhost:3306)/snippetbox?parseTime=true"

	pool := redisConfigure()
	sessionManager := scs.New()
	sessionManager.Store = redisstore.New(pool)

	// 路由的作用

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERRPR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
		snippet:  &models.SnippetModel{DB: db},
		session:  sessionManager,
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

func redisConfigure() *redis.Pool {
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	return pool
}
