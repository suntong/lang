////////////////////////////////////////////////////////////////////////////
// Porgram: go-mssqldb demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://bsilverstrim.blogspot.ca/2015/08/golang-and-mssql-databases-example.html
////////////////////////////////////////////////////////////////////////////

package main

// Notice in the import list there's one package prefaced by a ".",
// which allows referencing functions in that package without naming the library in
// the call (if using . "fmt", I can call Println as Println, not fmt.Println)
import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

const strVERSION string = "0.18 compiled on 8/11/2015"

// sqltest is a small application for demonstrating/testing/learning about SQL database connectivity from Go
func main() {

	// Flags
	ptrVersion := flag.Bool("version", false, "Display program version")
	ptrDeleteIt := flag.Bool("deletedb", false, "Delete the database")
	// ptrServer := flag.String("server", "localhost", "Server to connect to")
	// ptrUser := flag.String("username", "testuser", "Username for authenticating to database; if you use a backslash, it must be escaped or in quotes")
	// ptrPass := flag.String("password", "", "Password for database connection")
	ptrDBName := flag.String("dbname", "test_db", "Database name")

	flag.Parse()

	// Does the user just want the version of the application?
	if *ptrVersion == true {
		fmt.Println("Version " + strVERSION)
		os.Exit(0)
	}

	// Open connection to the database server; this doesn't verify anything until you
	// perform an operation (such as a ping).
	//db, err := sql.Open("mssql", "server="+*ptrServer+";user id="+*ptrUser+";password="+*ptrPass)

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

	db, err := sql.Open("mssql", c)
	if err != nil {
		fmt.Println("From Open() attempt: " + err.Error())
	}

	// When main() is done, this should close the connections
	defer db.Close()

	// Does the user want to delete the database?
	if *ptrDeleteIt == true {
		boolDBExist, err := CheckDB(db, *ptrDBName)
		if err != nil {
			fmt.Println("Error running CheckDB: " + err.Error())
			os.Exit(1)
		}
		if boolDBExist {
			fmt.Println("(sqltest) Deleting database " + *ptrDBName + "...")
			DropDB(db, *ptrDBName)
			os.Exit(0)
		} else {

			// Database doesn't seem to exist...
			fmt.Println("(sqltest) Database " + *ptrDBName + " doesn't appear to exist...")
			os.Exit(1)

		}
	}

	// Let's start the tests...
	fmt.Println("********************************")

	// Is the database running?
	strResult := PingServer(db)
	fmt.Println("(sqltest) Ping of Server Result Was: " + strResult)

	fmt.Println("********************************")

	// Does the database exist?
	boolDBExist, err := CheckDB(db, *ptrDBName)
	if err != nil {
		fmt.Println("(sqltest) Error running CheckDB: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("(sqltest) Database Existence Check: " + strconv.FormatBool(boolDBExist))

	fmt.Println("********************************")

	// If it doesn't exist, let's create the base database
	if !boolDBExist {

		CreateDBAndTable(db, *ptrDBName)
		fmt.Println("********************************")

	}

	// Enter a test record
	boolDBExist, err = CheckDB(db, *ptrDBName)
	if err != nil {
		fmt.Println("(sqltest) CheckDB() error: " + err.Error())
		os.Exit(1)
	}

	if boolDBExist == true {

		err := AddToContent(db, *ptrDBName, "Bob", 1437506592, "Hello!")
		if err != nil {
			fmt.Println("(sqltest) Error adding line to content: " + err.Error())
			os.Exit(1)
		}

		err = AddToContent(db, *ptrDBName, "user", 1437506648, "Now testing memory")
		if err != nil {
			fmt.Println("(sqltest) Error adding line to content: " + err.Error())
			os.Exit(1)
		}

		err = AddToContent(db, *ptrDBName, "user", 1437503394, "test, text!")
		if err != nil {
			fmt.Println("(sqltest) Error adding line to content: " + err.Error())
			os.Exit(1)
		}

		err = AddToContent(db, *ptrDBName, "Bob", 1437506592, "Hope this works!")
		if err != nil {
			fmt.Println("(sqltest) Error adding line to content: " + err.Error())
			os.Exit(1)
		}

	}

	fmt.Println("(sqltest) Completed entering test records.")

	fmt.Println("********************************")

	fmt.Println("(sqltest) Deleting records from a particular source.")

	// Delete from a source
	int64Deleted, err := RemoveFromContentBySource(db, "user")
	if err != nil {
		fmt.Println("(sqltest) Error deleting records by source: " + err.Error())
		os.Exit(1)
	} else {

		// How many records were removed?
		fmt.Println("Removed " + strconv.FormatInt(int64Deleted, 10) + " records")
		fmt.Println("********************************")

	}

	// Get the content
	slcstrSource, slcint64Timestamp, slcstrContent, err := GetContent(db)
	if err != nil {
		fmt.Println("(sqltest) Error getting content: " + err.Error())
	}

	// Now read the contents
	for i := range slcstrContent {

		fmt.Println("Entry " + strconv.Itoa(i) + ": " + strconv.FormatInt(slcint64Timestamp[i], 10) + ", from " + slcstrSource[i] + ": " + slcstrContent[i])

	}

}

// Package custom_DB_functions contains functions customized to manipulate MSSQL databases/tables
// for our application
//
// Version 0.15, 8-13-2015
//package custom_DB_functions

// import (
//  "database/sql"
//  _ "github.com/denisenkom/go-mssqldb"
//  "strconv"
// )

// PingServer uses a passed database handle to check if the database server works
func PingServer(db *sql.DB) string {

	err := db.Ping()
	if err != nil {
		return ("From Ping() Attempt: " + err.Error())
	}

	return ("Database Ping Worked...")

}

// CheckDB checks if the database "strDBName" exists on the MSSQL database engine.
func CheckDB(db *sql.DB, strDBName string) (bool, error) {

	// Does the database exist?
	result, err := db.Query("SELECT db_id('" + strDBName + "')")
	defer result.Close()
	if err != nil {
		return false, err
	}

	for result.Next() {
		var s sql.NullString
		err := result.Scan(&s)
		if err != nil {
			return false, err
		}

		// Check result
		if s.Valid {
			return true, nil
		} else {
			return false, nil
		}
	}

	// This return() should never be hit...
	return false, err
}

// CreateDBAndTable creates a new content database on the SQL Server along with
// the necessary tables. Keep in mind the user credentials that opened the database
// connection with sql.Open must have at least dbcreator rights to the database. The
// table (testtable) will have columns source (nvarchar), timestamp (bigint), and
// content (nvarchar).
func CreateDBAndTable(db *sql.DB, strDBName string) error {

	// Create the database
	_, err := db.Exec("CREATE DATABASE [" + strDBName + "]")
	if err != nil {
		return (err)
	}

	// Let's turn off AutoClose
	_, err = db.Exec("ALTER DATABASE [" + strDBName + "] SET AUTO_CLOSE OFF;")
	if err != nil {
		return (err)
	}

	// Create the tables
	_, err = db.Exec("USE " + strDBName + "; CREATE TABLE testtable (source nvarchar(100) NOT NULL, timestamp bigint NOT NULL, content nvarchar(4000) NOT NULL)")
	if err != nil {
		return (err)
	}

	return nil

}

// DropDB deletes the database strDBName.
func DropDB(db *sql.DB, strDBName string) error {

	// Drop the database
	_, err := db.Exec("DROP DATABASE [" + strDBName + "]")

	if err != nil {
		return err
	}

	return nil

}

// AddToContent adds new content to the database.
func AddToContent(db *sql.DB, strDBName string, strSource string, int64Timestamp int64, strContent string) error {

	// Add a record entry
	_, err := db.Exec("USE " + strDBName + "; INSERT INTO testtable (source, timestamp, content) VALUES ('" + strSource + "','" + strconv.FormatInt(int64Timestamp, 10) + "','" + strContent + "');")
	if err != nil {
		return err
	}

	return nil

}

// RemoveFromContentBySource removes a record from the database with source strSource. The
// int64 returned is a message indicating the number of rows affected.
func RemoveFromContentBySource(db *sql.DB, strSource string) (int64, error) {

	// Remove entries containing the source...
	result, err := db.Exec("DELETE FROM testtable WHERE source='" + strSource + "';")
	if err != nil {
		return 0, err
	}

	// What was the result?
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, nil

}

// Query the content in the database and return the source (string), timestamp (int64), and
// content (string) as slices
func GetContent(db *sql.DB) ([]string, []int64, []string, error) {

	var slcstrContent []string
	var slcint64Timestamp []int64
	var slcstrSource []string

	// Run the query
	rows, err := db.Query("SELECT source, timestamp, content FROM testtable")
	if err != nil {
		return slcstrSource, slcint64Timestamp, slcstrContent, err
	}
	defer rows.Close()

	for rows.Next() {

		// Holding variables for the content in the columns
		var source, content string
		var timestamp int64

		// Get the results of the query
		err := rows.Scan(&source, &timestamp, &content)
		if err != nil {
			return slcstrSource, slcint64Timestamp, slcstrContent, err
		}

		// Append them into the slices that will eventually be returned to the caller
		slcstrSource = append(slcstrSource, source)
		slcstrContent = append(slcstrContent, content)
		slcint64Timestamp = append(slcint64Timestamp, timestamp)
	}

	return slcstrSource, slcint64Timestamp, slcstrContent, nil

}

/*

$ HOST=localhost SQLUSER=sa SQLPASSWORD=sa DATABASE=test go run GoMssqlDb-Demo1.go
********************************
(sqltest) Ping of Server Result Was: Database Ping Worked...
********************************
(sqltest) Database Existence Check: false
********************************
********************************
(sqltest) Completed entering test records.
********************************
(sqltest) Deleting records from a particular source.
Removed 2 records
********************************
Entry 0: 1437506592, from Bob: Hello!
Entry 1: 1437506592, from Bob: Hope this works!

*/
