package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Hospital struct {
	UUID        bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"username" json:"username"`
	Password    string        `bson:"password" json:"password"`
	Name        string        `bson:"name" json:"name"`
	Email       string        `bson:"email" json:"email"`
	Address     string        `bson:"address" json:"address"`
	Description string        `bson:"description" json:"description"`
	Doctors     []Doctor      `bson:"doctors" json:"doctors"`
}

type Doctor struct {
	UUID     bson.ObjectId `bson:"_id" json:"id"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
	Token    string        `bson:"token" json:"token"`
	Name     string        `bson:"name" json:"name"`
}

type Supplier struct {
	UUID      bson.ObjectId   `bson:"_id" json:"id"`
	Username  string          `bson:"username" json:"username"`
	Password  string          `bson:"password" json:"password"`
	Token     string          `bson:"token" json:"token"`
	Name      string          `bson:"name" json:"name"`
	Address   string          `bson:"address" json:"address"`
	Inventory []InventoryItem `bson:"inventory" json:"inventory"`
}

func NewSupplier(Username, Password, Name, Address string) Supplier {
	token := string(GenerateRandomBytes(32))
	return Supplier{bson.NewObjectId(), Username, Password, token, Name, Address, []InventoryItem{}}
}

type InventoryItem struct {
	Type         MedicationType `bson:"medication_type" json:"medication_type"`
	Quantity     int            `bson:"quantity" json:"quantity"`
	PricePerUnit float32        `bson:"price_per_unit" json:"price_per_unit"`
}

func NewInventoryItem(t MedicationType, quantity int, price float32) InventoryItem {
	return InventoryItem{t, quantity, price}
}

type Order struct {
	UUID           bson.ObjectId  `bson:"_id" json:"id"`
	Medication     MedicationType `bson:"medication_type" json:"medication_type"`
	Supplier       bson.ObjectId  `bson:"supplier_id" json:"supplier_id"`
	TargetHospital bson.ObjectId  `bson:"target_hospital" json:"target_hospital"`
	TargetDoctor   bson.ObjectId  `bson:"target_doctor" json:"target_doctor"`
	Time           time.Time      `bson:"time" json:"time"`
	Urgency        int            `bson:"urgency" json:"urgency"`
	Quantity       int            `bson:"quantity" json:"quantity"`
	PricePerUnit   float32        `bson:"price_per_unit" json:"price_per_unit"`
}
