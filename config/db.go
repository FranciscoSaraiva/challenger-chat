package config

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// DB represents the Mongo Database
var DB *mgo.Database

// DBName represents the name of the Database.
var DBName = "OLXChatDB"
// DBTestName represents the test database for running go tests.
var DBTestName = "OLXChatDB-TEST"

// Initializer
func init() {
	fmt.Println("************")
	fmt.Println("FIRING IT UP -> " + DBTestName)
	fmt.Println("************")
	// get a mongo session
	s, err := mgo.Dial("mongodb://127.0.0.1:27017/" + DBTestName)
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB(DBTestName)

	if DB.Name == DBTestName {
		DB.DropDatabase()
	}

	log.Println("[Database] You connected to your mongo database.")
}

// GetCollection function that returns a collection in the database,
// that corresponds to the name inside the database, passed in the string.
func GetCollection(name string) *mgo.Collection {
	return DB.C(name)
}
