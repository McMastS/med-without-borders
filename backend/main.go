package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func InitializeServer(address string)

func main() {

	session, err := mgo.Dial("mongodb://htnadmin:htn2018@db.medwithoutborders.org:27017")
	if err != nil {
		logrus.WithError(err).Fatal("Could not initialize mongo client")
	}

	db = session.DB("users")

	// client, err := mongo.NewClient("mongodb://htnadmin:htn2018@db.medwithoutborders.org:27017")
	// if err != nil {
	// 	logrus.WithError(err).Fatal("Could not initialize mongo client")
	// }

	// err = client.Connect(context.TODO())
	// if err != nil {
	// 	logrus.WithError(err).Fatal("Could not connect to mongo client")
	// }

	// collection := client.Database("users").Collection("hospitals")

	// res, err := collection.InsertOne(context.Background(), map[string]string{"hello": "world"})
	// if err != nil {
	// 	logrus.WithError(err).Fatal("Could not insert document")
	// }

	// id := res.InsertedID

	// logrus.WithFields(logrus.Fields{"ID": id}).Info("Successfully Inserted Document")

}
