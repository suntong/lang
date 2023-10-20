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
