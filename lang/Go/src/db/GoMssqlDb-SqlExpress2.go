
// Go connection to Azure SQL Database
// https://docs.microsoft.com/en-us/learn/modules/azure-database-fundamentals/exercise-create-sql-database

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	//"errors"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	user = "sqluser"
	password = "<your_password>"
)

func main() {
	password = "Pa$$w0rd..."
	// Build connection string
	// connString := "server=localhost\\SQLExpress;port=1433;database=master"
	// WARN: You specified both instance name and port in the connection string, port will be used and instance name will be ignored
	// If without the port= and get:
	// unable to get instances from Sql Server Browser on host localhost: read udp 127.0.0.1:52494->127.0.0.1:1434: i/o timeout
	// --> Make sure that the "Sql Server Browser" service is up and running!
	// connString := "server=localhost\\SQLExpress;database=master"
	// connString := "server=localhost\\SQLExpress;user id=;password=;database=master"
	// login error: mssql: Login failed for user ''
	connString := fmt.Sprintf("server=localhost\\SQLExpress;user id=%s;password=%s;database=master", user, password)
	var err error
	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n\n")
	defer db.Close()

	var (
		f1, f2 string
	)
	rows, errdb := db.Query(`select @@servername ServerName, DB_NAME() DataBaseName`	)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&f1, &f2)
		if err != nil {
			fmt.Println("  Error Query db: select - ", err.Error())
		} else {
			fmt.Printf("- f1: %s\n  f2: %s\n", f1, f2)
		}
	}
	rows.Close()

	errdb = rows.Err()
	if errdb != nil {
		fmt.Println("  Error Query db: processing rows - ", errdb.Error())
	}
}


/* 

$ go run GoMssqlDb-DemoAzure.go
Connected!

- f1: Road Frames
  f2 HL Road Frame - Black, 58
- f1: Road Frames
  f2 HL Road Frame - Red, 58
- f1: Helmets
  f2 Sport-100 Helmet, Red
- f1: Helmets
  f2 Sport-100 Helmet, Black
- f1: Socks
  f2 Mountain Bike Socks, M
- f1: Socks
  f2 Mountain Bike Socks, L
- f1: Helmets
  f2 Sport-100 Helmet, Blue
- f1: Caps
  f2 AWC Logo Cap
. . .
  
*/
