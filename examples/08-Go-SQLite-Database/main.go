package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT);
    `
    db.Exec(sqlStmt)

    db.Exec("INSERT INTO users(name) VALUES (?)", "Alice")

    rows, _ := db.Query("SELECT id, name FROM users")
    defer rows.Close()
    for rows.Next() {
        var id int
        var name string
        rows.Scan(&id, &name)
        fmt.Println(id, name)
    }
}
