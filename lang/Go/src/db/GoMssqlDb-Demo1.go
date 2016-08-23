////////////////////////////////////////////////////////////////////////////
// Porgram: go-mssqldb demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://stackoverflow.com/questions/32010749/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb" // the underscore indicates the package is used
)

func main() {
	fmt.Println("starting app")

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

	// the user needs to be setup in SQL Server as an SQL Server user.
	// see create login and the create user SQL commands as well as the
	// SQL Server Management Studio documentation to turn on Hybrid Authentication
	// which allows both Windows Authentication and SQL Server Authentication.
	// also need to grant to the user the proper access permissions.
	// also need to enable TCP protocol in SQL Server Configuration Manager.
	//
	// you could also use Windows Authentication if you specify the fully qualified
	// user id which would specify the domain as well as the user id.
	// for instance you could specify "user id=domain\\user;password=userpw;".
	// condb, errdb := sql.Open("mssql", "server=localhost;user id=gouser;password=g0us3r;")
	condb, errdb := sql.Open("mssql", c)
	if errdb != nil {
		fmt.Println("  Error open db:", errdb.Error())
	}

	defer condb.Close()

	errdb = condb.Ping()
	if errdb != nil {
		log.Fatal(errdb)
	}

	// drop the database if it is there so we can recreate it
	// next we will recreate the database, put a table into it,
	// and add a few rows.
	_, errdb = condb.Exec("drop database mydbthing")
	if errdb != nil {
		fmt.Println("  Error Exec db: drop db - ", errdb.Error())
	}

	_, errdb = condb.Exec("create database mydbthing")
	if errdb != nil {
		fmt.Println("  Error Exec db: create db - ", errdb.Error())
	}

	_, errdb = condb.Exec("use  mydbthing")
	if errdb != nil {
		fmt.Println("  Error Exec db: using db - ", errdb.Error())
	}

	_, errdb = condb.Exec("create table junky (one int, two int)")
	if errdb != nil {
		fmt.Println("  Error Exec db: create table - ", errdb.Error())
	}

	_, errdb = condb.Exec("insert into junky (one, two) values (101, 201)")
	if errdb != nil {
		fmt.Println("  Error Exec db: insert table 1 - ", errdb.Error())
	}
	_, errdb = condb.Exec("insert into junky (one, two) values (102, 202)")
	if errdb != nil {
		fmt.Println("  Error Exec db: insert table 2 - ", errdb.Error())
	}
	_, errdb = condb.Exec("insert into junky (one, two) values (103, 203)")
	if errdb != nil {
		fmt.Println("  Error Exec db: insert table 3 - ", errdb.Error())
	}

	// Now that we have our database lets read some records and print them.
	var (
		one int
		two int
	)

	// documentation about a simple query and results loop is at URL
	// http://go-database-sql.org/retrieving.html
	// we use Query() and not Exec() as we expect zero or more rows to
	// be returned. only use Query() if rows may be returned.
	fmt.Println("  Query our table for the three rows we inserted.")
	rows, errdb := condb.Query("select one, two from junky")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one, &two)
		if err != nil {
			fmt.Println("  Error Query db: select - ", err.Error())
		} else {
			fmt.Printf("    - one %d and two %d\n", one, two)
		}
	}
	rows.Close()

	errdb = rows.Err()
	if errdb != nil {
		fmt.Println("  Error Query db: processing rows - ", errdb.Error())
	}

	fmt.Println("ending app")
}

/*

$ HOST=localhost SQLUSER=sa SQLPASSWORD=sa DATABASE=test go run GoMssqlDb-Demo1.go
starting app
  Error Exec db: drop db -  mssql: Cannot drop the database 'mydbthing', because it does not exist or you do not have permission.
  Query our table for the three rows we inserted.
    - one 101 and two 201
    - one 102 and two 202
    - one 103 and two 203
ending app

*/
