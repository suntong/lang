package main

import (
	"log"

	"github.com/kamva/mgm/v3"
)

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

func main() {
	crud()
}
