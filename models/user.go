package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name"`
	EMail     string        `json:"email"`
	Password  string        `json:"-"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:",omitempty"`
}

// NewUser creates a new User with ID and CreatedAt
func NewUser() *User {
	return &User{
		ID:        bson.NewObjectId(),
		CreatedAt: bson.Now(),
	}
}

func (u *User) cryptPassword(p []byte) error {
	var error error
	var hash []byte

	if hash, error = bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost); error != nil {
		return error
	}

	u.Password = string(hash[:])
	return nil
}

func (u *User) Validate() (bool, []ValidationError) {
	errors := []ValidationError{}

	// Validate: Name
	if len(u.Name) == 0 {
		errors = append(errors, ValidationError{
			"name", "Name is missing",
		})
	}

	// TODO Add valid regex or better way of validating emails
	if len(u.EMail) == 0 {
		errors = append(errors, ValidationError{
			"name", "E-Mail is missing",
		})
	}

	return (len(errors) == 0), errors
}

// Users collection
func Users() *mgo.Collection {
	return db.C("users")
}

// CreateUser
func CreateUser(newUser *User) error {
	newUser.cryptPassword([]byte(newUser.Password))
	return Users().Insert(newUser)
}
