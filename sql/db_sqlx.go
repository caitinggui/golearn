package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//var db *sqlx.DB

	db, err := sqlx.Open("sqlite3", ":memory:")
	//db, err := sqlx.Open("mysql", "root:hhlyctgllp@localhost")
	err = db.Ping()
	fmt.Println(db, err)
}
