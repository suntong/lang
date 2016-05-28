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
	fmt.Printf("Value: %v\n", value.(time.Time))
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
	_, err = db.Exec("CREATE TABLE unix_time (time datetime)")
	if err != nil {
		log.Fatalf("cannot create schema: %v", err)
	}

	_, err = db.Exec("INSERT INTO unix_time (time) VALUES (strftime('%Y-%m-%dT%H:%MZ','now'))")
	if err != nil {
		log.Fatalf("cannot insert: %v", err)
	}

	row := db.QueryRow("SELECT time FROM unix_time")
	fmt.Printf("Rec: %v\nScan: ", *row)

	var t time.Time
	err = row.Scan(TimeStamp{&t})
	if err != nil {
		log.Fatalf("cannot scan time: %v", err)
	}
	fmt.Println(t)
}

/*

Rec: {<nil> 0xc82001e180}
Scan: time.Time
Value: 0001-01-01 00:00:00 +0000 UTC
0001-01-01 00:00:00 +0000 UTC

*/
