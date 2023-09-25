// https://www.loginradius.com/blog/engineering/mongodb-as-datasource-in-golang/
// https://github.com/LoginRadius/engineering-blog-samples/tree/master/GoLang/MongoDriverForGolang

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/teris-io/shortid"
)

type anys []interface{}

func (a anys) ToStringSlice() []string {
	aString := make([]string, len(a))
	for i, v := range a {
		aString[i] = v.(string)
	}
	return aString
}

// Book - We will be using this Book type to perform crud operations
type Book struct {
	ID        string `bson:"_id"`
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies    int
}

var shortIDGenerator *shortid.Shortid

func main() {
	shortIDGenerator = shortid.MustNew(1, shortid.DefaultABC, uint64(time.Now().UnixNano()))

	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MGDB_CONN"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	booksCollection := client.Database("testdb").Collection("books")

	// Insert One document
	id := shortIDGenerator.MustGenerate()
	book1 := Book{id, "Animal Farm", "George Orwell", "0451526341", "Signet Classics", 100}
	insertResult, err := booksCollection.InsertOne(context.TODO(), book1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted single document: %+v\n", insertResult)
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	id1 := insertResult.InsertedID

	// Insert multiple documents
	id = shortIDGenerator.MustGenerate()
	book2 := Book{id, "Super Freakonomics", "Steven D. Levitt", "0062312871", "HARPER COLLINS USA", 100}
	id = shortIDGenerator.MustGenerate()
	book3 := Book{id, "The Alchemist", "Paulo Coelho", "0062315005", "HarperOne", 100}
	multipleBooks := []interface{}{book2, book3}

	insertManyResult, err := booksCollection.InsertMany(context.TODO(), multipleBooks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	ids := append(insertManyResult.InsertedIDs, id1)
	fmt.Printf("IDs: %#v\n", ids)
	// IDs: []interface {}{"EPpeg1l-Z", "Eiper15-ZZ", "kZSeg1l-Z"}

	//Update one document
	filter := bson.D{{"isbn", "0451526341"}}

	update := bson.D{
		{"$inc", bson.D{
			{"copies", 1},
		}},
	}

	updateResult, err := booksCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Find single document
	// A variable in which result will be decoded
	var result Book

	err = booksCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	// Find multiple documents by id
	// https://stackoverflow.com/questions/32264225/
	//objectIDs := []string{"ObjectId1", "ObjectId2"}
	// https://goplay.tools/snippet/w_A7Ca7MNO0
	objectIDs := anys(ids.([]interface{})).ToStringSlice()
	fmt.Printf("IDs: %#v\n", objectIDs)
	filter = bson.M{"_id": bson.M{"$in": objectIDs}}
	{
		cursor, err := booksCollection.Find(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		var books []Book
		if err = cursor.All(context.TODO(), &books); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Found multiple documents by ids of", ids)
		fmt.Println(books)
	}

	// Find multiple documents
	cursor, err := booksCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var books []Book
	if err = cursor.All(context.TODO(), &books); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found all documents")
	fmt.Println(books)

	//Delete Documents
	deleteCollection, err := booksCollection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the books collection\n", deleteCollection.DeletedCount)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

/*

$ go run main.go
Connected to MongoDB!
Inserted single document: &{InsertedID:k0359TIcu}
Inserted a single document:  k0359TIcu
Inserted multiple documents:  [ag_x9cqTl kgp5fcIcuu]
Matched 1 documents and updated 1 documents.
Found a single document: {ID:k0359TIcu Title:Animal Farm Author:George Orwell ISBN:0451526341 Publisher:Signet Classics Copies:101}
[{k0359TIcu Animal Farm George Orwell 0451526341 Signet Classics 101} {ag_x9cqTl Super Freakonomics Steven D. Levitt 0062312871 HARPER COLLINS USA 100} {kgp5fcIcuu The Alchemist Paulo Coelho 0062315005 HarperOne 100}]
Deleted 3 documents in the books collection
Connection to MongoDB closed.

*/
