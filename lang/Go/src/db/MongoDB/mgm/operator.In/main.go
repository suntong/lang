package main

import (
	"encoding/json"
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

type PersonT struct {
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Version     int    `json:"version,omitempty"`
	Scores      []int  `json:"scores,omitempty"`
	DateCreated struct {
		Date string `json:"$date,omitempty"`
	} `json:"dateCreated,omitempty"`
}

type person struct {
	MgmB    `bson:",inline"` // base type for mgm
	PersonT `bson:",inline"`
}

func crud() error {

	dataArray := []string{
		`{
      "name": "Andrea Le",
      "email": "andrea_le@fake-mail.com",
      "version": 5,
      "scores": [ 85, 95, 75 ],
      "dateCreated": { "$date": "2003-03-26" }
   }`,
		`{
      "email": "no_name@fake-mail.com",
      "version": 4,
      "scores": [ 90, 90, 70 ],
      "dateCreated": { "$date": "2001-04-15" }
   }`,
		`{
      "name": "Greg Powell",
      "email": "greg_powell@fake-mail.com",
      "version": 1,
      "scores": [ 65, 75, 80 ],
      "dateCreated": { "$date": "1999-02-10" }
   }`,
	}

	var persons []person
	for _, jsonStr := range dataArray {
		var person person
		err := json.Unmarshal([]byte(jsonStr), &person)
		if err != nil {
			log.Fatal(err)
		}
		person.NewID()
		persons = append(persons, person)
	}
	log.Printf("%#v\n", persons)
	coll := mgm.Coll(&person{})
	for _, person := range persons {
		err := coll.Create(&person)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil // personsColl.Delete(person)
}

func main() {
	crud()
}
