// https://www.alexedwards.net/blog/using-postgresql-jsonb

package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

type Item struct {
	ID    int
	Attrs Attrs
}

// The Attrs struct represents the data in the JSON/JSONB column,
// when we don't know in advance what keys and values from JSONB data.
// thus need to map the contents of the JSONB column to and from a map[string]interface{} instead. The big downside of this is that you will need to type assert any values that you retrieve from the database in order to use them.
type Attrs map[string]interface{}

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Attrs) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Attrs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func main() {
	// db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
	// db, err := sql.Open("postgres", "postgres://postgres:password@localhost/postgres")
	// X: pq: SSL is not enabled on the server
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n-------- Querying record(s) with unknown fields")
	query(db)
}

func query(db *sql.DB) {
	item := new(Item)
	item.Attrs = Attrs{
		"name":        "Passata",
		"ingredients": []string{"Tomatoes", "Onion", "Olive oil", "Garlic"},
		"organic":     true,
		"dimensions": map[string]interface{}{
			"weight": 250.00,
		},
	}

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err := db.Exec("INSERT INTO items (attrs) VALUES($1)", item.Attrs)
	if err != nil {
		log.Fatal(err)
	}

	// Similarly, we can also fetch data from the database, and the driver
	// will call the Scan() method to unmarshal the data to an Attr struct.
	item = new(Item)
	err = db.QueryRow("SELECT id, attrs FROM items ORDER BY id DESC LIMIT 1").Scan(&item.ID, &item.Attrs)
	if err != nil {
		log.Fatal(err)
	}

	// As you cannot use the struct fields as normal...
	name, ok := item.Attrs["name"].(string)
	if !ok {
		log.Fatal("unexpected type for name")
	}
	dimensions, ok := item.Attrs["dimensions"].(map[string]interface{})
	if !ok {
		log.Fatal("unexpected type for dimensions")
	}
	weight, ok := dimensions["weight"].(float64)
	if !ok {
		log.Fatal("unexpected type for weight")
	}
	weightKg := weight / 1000
	log.Printf("%s: %.2fkg", name, weightKg)

}
