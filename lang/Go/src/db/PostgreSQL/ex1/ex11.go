package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func main() {
	// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// URL for PostgreSQL connection.
	psqlInfo = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)
	connStr := psqlInfo
	connStr = os.Getenv("PGDB_CONN")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("-------- Connecting")
	fmt.Println("Successfully connected!")

	fmt.Println("\n-------- Querying single record")
	query1(db)
	fmt.Println("\n-------- Querying multiple records")
	query2(db)
}

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func query1(db *sql.DB) {
	// https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/
	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	var user User
	row := db.QueryRow(sqlStatement, 3)
	err := row.Scan(&user.ID, &user.Age, &user.FirstName,
		&user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
}

func query2(db *sql.DB) {
	// https://www.calhoun.io/querying-for-multiple-records-with-gos-sql-package/
	rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, firstName)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
