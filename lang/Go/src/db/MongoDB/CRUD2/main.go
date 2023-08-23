// https://blog.devops.dev/build-a-simple-mongodb-crud-example-with-go-and-docker-c5202af1b449

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
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(os.Getenv("MGDB_CONN")),
	)

	defer func() {
		cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("mongodb disconnect error : %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("connection error :%v", err)
		return
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		return
	}
	log.Println("ping success")

	// database and collection
	database := mongoClient.Database("demo")
	sampleCollection := database.Collection("sampleCollection")
	sampleCollection.Drop(ctx)
	// -- start from fresh collection every time. Therefore - drop

	// insert one
	insertedDocument := bson.M{
		"name":       "michael",
		"content":    "test content",
		"bank_money": 1000,
		"create_at":  time.Now(),
	}
	insertedResult, err := sampleCollection.InsertOne(ctx, insertedDocument)

	if err != nil {
		log.Fatalf("inserted error : %v", err)
		return
	}
	log.Println("======= inserted id ================")
	log.Printf("inserted ID is : %v", insertedResult.InsertedID)

	// query all data
	log.Println("== query all data ==")
	cursor, err := sampleCollection.Find(ctx, options.Find())
	if err != nil {
		log.Fatalf("find collection err : %v", err)
		return
	}
	var queryResult []bson.M
	if err := cursor.All(ctx, &queryResult); err != nil {
		log.Fatalf("query mongodb result")
		return
	}

	for _, doc := range queryResult {
		fmt.Println(doc)
	}

	// insert many data
	log.Println("=========== inserted many data ===============")
	insertedManyDocument := []interface{}{
		bson.M{
			"name":       "Andy",
			"content":    "new test content",
			"bank_money": 1500,
			"create_at":  time.Now().Add(36 * time.Hour),
		},
		bson.M{
			"name":       "Jack",
			"content":    "jack content",
			"bank_money": 800,
			"create_at":  time.Now().Add(12 * time.Hour),
		},
	}

	insertedManyResult, err := sampleCollection.InsertMany(ctx, insertedManyDocument)
	if err != nil {
		log.Fatalf("inserted many error : %v", err)
		return
	}

	for _, doc := range insertedManyResult.InsertedIDs {
		fmt.Println(doc)
	}

	log.Println("=========== query specific data =====================")
	// query specific data
	filter := bson.D{
		bson.E{
			Key: "bank_money",
			Value: bson.D{
				bson.E{
					Key:   "$gt",
					Value: 900,
				},
			},
		},
	}

	filterCursor, err := sampleCollection.Find(
		ctx,
		filter,
	)
	if err != nil {
		log.Fatalf("filter query data error : %v", err)
		return
	}
	var filterResult []bson.M
	err = filterCursor.All(ctx, &filterResult)
	if err != nil {
		log.Fatalf("filter result %v", err)
		return
	}

	for _, filterDoc := range filterResult {
		fmt.Println(filterDoc)
	}

	updateManyFilter := bson.D{
		bson.E{
			Key:   "name",
			Value: "michael",
		},
	}

	updateSet := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "bank_money",
					Value: 2000,
				},
			},
		},
	}
	// update
	updateManyResult, err := sampleCollection.UpdateMany(
		ctx,
		updateManyFilter,
		updateSet,
	)
	if err != nil {
		log.Fatalf("update error : %v", err)
		return
	}

	log.Println("========= updated modified count ===========")
	fmt.Println(updateManyResult.ModifiedCount)

	// check if updated with find solution
	checkedCursor, err := sampleCollection.Find(
		ctx,
		bson.D{
			bson.E{
				Key:   "name",
				Value: "michael",
			},
		},
	)
	if err != nil {
		log.Fatalf("check result error : %v", err)
		return
	}
	var checkedResult []bson.M
	err = checkedCursor.All(ctx, &checkedResult)
	if err != nil {
		log.Fatalf("get check information error : %v", err)
		return
	}
	log.Println("=========== checked updated result ==============")
	for _, checkedDoc := range checkedResult {
		fmt.Println(checkedDoc)
	}
	log.Println("===============================")
	// delete Many

	deleteManyResult, err := sampleCollection.DeleteMany(
		ctx,
		bson.D{
			bson.E{
				Key: "bank_money",
				Value: bson.D{
					bson.E{
						Key:   "$lt",
						Value: 1000,
					},
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("delete many data error : %v", err)
		return
	}
	log.Println("===== delete many data modified count =====")
	fmt.Println(deleteManyResult.DeletedCount)
}

/*

$ go run main.go
2023/08/23 12:15:27 ping success

2023/08/23 12:15:27 ======= inserted id ================
2023/08/23 12:15:27 inserted ID is : ObjectID("64e6309f45cd37a74b68667f")

2023/08/23 12:15:27 == query all data ==
map[_id:ObjectID("64e6309f45cd37a74b68667f") bank_money:1000 content:test content create_at:1692807327084 name:michael]

2023/08/23 12:15:27 =========== inserted many data ===============
ObjectID("64e6309f45cd37a74b686680")
ObjectID("64e6309f45cd37a74b686681")

2023/08/23 12:15:27 =========== query specific data =====================
map[_id:ObjectID("64e6309f45cd37a74b68667f") bank_money:1000 content:test content create_at:1692807327084 name:michael]
map[_id:ObjectID("64e6309f45cd37a74b686680") bank_money:1500 content:new test content create_at:1692936927203 name:Andy]

2023/08/23 12:15:27 ========= updated modified count ===========
1

2023/08/23 12:15:27 =========== checked updated result ==============
map[_id:ObjectID("64e6309f45cd37a74b68667f") bank_money:2000 content:test content create_at:1692807327084 name:michael]
2023/08/23 12:15:27 ===============================

2023/08/23 12:15:27 ===== delete many data modified count =====
1

*/
