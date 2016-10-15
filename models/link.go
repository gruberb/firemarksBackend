package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Link model
type Link struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	URL       string        `json:"url"`
	Name      string        `json:"name"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:",omitempty"`
}

// Validate run validations for the model
func (m *Link) Validate() (bool, []ValidationError) {
	errors := []ValidationError{}

	// Validate: Name
	if len(m.Name) == 0 {
		errors = append(errors, ValidationError{
			"name", "Name is missing",
		})
	}

	// Validate: URL
	if len(m.URL) == 0 {
		errors = append(errors, ValidationError{
			"url", "URL is missing",
		})
	}

	return (len(errors) == 0), errors
}

// NewLink creates a new Link with ID and CreatedAt
func NewLink() *Link {
	return &Link{
		ID:        bson.NewObjectId(),
		CreatedAt: bson.Now(),
	}
}

// Links ...
func Links() *mgo.Collection {
	return db.C("links")
}

// QueryLinks ...
func QueryLinks(results *[]Link) error {
	return Links().Find(nil).All(results)
}

// FindLink ...
func FindLink(id string, link *Link) error {
	objectID, err := stringToObjectID(id)
	if err != nil {
		return err
	}

	query := bson.M{"_id": objectID}
	return Links().Find(query).One(link)
}

// CreateLink ...
func CreateLink(newLink *Link) error {
	return Links().Insert(newLink)
}

// UpdateLink ...
func UpdateLink(id string, changes *Link) error {
	objectID, err := stringToObjectID(id)
	if err != nil {
		return err
	}

	changeSet := bson.M{"$set": changes}
	return Links().UpdateId(objectID, changeSet)
}

// DeleteLink ...
func DeleteLink(id string) error {
	objectID, err := stringToObjectID(id)
	if err != nil {
		return err
	}

	return Links().RemoveId(objectID)
}
