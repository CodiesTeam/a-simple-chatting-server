package db

import (
    "database/sql"
    "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDB(name string) {
    db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    _, err = db.Exec("CREATE DATABASE " + name)
    if err != nil {
        panic(err)
    }

    _, err = db.Exec("USE " + name)
    if err != nil {
        panic(err)
    }

    _, err = db.Exec("CREATE TABLE test ( id integer, name varchar(32) )")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Create db successfully!")
}