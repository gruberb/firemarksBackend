package models


import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)


type (
	// ValidationError represents a failed validation
	ValidationError struct {
		Field string		`json:"field"`
		Message string	`json:"message"`
	}
)


func stringToObjectID(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(id) {
		return "", errors.New("The id is not valid")
	}

	return bson.ObjectIdHex(id), nil
}
