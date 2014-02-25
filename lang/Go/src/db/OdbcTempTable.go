////////////////////////////////////////////////////////////////////////////
// Porgram: OdbcTempTable
// Purpose: Test temp table access with Go MSSQL odbc
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

Notice:

Contrary to its definition,
http://www.codeproject.com/Articles/42553/Quick-Overview-Temporary-Tables-in-SQL-Server-2005
http://stackoverflow.com/questions/2920836/local-and-global-temporary-table-in-sql-server

  Local temporary tables are visible only to their creators during the same
  connection to an instance of SQL Server as when the tables were first
  created or referenced. Local temporary tables are deleted after the user
  disconnects from the instance of SQL Server. Global temporary tables are
  visible to any user and any connection after they are created, and are
  deleted when all users that are referencing the table disconnect from the
  instance of SQL Server.

Even with the same ODBC connection, temp table (in form of #TempTable) cannot be used.
Else will get the following error during the "run query statement" phase:

  Query Error SQLExecute: {42S02} [Microsoft][ODBC SQL Server Driver][SQL Server]Invalid object name '#CT'.

Have to use global temp table instead (in form of ##TempTable)

The reason: Exec will release the connection right before returning.
Ref, http://jmoiron.net/blog/gos-database-sql/

*/

func main() {
	conn, err := sql.Open("odbc",
		"driver=sql server;server=localhost;database=tempdb;trusted_connection=yes;")
	if err != nil {
		fmt.Println("Connecting Error")
		return
	}
	defer conn.Close()

	log.Printf("Prepare data")
	// Use db.Exec() to executes a query without returning any rows.
	_, err = conn.Exec("SELECT CURRENT_TIMESTAMP AS TIMESTAMP INTO ##CT")
	if err != nil {
		fmt.Println("Exec Error", err)
		return
	}

	// Preparing Queries
	/*
	   You should, in general, always prepare queries to be used multiple times.
	   The result of preparing the query is a prepared statement, which can
	   have ? placeholders for parameters that you'll provide when you execute
	   the statement. This is much better than concatenating strings.
	*/

	log.Printf("Prepare query statement\n")
	stmt, err := conn.Prepare("SELECT convert(varchar, TIMESTAMP, 121) FROM ##CT")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Run query statement\n")
	// Use db.Query() to send the query to the database. Check errors as usual.
	row, err := stmt.Query()
	if err != nil {
		fmt.Println("Query Error", err)
		return
	}
	defer row.Close()

	// Iterate over the row with row.Next()
	// and read the columns in each row into variables with row.Scan()
	fmt.Printf("\nQuery result:\n")
	for row.Next() {
		var timestamp string
		if err := row.Scan(&timestamp); err == nil {
			fmt.Println(timestamp)
		}
	}
	// Check for errors after done iterating over the row. Should always do.
	err = row.Err()
	if err != nil {
		fmt.Printf("\nFatal: %s\n", err)
	}

	fmt.Printf("\nFinished correctly\n")
	return
}
