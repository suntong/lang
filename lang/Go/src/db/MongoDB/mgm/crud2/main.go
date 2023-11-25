package main

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
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

type BookT struct {
	Name  string `json:"name" bson:"name"`
	Pages int    `json:"pages" bson:"pages"`
}

type book struct {
	MgmB  `bson:",inline"` // base type for mgm
	BookT `bson:",inline"`
}

func newBook(name string, pages int) *book {
	return &book{BookT: BookT{
		Name:  name,
		Pages: pages,
	}, MgmB: MgmB{ID: NewID()}}
}

func crud() error {

	book := newBook("Test", 124)
	log.Printf("%#v\n", book)
	booksColl := mgm.Coll(book)

	if err := booksColl.Create(book); err != nil {
		return err
	}

	book.Name = "Moulin Rouge!"
	if err := booksColl.Update(book); err != nil {
		return err
	}
	log.Printf("%#v\n", book)

	return nil // booksColl.Delete(book)
}

func simpleFind() {
	log.Println("\n\n## SimpleFind")
	result := []book{}
	if err := mgm.Coll(&book{}).SimpleFind(&result, bson.M{}, &options.FindOptions{
		Sort: map[string]int{"updated_at": -1}}); err != nil {
		return
	}
	log.Printf("%#v\n", result)
}

func main() {
	//crud()
	simpleFind()
}
