package main

import (
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/teris-io/shortid"
)

// MgmB is the base type for mgm, Mongo Go Models
type MgmB struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	mgm.DateFields `bson:",inline"`
}

// == Model interface contains base methods that must be implemented by
// each model. If you're using the `DefaultModel` struct in your model,
// you don't need to implement any of these methods.

// PrepareID method prepares the ID value to be used for filtering
// e.g convert hex-string ID value to bson.ObjectId
func (f *MgmB) PrepareID(id interface{}) (interface{}, error) {
	return id, nil
}

// GetID method returns a model's ID
func (f *MgmB) GetID() interface{} {
	return f.ID
}

// SetID sets the value of a model's ID field.
func (f *MgmB) SetID(id interface{}) {
	f.ID = id.(string)
}

// == Model interface extra to mgm, to use shortid

var shortIDGenerator *shortid.Shortid

func init() {
	// init a single (1) worker *Shortid
	shortIDGenerator = shortid.MustNew(1, shortid.DefaultABC, uint64(time.Now().UnixNano()))
}

// NewID creates a new shortID.
func NewID() string {
	return shortIDGenerator.MustGenerate()
}

// MakeID makes/creates the model's ID.
func (f *MgmB) MakeID() {
	f.ID = NewID()
}
