// https://www.loginradius.com/blog/engineering/mongodb-as-datasource-in-golang/
// https://github.com/LoginRadius/engineering-blog-samples/tree/master/GoLang/MongoDriverForGolang

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Book - We will be using this Book type to perform crud operations
type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies    int
}

func main() {
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
	book1 := Book{"Animal Farm", "George Orwell", "0451526341", "Signet Classics", 100}
	insertResult, err := booksCollection.InsertOne(context.TODO(), book1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert multiple documents
	book2 := Book{"Super Freakonomics", "Steven D. Levitt", "0062312871", "HARPER COLLINS USA", 100}
	book3 := Book{"The Alchemist", "Paulo Coelho", "0062315005", "HarperOne", 100}
	multipleBooks := []interface{}{book2, book3}

	insertManyResult, err := booksCollection.InsertMany(context.TODO(), multipleBooks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

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

	// Find multiple documents
	cursor, err := booksCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var books []Book
	if err = cursor.All(context.TODO(), &books); err != nil {
		log.Fatal(err)
	}
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
Inserted a single document:  ObjectID("64e626c83e1752e13beacad4")
Inserted multiple documents:  [ObjectID("64e626c83e1752e13beacad5") ObjectID("64e626c83e1752e13beacad6")]
Matched 1 documents and updated 1 documents.
Found a single document: {Title:Animal Farm Author:George Orwell ISBN:0451526341 Publisher:Signet Classics Copies:101}
[{Animal Farm George Orwell 0451526341 Signet Classics 101} {Super Freakonomics Steven D. Levitt 0062312871 HARPER COLLINS USA 100} {The Alchemist Paulo Coelho 0062315005 HarperOne 100}]
Deleted 3 documents in the books collection
Connection to MongoDB closed.

*/
