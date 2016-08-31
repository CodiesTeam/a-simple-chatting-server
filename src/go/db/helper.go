package db

import (
    "database/sql"
    "fmt"
    "os"
	_ "github.com/go-sql-driver/mysql"
)

var (
    username string
    password string
    host string
    port string
    database string

    db *sql.DB
)

func config() {
    // initialized configurations
    username = os.Getenv("MYSQL_USERNAME")
    password = os.Getenv("MYSQL_PASSWORD")
    host = os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
    if host == "" {
        host = "localhost"
    }

    port = os.Getenv("MYSQL_PORT_3306_TCP_PORT")
    if port == "" {
        port = "3306"
    }

    database = os.Getenv("MYSQL_INSTANCE_NAME")
}

func connetDB() error {
    // connect database with configurations
    uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
    if db, err := sql.Open("mysql", uri); err != nil {
        panic(err)
    }
    return
}

func init() {
	config()
    connetDB()
}