package main

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}

type bookT struct {
	Name  string `json:"name" bson:"name"`
	Pages int    `json:"pages" bson:"pages"`
}

type book struct {
	MgmB // base type for mgm
	bookT
}

func newBook(name string, pages int) *book {
	return &book{bookT: bookT{
		Name:  name,
		Pages: pages,
	}, MgmB: MgmB{ID: NewID()}}
}

func crud() error {

	book := newBook("Test", 124)
	booksColl := mgm.Coll(book)

	if err := booksColl.Create(book); err != nil {
		return err
	}

	book.Name = "Moulin Rouge!"
	if err := booksColl.Update(book); err != nil {
		return err
	}

	return booksColl.Delete(book)
}

func main() {
	crud()
}
