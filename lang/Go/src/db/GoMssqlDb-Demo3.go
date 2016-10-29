////////////////////////////////////////////////////////////////////////////
// Porgram: go-mssqldb demo
// Purpose: go-mssqldb transaction demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://play.golang.org/p/0L3Vgk8C_F
////////////////////////////////////////////////////////////////////////////

package main

import (
	"database/sql"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func main() {

	var params map[string]string
	params = map[string]string{
		"server":   os.Getenv("HOST"),
		"user id":  os.Getenv("SQLUSER"),
		"password": os.Getenv("SQLPASSWORD"),
		"database": os.Getenv("DATABASE"),
	}
	var c string
	for n, v := range params {
		c += n + "=" + v + ";"
	}

	err := Init(c)
	defer Close()
	if err != nil {
		panic(err)
	}

	tran, err := db.Begin()
	if err != nil {
		panic(err)
	}
	MustInsertNewTestItem(1, "one", tran)
	one := MustGetTestValue(1, tran)
	if one != "one" {
		panic(`one != "one"`)
	}
	MustInsertNewTestItem(2, "two", tran)
	MustUpdateTestItem(2, "two modified", tran)

	err = tran.Rollback()
	if err != nil {
		panic(err)
	}
}

func Init(connectionString string) error {
	conn, err := sql.Open("mssql", connectionString)
	if err == nil {
		db = conn
	}
	return err
}

func Close() error {
	return db.Close()
}

func MustInsertNewTestItem(id int32, value string, tran *sql.Tx) {
	_, err := tran.Exec("INSERT INTO TestTable (Id, Value) VALUES ($1, $2)", id, value)
	if err != nil {
		panic(err)
	}
}

func MustUpdateTestItem(id int32, value string, tran *sql.Tx) {
	_, err := tran.Exec("UPDATE TestTable SET Value = $2 WHERE Id = $1", id, value)
	if err != nil {
		panic(err)
	}
}

func MustGetTestValue(id int32, tran *sql.Tx) string {
	stmt, err := tran.Prepare("SELECT Value FROM TestTable WHERE Id = $1")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	value := ""
	err = row.Scan(&value)
	if err != nil {
		panic(err)
	}
	return value
}

/*

$ HOST=localhost SQLUSER=sa SQLPASSWORD=sa DATABASE=test go run GoMssqlDb-Demo1.go

*/
