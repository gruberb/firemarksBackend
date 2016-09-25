package models

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


// Link model
type Link struct {
	ID					bson.ObjectId `json:"id" bson:"_id,omitempty"`
	URL 				string `json:"url"`
	Name 				string `json:"name"`
	CreatedAt 	time.Time `json:"created_at,omitempty" bson:",omitempty"`
}


// Validate run validations for the model
func (m *Link) Validate() (bool, []ValidationError) {
	errors := []ValidationError{}

	// Validate: Name
	if len(m.Name) == 0 {
		errors = append(errors, ValidationError{
			"Name is missing",
		})
	}

	return (len(errors) == 0), errors
}


// NewLink creates a new Link with ID and CreatedAt
func NewLink() *Link {
	return &Link{
		ID: bson.NewObjectId(),
		CreatedAt: bson.Now(),
	}
}


// LinkCollection returns the collection for links
func LinkCollection() *mgo.Collection {
	return db.C("links")
}
