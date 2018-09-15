package main

import (
	"time"

	"github.com/google/uuid"
)

type MedicationType int

const (
	AcetylsalicylicAcid MedicationType = 1
	Adenosine           MedicationType = 2
	Amiloride           MedicationType = 3
	Amiodarone          MedicationType = 4
	Bisoprolol          MedicationType = 5
	Clopidogrel         MedicationType = 6
	Digoxin     	    MedicationType = 7
)

type Hospital struct {
	UUID        uuid.UUID
	Username    string
	Password    string
	Name        string
	Address     string
	Description string
	PhotoURL    string
	Doctors     []Doctor
}

type Doctor struct {
	UUID        uuid.UUID
	Username    string
	Password    string
	Name        string
	Description string
	Photo       string
}

type 

type Order struct {
	Medication     MedicationType
	Supplier       uuid.UUID
	TargetHospital uuid.UUID
	TargetDoctor   uuid.UUID
	Time           time.Time
	Urgency        int
	Quantity       int
	PricePerUnit   float32
}


