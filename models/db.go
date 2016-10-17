package models

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

// ConnectDB starts the database session and shares the instance with all models
func ConnectDB(dbAddress string, dbName string) {
	session, err := mgo.Dial(dbAddress)
	if err != nil {
		fmt.Printf("Failed to connect to database server at `%s`\n", dbAddress)
		panic(err)
	}
	db = session.DB(dbName)
}

// DisconnectDB terminates the database session
func DisconnectDB() {
	session.Close()
}
