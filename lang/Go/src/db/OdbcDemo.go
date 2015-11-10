////////////////////////////////////////////////////////////////////////////
// Porgram: OdbcDemo
// Purpose: Go MSSQL odbc demo, using the code.google.com/p/odbc driver
// Authors: Tong Sun (c) 2013-2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/alexbrainman/odbc"
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
		"driver=sql server;server=(local);database=tempdb;trusted_connection=yes;")
	if err != nil {
		fmt.Println("Connecting Error")
		return
	}
	defer conn.Close()
	fmt.Println("Sql Server Connected")

	test1(conn)
	test2(conn)
	testCreateInsertDelete(conn)
}

func test1(conn *sql.DB) {
	stmt, err := conn.Prepare("select top 5 database_id, name from sys.databases WHERE database_id >= 5")
	if err != nil {
		fmt.Println("Query Preparation Error", err)
		return
	}
	defer stmt.Close()
	fmt.Println("Query Prepared")

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
}

func test2(conn *sql.DB) {
	// Preparing Queries
	/*
		You should, in general, always prepare queries to be used multiple times.
		The result of preparing the query is a prepared statement, which can
		have ? placeholders for parameters that you'll provide when you execute
		the statement. This is much better than concatenating strings.
	*/

	stmt, err := conn.Prepare("select top 5 database_id, name from sys.databases WHERE database_id >= ?")
	if err != nil {
		log.Fatal(err)
	}

	row, err := stmt.Query(1)
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

func testCreateInsertDelete(db *sql.DB) {
	type friend struct {
		age       int
		isGirl    bool
		weight    float64
		dob       time.Time
		data      []byte
		canBeNull sql.NullString
	}
	var friends = map[string]friend{
		"glenda": {
			age:       5,
			isGirl:    true,
			weight:    15.5,
			dob:       time.Date(2000, 5, 10, 11, 1, 1, 0, time.Local),
			data:      []byte{0x0, 0x0, 0xb, 0xad, 0xc0, 0xde},
			canBeNull: sql.NullString{"aa", true},
		},
		"gopher": {
			age:       3,
			isGirl:    false,
			weight:    26.12,
			dob:       time.Date(2009, 5, 10, 11, 1, 1, 123e6, time.Local),
			data:      []byte{0x0},
			canBeNull: sql.NullString{"bbb", true},
		},
	}

	fmt.Printf("\nTesting create, insert & delete:\n")

	// create table
	db.Exec("drop table dbo.temp")
	db.Exec("create table dbo.temp (name varchar(20), age int, isGirl bit, weight decimal(5,2), dob datetime, data varbinary(10) null, canBeNull varchar(10) null)")

	func() {
		s, err := db.Prepare("insert into dbo.temp (name, age, isGirl, weight, dob, data, canBeNull) values (?, ?, ?, ?, ?, cast(? as varbinary(10)), ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer s.Close()
		for name, f := range friends {
			_, err := s.Exec(name, f.age, f.isGirl, f.weight, f.dob, f.data, f.canBeNull)
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = s.Exec("chris", 25, 0, 50, time.Date(2015, 12, 25, 0, 0, 0, 0, time.Local), "ccc", nil)
		if err != nil {
			log.Fatal(err)
		}
		_, err = s.Exec("null", 0, 0, 0, time.Date(2015, 12, 25, 1, 2, 3, 0, time.Local), nil, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// read from the table and verify returned results
	rows, err := db.Query("select name, age, isGirl, weight, dob, data, canBeNull from dbo.temp")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var is friend
		err = rows.Scan(&name, &is.age, &is.isGirl, &is.weight, &is.dob, &is.data, &is.canBeNull)
		if err != nil {
			log.Fatal(err)
		}
		want, ok := friends[name]
		if !ok {
			switch name {
			case "chris":
				// we know about chris, we just do not like him
			case "null":
				if is.canBeNull.Valid {
					fmt.Printf("null's canBeNull is suppose to be NULL, but is %v", is.canBeNull)
				}
			default:
				fmt.Printf("found %s, who is not my friend", name)
			}
			continue
		}
		if is.age < want.age {
			fmt.Printf("I did not know, that %s is so young (%d, but %d expected)", name, is.age, want.age)
			continue
		}
		if is.age > want.age {
			fmt.Printf("I did not know, that %s is so old (%d, but %d expected)", name, is.age, want.age)
			continue
		}
		if is.isGirl != want.isGirl {
			if is.isGirl {
				fmt.Printf("I did not know, that %s is a girl", name)
			} else {
				fmt.Printf("I did not know, that %s is a boy", name)
			}
			continue
		}
		if is.weight != want.weight {
			fmt.Printf("I did not know, that %s weights %dkg (%dkg expected)", name, is.weight, want.weight)
			continue
		}
		if !is.dob.Equal(want.dob) {
			fmt.Printf("I did not know, that %s's date of birth is %v (%v expected)", name, is.dob, want.dob)
			continue
		}
		if !bytes.Equal(is.data, want.data) {
			fmt.Printf("I did not know, that %s's data is %v (%v expected)", name, is.data, want.data)
			continue
		}
		if is.canBeNull != want.canBeNull {
			fmt.Printf("canBeNull for %s is wrong (%v, but %v expected)", name, is.canBeNull, want.canBeNull)
			continue
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// clean after ourselves
	db.Exec("drop table dbo.temp")
	fmt.Printf(" No news is good news\n")
}
