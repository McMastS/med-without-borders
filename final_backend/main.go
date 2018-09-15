package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MedicationID is a unique ID for each type of supported medication.
type MedicationID int

const (
	Aspirin        MedicationID = 1
	Amiloride      MedicationID = 2
	Amiodarone     MedicationID = 3
	Bisoprolol     MedicationID = 4
	Clopidogrel    MedicationID = 5
	Digoxin        MedicationID = 6
	Furosemide     MedicationID = 7
	Losartan       MedicationID = 8
	Methyldopa     MedicationID = 9
	Nifedipine     MedicationID = 10
	Spironolactone MedicationID = 11
	Streptokinase  MedicationID = 12
	Verapamil      MedicationID = 13
)

type InitialSourceData struct {
	Username    string          `bson:"username" json:"username"`
	Password    string          `bson:"password" json:"password"`
	Name        string          `bson:"name" json:"name"`
	Address     string          `bson:"address" json:"address"`
	PhoneNumber string          `bson:"phone_number" json:"phone_number"`
	PhotoURL    string          `bson:"photo_url" json:"photo_url"`
	Inventory   []InventoryItem `bson:"inventory" json:"inventory"`
}

type Source struct {
	UUID        bson.ObjectId   `bson:"_id" json:"id"`
	Token       string          `bson:"token" json:"token"`
	Username    string          `bson:"username" json:"username"`
	Password    string          `bson:"password" json:"password"`
	Name        string          `bson:"name" json:"name"`
	Address     string          `bson:"address" json:"address"`
	PhoneNumber string          `bson:"phone_number" json:"phone_number"`
	PhotoURL    string          `bson:"photo_url" json:"photo_url"`
	Inventory   []InventoryItem `bson:"inventory" json:"inventory"`
}

type InventoryItem struct {
	ID           MedicationID `bson:"med_id" json:"med_id"`
	Quantity     int          `bson:"quantity" json:"quantity"`
	PricePerUnit float32      `bson:"price_per_unit" json:"price_per_unit"`
}

var db *mgo.Database

//InitializeServer initializes the mgo.Database global variable.
func InitializeServer(address, database string) {
	session, err := mgo.Dial(address)
	if err != nil {
		logrus.WithError(err).Fatal("Could not initialize mongo client")
	}
	db = session.DB(database)
}

const characterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomBytes creates a random token of length n
func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = characterBytes[rand.Intn(len(characterBytes))]
	}
	return b
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateNewUser(sourceData InitialSourceData) Source {
	pass, err := HashPassword(sourceData.Password)
	if err != nil {
		logrus.WithError(err).Fatal()
	}
	return Source{bson.NewObjectId(), string(GenerateRandomBytes(32)), sourceData.Username, pass, sourceData.Name, sourceData.Address, sourceData.PhoneNumber, sourceData.PhotoURL, sourceData.Inventory}
}

func InsertNewUserIntoDB(data Source) error {
	return db.C("sources").Insert(&data)
}

type SourceCreationResponse struct {
	Response int    `bson:"response" json:"response"`
	UUID     string `bson:"uuid" json:"uuid"`
	Token    string `bson:"token" json:"token"`
}

func NewSourceEndpoint(w http.ResponseWriter, r *http.Request) {
	var sourceData InitialSourceData
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&sourceData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	source := CreateNewUser(sourceData)
	response := SourceCreationResponse{Response: 1, UUID: string(source.UUID), Token: source.Token}

	err = InsertNewUserIntoDB(source)
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	json, err := json.Marshal(response)
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	fmt.Fprintf(w, string(json))
}

func GetUserForUUID(uuid string) (Source, error) {
	var source Source
	err := db.C("sources").Find(bson.M{"_id": uuid}).One(&source)
	if err != nil {
		return Source{}, err
	}
	return source, nil
}

type SessionLoginRequest struct {
	UUID  string `bson:"uuid" json:"uuid"`
	Token string `bson:"token" json:"token"`
}

type SessionLoginResponse struct {
	Response int    `bson:"response" json:"response"`
	Name     string `bson:"name" json:"name"`
}

func LoginSessionEndpoint(w http.ResponseWriter, r *http.Request) {
	var sessionLoginRequest SessionLoginRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&sessionLoginRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := GetUserForUUID(sessionLoginRequest.UUID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if user.Token == sessionLoginRequest.Token {
		json, err := json.Marshal(SessionLoginResponse{1, user.Name})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, string(json))
	} else {
		json, err := json.Marshal(SessionLoginResponse{0, ""})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, string(json))
	}

}

type NormalLoginRequest struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

func GetUserForUsername(username string) (Source, error) {
	var source Source
	err := db.C("sources").Find(bson.M{"username": username}).One(&source)
	if err != nil {
		return Source{}, err
	}
	return source, nil
}

func LoginNormalEndpoint(w http.ResponseWriter, r *http.Request) {
	var normalLoginRequest NormalLoginRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&normalLoginRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := GetUserForUsername(normalLoginRequest.Username)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if CheckPasswordHash(normalLoginRequest.Password, user.Password) {
		json, err := json.Marshal(SessionLoginResponse{1, user.Name})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, string(json))
	} else {
		json, err := json.Marshal(SessionLoginResponse{0, ""})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, string(json))
	}
}

type SupplierData struct {
	ID       string  `bson:"id" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Quantity int     `bson:"quantity" json:"quantity"`
	Price    float32 `bson:"price_per_unit" json:"price_per_unit"`
}

func GatherSupplierDataForMedicine(id MedicationID) ([]SupplierData, error) {
	var sources []Source
	err := db.C("sources").Find(bson.M{}).All(&sources)

	if err != nil {
		return []SupplierData{}, nil
	}

	var supplierData []SupplierData

	for _, s := range sources {
		for _, i := range s.Inventory {
			if i.ID == id {
				supplierData = append(supplierData, SupplierData{string(s.UUID), s.Name, i.Quantity, i.PricePerUnit})
				break
			}
		}
	}

	return supplierData, nil
}

func GetMedicineEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idNum, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id := MedicationID(idNum)

	suppliers, err := GatherSupplierDataForMedicine(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	data, err := json.Marshal(suppliers)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, string(data))
}

func main() {
	InitializeServer("mongodb://htnadmin:htn2018@db.medwithoutborders.org:27017", "htn")

	router := mux.NewRouter()

	router.HandleFunc("/source/new_user", NewSourceEndpoint).Methods("POST")
	router.HandleFunc("/source/login_session", LoginSessionEndpoint).Methods("POST")
	router.HandleFunc("/source/login_normal", LoginNormalEndpoint).Methods("POST")

	router.HandleFunc("/medicine/{id}", GetMedicineEndpoint).Methods("GET")

	http.ListenAndServe(":80", router)
}
