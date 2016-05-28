package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	//_ "code.google.com/p/gosqlite/sqlite3"
	//_ "github.com/mxk/go-sqlite/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type TimeStamp struct{ *time.Time }

func (t TimeStamp) Scan(value interface{}) error {
	fmt.Printf("%T\n", value)
	//...
	return nil
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("cannot open an SQLite memory database: %v", err)
	}
	defer db.Close()

	// sqlite> select strftime('%J', '2015-04-13T19:22:19.773Z'), strftime('%J', '2015-04-13T19:22:19');
	_, err = db.Exec("CREATE TABLE unix_time (time datetime); INSERT INTO unix_time (time) VALUES (strftime('%Y-%m-%dT%H:%MZ','now'))")
	if err != nil {
		log.Fatalf("cannot create schema: %v", err)
	}

	row := db.QueryRow("SELECT time FROM unix_time")
	var t time.Time
	err = row.Scan(TimeStamp{&t})
	if err != nil {
		log.Fatalf("cannot scan time: %v", err)
	}
	fmt.Println(t)
}
