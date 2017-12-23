package db

import (
	"gopkg.in/mgo.v2"
	"time"
)

type MongoServerConfig struct {
	URL              string        `json:"url"`
	AuthDatabaseName string        `json:"authDatabaseName"`
	User             string        `json:"user"`
	Password         string        `json:"password"`
	DatabaseName     string        `json:"dataBaseName"`
}

type DB struct {
	Session *mgo.Session
}

func NewDB(configuration MongoServerConfig) (*DB, error) {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{configuration.URL},
		Timeout:  60 * time.Second,
		Database: configuration.AuthDatabaseName,
		Username: configuration.User,
		Password: configuration.Password,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	session, err := mgo.DialWithInfo(mongoDBDialInfo)

	//	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &DB{Session:session}, nil
}

func (db *DB) Close() {
	db.Session.Close()
}