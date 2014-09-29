////////////////////////////////////////////////////////////////////////////
// Porgram: DbLatching
// Purpose: Collect DB latching info into performance db
// Authors: Tong Sun (c) 2013, All rights reserved
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

const progname = "DbLatching"

const tbLatchDetails = "LatchDetails"
const tbLatchSummary = "LatchSummary"

var (
	fServer = flag.String("s", "",
		"server to garther DB latching info from\n\tDefault: no default. Must specify.\n")

	fPerfDb = flag.String("d", "perfdb",
		"performance db to hold the DB latching info\n\tDefault: perfdb\n")

	// Duration flags accept any input valid for time.ParseDuration
	fInterval = flag.Duration("i", 150*time.Second,
		"Interval to query for DB latching info\n\tDefault: 2.5m (2.5 minutes)\n")

	fLast = flag.Duration("l", 7200*time.Second,
		"Last, how long the recording will last\n\tDefault: 2hrs\n")
)

func usage() {
	// Fprintf allows us to print to a specifed file handle or stream
	fmt.Fprintf(os.Stderr, "\nUsage:\n %s -s DbServer [flags...]\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n-i: specify with a string, a possibly sequence of decimal numbers,\n each with optional fraction and a unit suffix, such as '300ms', '1.5h'\n or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.\n")
	fmt.Fprintf(os.Stderr, "\nDB latching info is recorded in table '%s' and '%s'.\n",
		tbLatchDetails, tbLatchSummary)
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

	for {

		// Get current timestamp
		_, err = conn.Exec("SELECT CURRENT_TIMESTAMP AS TIMESTAMP INTO ##CT")
		if err != nil {
			log.Println(progname+" Exec Error", err)
			return
		}

		// Garther DB latching info
		_, err = conn.Exec("SELECT 'CurrTest' CurrTest, TIMESTAMP CurrTime, L.*" +
			"  INTO ##Latch" +
			"  FROM ##CT, (" +
			"  SELECT spid, kpid, blocked, waittype, waittime, lastwaittype, waitresource," +
			"       uid, cpu, physical_io, memusage, login_time, last_batch, ecid, open_tran, status," +
			"       sid, hostname, program_name, hostprocess, cmd, nt_domain, nt_username, net_address," +
			"       net_library, loginame, context_info, sql_handle, stmt_start, stmt_end, request_id" +
			"       from master..sysprocesses sp cross apply fn_get_sql(sql_handle)" +
			"       WHERE status <> 'sleeping' AND waitresource = '2:1:103'" +
			"       ) L")
		if err != nil {
			log.Println(progname+" Exec Error", err)
			return
		}
		//log.Println("DB latching info garthered")

		// LatchDetails
		/*
			INSERT INTO perfdb..LatchDetails
			SELECT * FROM #Latch
		*/
		_, err = conn.Exec("INSERT INTO " + *fPerfDb + ".." + tbLatchDetails +
			" SELECT * FROM ##Latch")
		if err != nil {
			log.Println(progname+" Exec Error for "+tbLatchDetails, err)
			return
		}

		// LatchSummary
		_, err = conn.Exec("SELECT CurrTest, CurrTime TimeAbs, CurrTime TimeRel, Counts " +
			"  INTO ##Summary" +
			"  FROM ##CT, (" +
			"    SELECT CurrTest, CurrTime, COUNT(*) Counts FROM ##Latch GROUP BY CurrTest, CurrTime" +
			"  ) Agg")
		if err != nil {
			log.Println(progname+" Exec Error for "+tbLatchSummary, err)
			return
		}
		/*
			INSERT INTO perfdb..LatchSummary
			SELECT * FROM #Summary
		*/
		_, err = conn.Exec("INSERT INTO " + *fPerfDb + ".." + tbLatchSummary +
			" SELECT * FROM ##Summary")
		if err != nil {
			log.Println(progname+" Exec Error", err)
			return
		}
		//log.Println("DB latching info stored")

		row, err := conn.Query("SELECT CONVERT(char,TimeAbs,114) Time, Counts FROM ##Summary")
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
				log.Printf(progname+": DB latches totals to %03d", counts)
			}
		}
		if c == 0 {
			log.Println(progname + ": Zero DB latching found")

			// put a '0' record in the tbLatchSummary
			_, err = conn.Exec("INSERT INTO " + *fPerfDb + ".." + tbLatchSummary +
				" SELECT 'CurrTest' CurrTest, TIMESTAMP TimeAbs, TIMESTAMP TimeRel, 0 Counts " +
				"  FROM ##CT")
			if err != nil {
				log.Println(progname+" Exec Error for "+tbLatchSummary, err)
				return
			}

		}

		// Clean up
		_, err = conn.Exec("DROP TABLE ##CT; DROP TABLE ##Latch; DROP TABLE ##Summary")
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
