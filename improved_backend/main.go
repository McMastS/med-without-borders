package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func InitializeServer(address, database string) {
	session, err := mgo.Dial(address)
	if err != nil {
		logrus.WithError(err).Fatal("could not initialize mongo client")
	}
	db = session.DB(database)
}

func main() {

}
