package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/field"
	"github.com/kamva/mgm/v3/operator"
	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	Pages            int                `json:"pages" bson:"pages"`
	AuthorID         primitive.ObjectID `json:"author_id" bson:"author_id"`
}

func newBook(name string, pages int, authID primitive.ObjectID) *book {
	return &book{
		Name:     name,
		Pages:    pages,
		AuthorID: authID,
	}
}

type author struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

func newAuthor(name string) *author {
	return &author{
		Name: name,
	}
}

type book2 struct {
	book   `bson:",inline"`
	Author []author `bson:"author"` // will get `author:[]` if just use `author`
}

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

func lookup() error {

	// Author model's collection
	authorColl := mgm.Coll(&author{})

	// db.books.aggregate([ { $match: { "name": "Test2", "pages": 126 } }, { $lookup: { from: "authors", localField: "author_id", foreignField: "_id", as: "author" } }] )
	// https://github.com/Kamva/mgm#aggregation
	pipeline := A{
		builder.S(builder.New(operator.Match, M{"name": "Test2", "pages": 126})),
		builder.S(builder.Lookup(authorColl.Name(), "author_id", field.ID, "author")),
	}

	// X: err := mgm.Coll(&book{}).SimpleAggregate(&result, pipeline)
	// Fatal error: cannot marshal type primitive.A to a BSON Document

	cur, err := mgm.Coll(&book{}).Aggregate(mgm.Ctx(), pipeline)
	checkError(err)

	defer cur.Close(nil)

	for cur.Next(nil) {
		var result = book2{} // M
		err := cur.Decode(&result)
		checkError(err)

		// do something with result....
		fmt.Printf("2] %+v\n", result)
	}

	return nil
}

func main() {
	lookup()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
