package main

import "gopkg.in/mgo.v2/bson"

type SupplierOverview struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Address  string        `bson:"address" json:"address"`
	Quantity int           `bson:"quantity" json:"quantity"`
	Price    float32       `bson:"price_per_unit" json:"price_per_unit"`
}

func GetSuppliersForMedicineType(mt MedicationType) ([]SupplierOverview, error) {
	var supplierData []Supplier
}
