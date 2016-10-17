package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// URL ...
type URL struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Address   string        `json:"address"`
	Hash      string        `json:"hash"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:",omitempty"`
}

// NewURL ...
func NewURL() *URL {
	return &URL{
		ID:        bson.NewObjectId(),
		CreatedAt: bson.Now(),
	}
}

// CreateURLHash ...
func CreateURLHash(a string) string {
	hash := md5.Sum([]byte(a))
	return hex.EncodeToString(hash[:])
}
