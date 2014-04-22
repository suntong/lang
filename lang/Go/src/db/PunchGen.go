////////////////////////////////////////////////////////////////////////////
// Porgram: PunchGen
// Purpose: Generate punch xml for RawPunchImport
// Authors: Tong Sun (c) 2014, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

package main

import (
	"bitbucket.org/kardianos/table"
	_ "code.google.com/p/odbc"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const progname = "PunchGen" // os.Args[0]

var (
	fSqlConnectionString = flag.String("c", "",
		"ConnectionString of Go MSSQL Odbc to MS SQL Server\n\tDefault: empty, which means using the -cs -cd setting. Sample: \n"+
			"  'driver=sql server;server=(local);database=LoadTest2010;uid=user;pwd=pass'\n")

	fServer = flag.String("cs", "",
		"Connection Server, server to get employee BadgeId info from\n\tDefault: no default. Must specify.\n")

	fDb = flag.String("cd", "",
		"Connection DB, db that holds the employee BadgeId info\n\tDefault: no default. Must specify.\n")

	fDeviceName = flag.String("dn", "",
		"DeviceName for RawPunchImport\n\tDefault: no default. Must specify.\n")

	fCodeIn = flag.Int("ci", 1,
		"Code of the punch-in event\n\tDefault: 1\n")

	fCodeOut = flag.Int("co", 6,
		"Code of the punch-out event\n\tDefault: 6\n")

	fStartDate = flag.String("sd", "",
		"Start date for punch generation\n\tDefault: no default. Must specify\n")

	fDurGen = flag.Duration("dg", 30*24*time.Hour,
		"Duration for punch generation \n\tDefault: 30 days\n")

	fStartTime = flag.String("st", "08:00:00",
		"Start time for punch generation\n\tDefault: 8am\n")

	fDurShift = flag.Duration("ds", 5*time.Hour,
		"Duration of the shift for the punch\n\tDefault: 5h\n")

	fEmpFilter = flag.String("f", "1=1",
		"Filter for selecting employees for RawPunchImport\n\tDefault: all employees\n")
)

const badgeIdSql1 = "SELECT e.XRefCode BadgeId FROM Employee e, EmploymentStatus s WHERE s.EmploymentStatusId = 4 AND s.ClientId = e.ClientId AND "
const badgeIdSql2 = " ORDER BY e.XRefCode"

const formDate = "2006-01-02"
const formTime = "15:04:05"

const durOneDay = 24 * time.Hour

func main() {
	flag.Usage = usage
	flag.Parse()

	// The cs, cd,dn & sd are mandatory flag
	// fmt.Fprintln(os.Stderr, progname, len(flag.Args()))
	if len(*fServer)*len(*fDb)*len(*fDeviceName)*len(*fStartDate) == 0 {
		if len(os.Args) > 1 {
			fmt.Fprintf(os.Stderr, "\n[%s] Error: not enough mandatory flag specified.\n",
				progname)
		}
		usage()
	}

	startDate, _ := time.Parse(formDate, *fStartDate)
	startTime, _ := time.Parse(formDate+formTime, *fStartDate+*fStartTime)
	durStartTime := startTime.Sub(startDate)
	_ = durStartTime

	// Construct the Go MSSQL odbc SqlConnectionString
	// https://code.google.com/p/odbc/source/browse/mssql_test.go
	var c string
	if *fSqlConnectionString == "" {
		var params map[string]string
		params = map[string]string{
			"driver":             "sql server",
			"server":             *fServer,
			"database":           *fDb,
			"trusted_connection": "yes",
		}

		for n, v := range params {
			c += n + "=" + v + ";"
		}

		log.Printf("[%s] Generating punches from %s@%s ",
			progname, *fDb, *fServer)

	} else {
		c = *fSqlConnectionString
		log.Printf("[%s] Generating punches from ConnectionString %s ",
			progname, *fSqlConnectionString)
	}
	// log.Println("Connection string: " + c)
	_startTime, _ := time.Parse(formTime, *fStartTime)

	conn, err := sql.Open("odbc", c)
	if err != nil {
		fmt.Println("Connecting Error")
		return
	}
	defer conn.Close()

	sql := badgeIdSql1 + *fEmpFilter + badgeIdSql2
	//log.Printf("[%s] SQL: %s ", progname, sql)

	badges, err := table.Get(conn, sql)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[%s] Generating punches for %d people ... ",
		progname, len(badges.Rows))
	// log.Printf("[%s]  starting\n\t\t\t\t from %s\n\t\t\t\t to   %s",
	// progname, startTime.Format(formDate+"T"+formTime),
	// startTime.Add(*fDurGen).Format(formDate+"T"+formTime))
	log.Printf("[%s]  starting from %s to %s\n\t\t\t\t\t from  %s  to  %s",
		progname, startTime.Format(formDate),
		startTime.Add(*fDurGen).Format(formDate),
		_startTime.Format(formTime),
		_startTime.Add(*fDurShift).Format(formTime))

	fmt.Println("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n<RawPunchImport>")
	for genDate := startDate; genDate.Before(startTime.Add(*fDurGen)); genDate = genDate.Add(durOneDay) {
		log.Printf("[%s]  generating for %s", progname, genDate.Format(formDate))

		for _, badge := range badges.Rows {
			badgeId := badge.MustGet("BadgeId").(string)
			// punch-in
			fmt.Printf("\t<RawPunch>\n\t\t<BadgeId>%s</BadgeId>"+
				"\n\t\t<DeviceName>%s</DeviceName>"+
				"\n\t\t<Time>%s</Time>"+
				"\n\t\t<Type>%d</Type>\n\t</RawPunch>\n",
				badgeId, *fDeviceName,
				startTime.Format(formDate+"T"+formTime+".000"), *fCodeIn)
			// punch-out
			fmt.Printf("\t<RawPunch>\n\t\t<BadgeId>%s</BadgeId>"+
				"\n\t\t<DeviceName>%s</DeviceName>"+
				"\n\t\t<Time>%s</Time>"+
				"\n\t\t<Type>%d</Type>\n\t</RawPunch>\n",
				badgeId, *fDeviceName,
				startTime.Add(*fDurShift).Format(formDate+"T"+formTime+".000"),
				*fCodeOut)
		}
	}
	fmt.Println("</RawPunchImport>")

	log.Printf("[%s] RawPunchXml generation finished correctly.\n", progname)
	return
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage:\n\n %s [flags...] \n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	os.Exit(0)
}
