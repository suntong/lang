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

// The Attrs struct represents the data in the JSON/JSONB column. We can use
// struct tags to control how each field is encoded.
type Attrs struct {
	Name        string   `json:"name,omitempty"`
	Ingredients []string `json:"ingredients,omitempty"`
	Organic     bool     `json:"organic,omitempty"`
	Dimensions  struct {
		Weight float64 `json:"weight,omitempty"`
	} `json:"dimensions,omitempty"`
}

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

	// Initialize a new Attrs struct and add some values.
	attrs := new(Attrs)
	attrs.Name = "Pesto"
	attrs.Ingredients = []string{"Basil", "Garlic", "Parmesan", "Pine nuts", "Olive oil"}
	attrs.Organic = false
	attrs.Dimensions.Weight = 100.00

	// The database driver will call the Value() method and and marshall the
	// attrs struct to JSON before the INSERT.
	_, err = db.Exec("INSERT INTO items (attrs) VALUES($1)", attrs)
	if err != nil {
		log.Fatal(err)
	}

	// Similarly, we can also fetch data from the database, and the driver
	// will call the Scan() method to unmarshal the data to an Attr struct.
	item := new(Item)
	err = db.QueryRow("SELECT id, attrs FROM items ORDER BY id DESC LIMIT 1").Scan(&item.ID, &item.Attrs)
	if err != nil {
		log.Fatal(err)
	}

	// You can then use the struct fields as normal...
	weightKg := item.Attrs.Dimensions.Weight / 1000
	log.Printf("Item: %d, Name: %s, Weight: %.2fkg", item.ID, item.Attrs.Name, weightKg)
}
