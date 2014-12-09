////////////////////////////////////////////////////////////////////////////
// Porgram: DbLocking
// Purpose: Collect DB locking info into performance db
// Authors: Tong Sun (c) 2014, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

package main

import (
	_ "code.google.com/p/odbc"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const progname = "DbLocking"

const tbDbLocking = "Locking"

// time layout
const timelayout = "2006/01/02 03:04:05 "

var (
	fServer = flag.String("s", "",
		"server to garther DB locking info from\n\tDefault: no default. Must specify.\n")

	fPerfDb = flag.String("d", "perfdb",
		"performance db to hold the DB locking info\n\tDefault: perfdb\n")

	// Duration flags accept any input valid for time.ParseDuration
	fInterval = flag.Duration("i", 5*time.Second,
		"Interval to query for DB locking info\n\tDefault: 5sec\n")

	fLast = flag.Duration("l", 7200*time.Second,
		"Last, how long the recording will last\n\tDefault: 2hrs\n")
)

func usage() {
	// Fprintf allows us to print to a specifed file handle or stream
	fmt.Fprintf(os.Stderr, "\nUsage:\n %s -s DbServer [flags...]\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n-i: specify with a string, a possibly sequence of decimal numbers,\n each with optional fraction and a unit suffix, such as '300ms', '1.5h'\n or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.\n")
	fmt.Fprintf(os.Stderr, "\nDB locking info is recorded in table '%s'.\n",
		tbDbLocking)
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(*fServer) < 1 {
		// fmt.Fprintf(os.Stderr, "\nError:\n Flag -s is mandatory.\n\n")
		usage()
	}

	// Construct the Go MSSQL odbc SqlConnectionString
	var params map[string]string
	params = map[string]string{
		"driver":             "sql server",
		"server":             *fServer,
		"database":           *fPerfDb,
		"trusted_connection": "yes",
	}

	var c string
	for n, v := range params {
		c += n + "=" + v + ";"
	}
	//log.Println(c)

	// establish the Sql Server connection
	conn, err := sql.Open("odbc", c)
	if err != nil {
		log.Println(progname+" Connecting Error", err)
		return
	}
	defer conn.Close()

	log.Printf(progname+": Program starts with\n\t checking interval of %v\n"+
		"\t run-time duration of %v", *fInterval, *fLast)


	startTime := time.Now()

	// SELECT REPLICATE('-',96) CurrTest, ...
	for {

		// Get Lock Summary
		// select GETDATE(), lastwaittype, LEFT(waitresource, 18), COUNT(1) from master..sysprocesses sp cross apply fn_get_sql(sql_handle) where lastwaittype LIKE 'LCK%' GROUP BY lastwaittype, LEFT(waitresource, 18)
		_, err = conn.Exec("SELECT 'CurrTest' CurrTest, CURRENT_TIMESTAMP CurrTime, " +
			"LastWaitType, LEFT(waitresource, 18) Resource, COUNT(1) CNT " +
			"INTO ##Lock " +
			"FROM master..sysprocesses sp cross apply fn_get_sql(sql_handle) " +
			"WHERE lastwaittype LIKE 'LCK%' " +
			"GROUP BY lastwaittype, LEFT(waitresource, 18)")
		if err != nil {
			log.Println(progname+" Exec Error", err)
			return
		}
		//log.Println("DB locking info retrieved\n")

		_, err = conn.Exec("INSERT INTO " + *fPerfDb + ".." + tbDbLocking +
			" (CurrTest, CurrTime, LastWaitType, Resource, CNT) SELECT * FROM ##Lock")
		if err != nil {
			log.Println(progname+" Exec Error for "+tbDbLocking, err)
			return
		}
		//log.Println("DB locking info stored\n")

		row, err := conn.Query("SELECT CONVERT(char,CurrTime,114) Time, COUNT(*) Counts FROM ##Lock GROUP BY CurrTime")
		if err != nil {
			log.Println(progname+" Query Error", err)
			return
		}
		defer row.Close()

		// Show
		c := 0
		for row.Next() {
			c++
			//var id int
			var (
				time   string
				counts int
			)
			if err := row.Scan(&time, &counts); err == nil {
				log.Printf(progname+": DB locks total to %03d", counts)
			}
		}
		if c == 0 {
			fmt.Printf(time.Now().Format(timelayout) + progname +
				": Zero DB locking found\r")
		}

		// Clean up
		_, err = conn.Exec("DROP TABLE ##Lock")
		if err != nil {
			log.Println(progname+" Cleanup Exec Error", err)
			return
		}

		//break out of the for loop when due time passed
		elapsed := time.Since(startTime)
		if elapsed >= *fLast { 
			log.Println(progname+": Run-time duration reached, exiting")
			break
		}

		// delay the given interval
		time.Sleep(*fInterval)

	}
	return
}
