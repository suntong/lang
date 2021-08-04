
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

var db *sql.DB
var server = "sqlserver210802.database.windows.net"
var port = 1433
var user = "sqluser"
var password = "<your_password>"
var database = "db1"

func main() {
	password = "Pa$$w0rd..."
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
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
	rows, errdb := db.Query(`
SELECT TOP 20 pc.Name as CategoryName, p.name as ProductName
  FROM SalesLT.ProductCategory pc
  JOIN SalesLT.Product p ON pc.productcategoryid = p.productcategoryid;
`	)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&f1, &f2)
		if err != nil {
			fmt.Println("  Error Query db: select - ", err.Error())
		} else {
			fmt.Printf("- f1: %s\n  f2 %s\n", f1, f2)
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
