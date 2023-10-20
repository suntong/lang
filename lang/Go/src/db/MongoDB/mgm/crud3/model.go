package main

import (
	"github.com/kamva/mgm/v3"
)

type book struct {
	ID             string `json:"_id" bson:"_id,omitempty"`
	Name           string `json:"name" bson:"name"`
	Pages          int    `json:"pages" bson:"pages"`
	mgm.DateFields `bson:",inline"`
}

func newBook(name string, pages int) *book {
	return &book{
		ID: "123456789",
		Name:  name,
		Pages: pages,
	}
}


// GetID method returns a model's ID
func (b *book) GetID() interface{} {
	return b.ID
}
