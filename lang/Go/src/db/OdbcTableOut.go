////////////////////////////////////////////////////////////////////////////
// Porgram: OdbcTableOut
// Purpose: To test go Stderr bufferring during lengthy process
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Go MSSQL odbc: http://godoc.org/bitbucket.org/kardianos/table

// Style: gofmt -tabs=false -tabwidth=4 -w

/*

Problem statement:

In a lengthy process, I output a dot every so often, and a CR at the end.

But it turns out that the dots all show up at once when the lengthy process finished,
not one by one during it.

Data preparation:

seq -w 120 | sed "s/^/\t, '11111...11111' AS C/"

-- run this first
SELECT ...
  INTO BigTable

-- then run this line 10~20 times
INSERT INTO BigTable
SELECT ...

Final result detailed in
http://pastebin.com/YxCij3Cs

Problem duplication:

There are two SQL statements there. Please run the first SQL statement first,
in tempdb on local machine, then run the second SQL statement 10~20 times,
as instructed in the code. And then run

  go run OdbcTableOut.go "" BigTable

supply the first parameter as the connection string if you are not
running the above in tempdb on local machine.

> In Go neither os.Stdout nor os.Stderr is buffered.

Then we should see the dots progressing. However, you will find,
in this simple example,  the dots all show up at once after the
lengthy process finished, not one by one during it.
I.e., "nothing" then all-of-sudden "everything".

*/

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
	connStr := "driver=sql server;server=(local);database=tempdb;trusted_connection=yes"
	if len(os.Args) > 1 && len(os.Args[1]) > 1 {
		connStr = os.Args[1]
	}
	//log.Printf("Connecting with '%s'\n", connStr)
	conn, err := sql.Open("odbc", connStr)
	if err != nil {
		fmt.Println("Connecting Error")
		return
	}
	defer conn.Close()

	connTbl := "sys.databases"
	if len(os.Args) > 2 {
		connTbl = os.Args[2]
	}
	table, err := table.Get(conn, "select * from "+connTbl)
	if err != nil {
		log.Fatal(err)
	}

	dumpTable(table)

	fmt.Fprintf(os.Stderr, "\nFinished correctly\n")
	return
}

func dumpTable(table *table.Buffer) {
	// open the output file
	file, err := os.Create("out.csv")
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// output header
	for i, element := range table.ColumnName {
		if i != 0 {
			fmt.Fprintf(file, ",")
		}
		fmt.Fprintf(file, "\"%s\"", element)
	}
	fmt.Fprintf(file, "\n")

	// output body
	const layout = "01/02/2006 15:04:05.999"
	for _, row := range table.Rows {
		for i, colname := range table.ColumnName {
			if i != 0 {
				fmt.Fprintf(file, ",")
			}
			switch x := row.MustGet(colname).(type) {
			case string: // x is a string
				fmt.Fprintf(file, "\"%s\"", x)
			case int: // now x is an int
				fmt.Fprintf(file, "\"%d\"", x)
			case int32: // now x is an int32
				fmt.Fprintf(file, "\"%d\"", x)
			case int64: // now x is an int64
				fmt.Fprintf(file, "\"%d\"", x)
			case float32: // now x is an float32
				fmt.Fprintf(file, "\"%f\"", x)
			case float64: // now x is an float64
				fmt.Fprintf(file, "\"%f\"", x)
			case time.Time: // now x is a time.Time
				fmt.Fprintf(file, "\"%s\"", x.Format(layout))
			default:
				fmt.Fprintf(file, "\"%s\"", x)
			}
		}
		fmt.Fprintf(file, "\n")
		fmt.Fprintf(os.Stderr, ".")
	}
	fmt.Fprintf(os.Stderr, "\n")

}
