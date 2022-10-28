
> The `cmd` directory will contain the application-specific code for the executable applications in the project. For now we’ll have just one executable application — the web application — which will live under the cmd/web directory.


> The `internal` directory will contain the ancillary non-application-specific code used in the project. We’ll use it to hold potentially reusable code like validation helpers and the SQL database models for the project.

> The `ui` directory will contain the user-interface assets used by the web application. Specifically, the ui/html directory will contain HTML templates, and the ui/static directory will contain static files (like CSS and images).

> It’s important to point out that the directory name `internal` carries a special meaning and behavior in Go: **any packages which live under this directory can only be imported by code**
inside the parent of the internal directory. In our case, this means that any packages which live in internal can only be imported by code inside our snippetbox project directory.


> http 定义了 ServeHTTP(w http.ResponseWriter, r *http.Request) 所以任何对象只要实现这个方法即可 `mux.Handle("/", &home{})`
> log.Fatal() 输出 message 并执行 os.Exit(1)

# SQL
1. 安装驱动
2. 连接数据库 
3. 使用
4. 事务
5. 防止 sql 注入

```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
),
(   'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
),
(
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```


```golang
type ExampleModel struct { 
    DB *sql.DB
}
func (m *ExampleModel) ExampleTransaction() error {
        // Calling the Begin() method on the connection pool creates a new sql.Tx // object, which represents the in-progress database transaction.
        tx, err := m.DB.Begin()
        if err != nil {
            return err 
        }
        // Defer a call to tx.Rollback() to ensure it is always called before the // function returns. If the transaction succeeds it will be already be
        // committed by the time tx.Rollback() is called, making tx.Rollback() a // no-op. Otherwise, in the event of an error, tx.Rollback() will rollback // the changes before the function returns.
        defer tx.Rollback()
        // Call Exec() on the transaction, passing in your statement and any
        // parameters. It's important to notice that tx.Exec() is called on the
        // transaction object just created, NOT the connection pool. Although we're // using tx.Exec() here you can also use tx.Query() and tx.QueryRow() in
        // exactly the same way.
        _, err = tx.Exec("INSERT INTO ...")
        if err != nil {
            return err 
        }
        // Carry out another transaction in exactly the same way.
        _, err = tx.Exec("UPDATE ...") 
        if err != nil {
            return err 
        }
        // If there are no errors, the statements in the transaction can be committed // to the database with the tx.Commit() method.
        err = tx.Commit()
        return err
    }

```