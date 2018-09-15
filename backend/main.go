package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

//InitializeServer initializes the mgo.Database global variable.
func InitializeServer(address, database string) {
	session, err := mgo.Dial(address)
	if err != nil {
		logrus.WithError(err).Fatal("Could not initialize mongo client")
	}
	db = session.DB(database)
}

func NewHospitalEndPoint(w http.ResponseWriter, r *http.Request) {

}

func NewDoctorEndPoint(w http.ResponseWriter, r *http.Request) {

}

func NewSupplierEndPoint(w http.ResponseWriter, r *http.Request) {

}

func GetMedicineEndPoint(w http.ResponseWriter, r *http.Request) {

}

func GetUserEndPoint()

func main() {

	InitializeServer("mongodb://htnadmin:htn2018@db.medwithoutborders.org:27017", "htn")

	router := mux.NewRouter()

	router.HandleFunc("/hospital/new", NewHospitalEndPoint).Methods("POST")
	router.HandleFunc("/doctor/new", NewDoctorEndPoint).Methods("POST")
	router.HandleFunc("/supplier/new", NewSupplierEndPoint).Methods("POST")

	router.handleFunc("/user/{id}")

	router.HandleFunc("/medicine/{id}", GetMedicineEndPoint).Methods("GET")
	router.HandleFunc("/user/{id}")

	/*
		get all medicine sources
		get user data
	*/

	// err := InsertNewOrder(Order{UUID: bson.NewObjectId(), Supplier: bson.NewObjectId(), TargetHospital: bson.NewObjectId(), TargetDoctor: bson.NewObjectId(), Medication: Aspirin, Time: time.Now(), Urgency: 10, Quantity: 100, PricePerUnit: 10.3})
	// if err != nil {
	// 	logrus.WithError(err).Fatal("Could not insert new order")
	// }

	// orders, err := GetAllOrders()
	// if err != nil {
	// 	logrus.WithError(err).Fatal("Could not retrieve all orders")
	// }

	// logrus.WithFields(logrus.Fields{"Orders": orders}).Info("The returned orders")

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
