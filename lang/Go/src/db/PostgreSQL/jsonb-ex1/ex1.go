package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
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

	fmt.Println("\n-------- Querying record(s)")
	query1(db)
}

// https://www.mohitkhare.com/blog/json-in-postgres-with-golang/

type Order struct {
	CartID int  `json:"cart_id"`
	UserID int  `json:"user_id"`
	Cart   Cart `json:"cart"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

type CartItem struct {
	ItemID   int    `json:"item_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity,omitempty"`
	Price    int    `json:"price,omitempty"`
}

func (c Cart) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Cart) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c)
}

func query1(db *sql.DB) {
	// https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/
	sqlStatement := `SELECT cart_id, cart FROM orders WHERE cart_id=$1;`
	var cartId int
	var c Cart
	row := db.QueryRow(sqlStatement, 1)
	err := row.Scan(&cartId, &c)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(cartId, c)
	default:
		panic(err)
	}
}
