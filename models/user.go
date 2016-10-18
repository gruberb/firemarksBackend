package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name"`
	EMail     string        `json:"email"`
	Password  string        `form:"password"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type PublicUser struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name"`
	EMail string        `json:"email"`
	Token string        `json:"token"`
}

// NewUser creates a new User with ID and CreatedAt
func NewUser() *User {
	return &User{
		ID:        bson.NewObjectId(),
		CreatedAt: bson.Now(),
	}
}

// Users collection
func Users() *mgo.Collection {
	return db.C("users")
}

func (u *User) cryptPassword(p string) error {
	// TODO: Add secret salt to password to make it more robust
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash[:])
	return nil
}

// Validate run validations for the model
func (u *User) Validate() (bool, []ValidationError) {
	errors := []ValidationError{}

	// Validate: Name
	if len(u.Name) == 0 {
		errors = append(errors, ValidationError{
			"name", "Name is missing",
		})
	}

	// TODO Add valid regex or better way of validating emails
	// TODO Make email unique so that no other user can use the same email (on creation only)
	if len(u.EMail) == 0 {
		errors = append(errors, ValidationError{
			"email", "E-Mail is missing",
		})
	}

	if len(u.Password) == 0 {
		errors = append(errors, ValidationError{
			"password", "Password is missing",
		})
	}

	return (len(errors) == 0), errors
}

// QueryUsers ...
func QueryUsers(results *[]User) error {
	return Users().Find(nil).All(results)
}

// FindLink ...
func FindUser(email string, user *User) error {
	query := bson.M{"email": email}
	return Users().Find(query).One(user)
}

// CreateUser ...
func CreateUser(newUser *User) error {
	if err := newUser.cryptPassword(newUser.Password); err != nil {
		return err
	}

	return Users().Insert(newUser)
}
