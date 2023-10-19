package main

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	_ = mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}

type bookT struct {
        Name             string `json:"name" bson:"name"`
        Pages            int    `json:"pages" bson:"pages"`
}


type book struct {
	mgm.DefaultModel `bson:",inline"`
	bookT
}

func newBook(name string, pages int) *book {
	return &book{bookT: bookT{
		Name:  name,
		Pages: pages,
	}}
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
