////////////////////////////////////////////////////////////////////////////
// Porgram: OdbcDemo
// Purpose: Go MSSQL odbc demo, using the code.google.com/p/odbc driver 
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

package main

import (
  _ "code.google.com/p/odbc"
  "database/sql"
  "fmt"
  "log"
)

/*
Notice: Load the driver anonymously, aliasing its package qualifier to _
so none of its exported names are visible to our code. Under the hood,
the driver registers itself as being available to the database/sql package,
but in general nothing else happens.
-- https://github.com/VividCortex/go-database-sql-tutorial
*/

func main() {
  conn, err := sql.Open("odbc",
    "driver=sql server;server=localhost;database=tempdb;trusted_connection=yes;")
  if err != nil {
    fmt.Println("Connecting Error")
    return
  }
  defer conn.Close()

  stmt, err := conn.Prepare("select top 5 database_id, name from sys.databases WHERE database_id >= 5")
  if err != nil {
    fmt.Println("Query Error", err)
    return
  }
  defer stmt.Close()

  // Use db.Query() to send the query to the database. Check errors as usual.
  row, err := stmt.Query()
  if err != nil {
    fmt.Println("Query Error", err)
    return
  }
  defer row.Close()

  // Iterate over the row with row.Next()
  // and read the columns in each row into variables with row.Scan()
  fmt.Printf("\nResult set 1:\n")
  for row.Next() {
    //var id int
    var (
      id   int
      name string
    )
    if err := row.Scan(&id, &name); err == nil {
      fmt.Println(id, name)
    }
  }
  // Check for errors after done iterating over the row. Should always do.
  err = row.Err()
  if err != nil {
    fmt.Printf("\nFatal: %s\n", err)
  }

  // Preparing Queries
  /*
  	You should, in general, always prepare queries to be used multiple times.
  	The result of preparing the query is a prepared statement, which can
  	have ? placeholders for parameters that you'll provide when you execute
  	the statement. This is much better than concatenating strings.
  */

  stmt, err = conn.Prepare("select top 5 database_id, name from sys.databases WHERE database_id >= ?")
  if err != nil {
    log.Fatal(err)
  }

  row, err = stmt.Query(1)
  if err != nil {
    log.Fatal(err)
  }
  defer row.Close()

  fmt.Printf("\nResult set 2:\n")
  for row.Next() {
    var (
      id   int
      name string
    )
    if err := row.Scan(&id, &name); err == nil {
      fmt.Println(id, name)
    } else {
      log.Fatal(err)
    }
  }
  err = row.Err()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("\nFinished correctly\n")
  return
}
