package main

import (
	"errors"
	"math/rand"

	"gopkg.in/mgo.v2/bson"
)

const characterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomBytes creates a random token of length n
func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = characterBytes[rand.Intn(len(characterBytes))]
	}
	return b
}

// GetHospitalForID returns a hospital object for a given ID if it exists
func GetHospitalForID(id bson.ObjectId) (Hospital, error) {
	var hospital Hospital
	err := db.C("hospitals").FindId(id).One(&hospital)
	return hospital, err
}

// GetDoctorForID returns a doctor given a hospital and a doctor ID
func GetDoctorForID(hospital, doctor bson.ObjectId) (Doctor, error) {
	h, err := GetHospitalForID(hospital)
	if err != nil {
		return Doctor{}, err
	}

	for _, d := range h.Doctors {
		if d.UUID == doctor {
			return d, nil
		}
	}

	return Doctor{}, errors.New("cannot find doctor")
}

// GetAllOrders returns all orders in the database
func GetAllOrders() ([]Order, error) {
	var orders []Order
	err := db.C("orders").Find(bson.M{}).All(&orders)
	return orders, err
}

// InsertNewOrder pushes a new order to the database.
func InsertNewOrder(order Order) error {
	return db.C("orders").Insert(&order)
}
