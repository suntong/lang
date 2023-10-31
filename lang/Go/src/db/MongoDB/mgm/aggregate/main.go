package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"github.com/kamva/mgm/v3/field"
	"go.mongodb.org/mongo-driver/bson"
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
	book
	author []author
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

func seed() {
	author := newAuthor("Mehran")
	_ = mgm.Coll(author).Create(author)

	book := newBook("Test", 124, author.ID)
	_ = mgm.Coll(book).Create(book)

}

func delSeededData() {
	_, _ = mgm.Coll(&book{}).DeleteMany(nil, bson.M{})
	_, _ = mgm.Coll(&author{}).DeleteMany(nil, bson.M{})
}

func lookup1() {
	authorCollName := mgm.Coll(&author{}).Name()
	//result := []book{} // X: will miss the author field!
	// var result bson.M
	// X: results argument must be a pointer to a slice, but was a pointer to map
	// result := []book2{} // X: all empty
	result := []bson.M{}
	err := mgm.Coll(&book{}).SimpleAggregate(&result,
		builder.S(builder.Lookup(authorCollName, "author_id", "_id", "author")))
	checkError(err)
	fmt.Printf("1] %v\n   %+[1]v\n", result)
}

func lookup() error {

	// Author model's collection
	authorColl := mgm.Coll(&author{})

	pipeline := bson.A{
		builder.S(builder.Lookup(authorColl.Name(), "author_id", field.ID, "author")),
	}

	cur, err := mgm.Coll(&book{}).Aggregate(mgm.Ctx(), pipeline)
	checkError(err)

	defer cur.Close(nil)

	for cur.Next(nil) {
		var result bson.M
		err := cur.Decode(&result)
		checkError(err)

		// do something with result....
		fmt.Printf("2] %+v\n", result)
	}

	return nil
}

func main() {
	seed()
	defer delSeededData()

	lookup1()
	lookup()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
