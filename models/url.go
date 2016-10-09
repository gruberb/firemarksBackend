package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Link model
type Url struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Address   string
	Hash      string
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func NewUrl() *Url {
	return &Url{
		ID:        bson.NewObjectId(),
		CreatedAt: bson.Now(),
	}
}

func CreateUrlHash(a string) string {
	hash := md5.Sum([]byte(a))
	return hex.EncodeToString(hash[:])
}
