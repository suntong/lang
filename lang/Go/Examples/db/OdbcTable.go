////////////////////////////////////////////////////////////////////////////
// Porgram: OdbcTable
// Purpose: Go MSSQL odbc demo, from SQL to table
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// http://godoc.org/bitbucket.org/kardianos/table

// Style: gofmt -tabs=false -tabwidth=2 -w

package main

import (
  "bitbucket.org/kardianos/table"
  _ "code.google.com/p/odbc"
  "database/sql"
  "fmt"
  "log"
  "os"
  "time"
)

func main() {
  conn, err := sql.Open("odbc",
    "driver=sql server;server=localhost;database=tempdb;trusted_connection=yes;")
  if err != nil {
    fmt.Println("Connecting Error")
    return
  }
  defer conn.Close()

  table, err := table.Get(conn, "select top 5 database_id, name, create_date from sys.databases WHERE database_id >= ?", 1)
  if err != nil {
    log.Fatal(err)
  }

  dumpTable(table)
  
  testTVF(conn)
  testExec(conn)
  
  fmt.Fprintf(os.Stderr, "\nFinished correctly\n")
  return
}

/*

testTVF: to test Table-Valued Functions of MS SQL Server

Need to define the following in tempdb:

	Create Function Instances (@Bottom int)
	Returns Table     
	As
	Return
	(
	select top 5 database_id, name, create_date from sys.databases WHERE database_id >= @Bottom
	)

Here is how to run:

	Select * from Instances(5)

*/
func testTVF(conn *sql.DB) {
  table, err := table.Get(conn, "Select * from Instances(?)", 5)
  if err != nil {
    log.Fatal(err)
  }

  dumpTable(table)

}

func testExec(conn *sql.DB) {
  table, err := table.Get(conn, "EXEC sp_executesql N'Select * from Instances(5)'")
  if err != nil {
    log.Fatal(err)
  }

  dumpTable(table)

}

func dumpTable(table *table.Buffer) {
  for i, element := range table.ColumnName {
    if i != 0 {
      fmt.Printf(",")
    }
    fmt.Printf("%s", element)
  }
  fmt.Printf("\n")

  for _, row := range table.Rows {
    for i, colname := range table.ColumnName {
      if i != 0 {
        fmt.Printf(",")
      }
      switch x := row.MustGet(colname).(type) {
      case string: // x is a string
        fmt.Printf("\"%s\"", x)
      case int: // now x is an int
        fmt.Printf("%d", x)
      case int32: // now x is an int32
        fmt.Printf("%d", x)
      case int64: // now x is an int64
        fmt.Printf("%d", x)
      case float32: // now x is an float32
        fmt.Printf("%f", x)
      case float64: // now x is an float64
        fmt.Printf("%f", x)
      case time.Time: // now x is a time.Time
        fmt.Printf("%s", x.Format(time.RFC3339))
      default:
        fmt.Printf("%s", x)
      }
    }
    fmt.Printf("\n")
  }
}
