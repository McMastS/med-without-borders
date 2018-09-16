package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

// InitialSourceData holds data
type InitialSourceData struct {
	Username    string `bson:"username" json:"username"`
	Password    string `bson:"password" json:"password"`
	Name        string `bson:"name" json:"name"`
	Address     string `bson:"address" json:"address"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
	PhotoURL    string `bson:"photo_url" json:"photo_url"`
	Inventory   string `bson:"inventory" json:"inventory"`
	Prices      string `bson:"prices" json:"prices"`
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
	PricePerUnit float64      `bson:"price_per_unit" json:"price_per_unit"`
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

	inventory := []InventoryItem{}

	for i, str := range strings.Split(sourceData.Inventory, " ") {
		quantity, err := strconv.Atoi(str)
		if err != nil {
			logrus.WithError(err).Fatal()
		}

		if quantity > 0 {

			price, err := strconv.ParseFloat(strings.Split(sourceData.Prices, " ")[i], 32)
			if err != nil {
				logrus.WithError(err).Fatal()
			}

			inventory = append(inventory, InventoryItem{ID: MedicationID(i + 1), Quantity: quantity, PricePerUnit: price})
		}
	}

	return Source{bson.NewObjectId(), string(GenerateRandomBytes(32)), sourceData.Username, pass, sourceData.Name, sourceData.Address, sourceData.PhoneNumber, sourceData.PhotoURL, inventory}
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	} else {

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
		response := SourceCreationResponse{Response: 1, UUID: source.UUID.Hex(), Token: source.Token}

		err = InsertNewUserIntoDB(source)
		if err != nil {
			logrus.WithError(err).Fatal()
		}

		json, err := json.Marshal(response)
		if err != nil {
			logrus.WithError(err).Fatal()
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
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

type SupplierDataRequest struct {
	ID           int    `bson:"id" json:"id"`
	LoggedInUUID string `bson:"uuid" json:"uuid"`
}

type SupplierData struct {
	ID       string  `bson:"id" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Quantity int     `bson:"quantity" json:"quantity"`
	Price    float64 `bson:"price_per_unit" json:"price_per_unit"`
	Distance float64 `bson:"distance" json:"distance"`
}

func GatherSupplierDataForMedicine(dr SupplierDataRequest) ([]SupplierData, error) {
	var sources []Source
	err := db.C("sources").Find(bson.M{}).All(&sources)

	if err != nil {
		return []SupplierData{}, nil
	}

	loggedUser, err := GetUserForUUID(dr.LoggedInUUID)
	if err != nil {
		return []SupplierData{}, nil
	}

	x1, x2, err := GetCoordinatesFromAddress(loggedUser.Address)

	var supplierData []SupplierData

	for _, s := range sources {
		for _, i := range s.Inventory {
			if i.ID == MedicationID(dr.ID) {

				supplierUser, err := GetUserForUUID(s.UUID.Hex())
				if err != nil {
					continue
				}

				y1, y2, err := GetCoordinatesFromAddress(supplierUser.Address)

				distance := math.Sqrt((y1-x1)*(y1-x1) + (y2-x2)*(y2-x2))

				supplierData = append(supplierData, SupplierData{s.UUID.Hex(), s.Name, i.Quantity, i.PricePerUnit, distance})
				break
			}
		}
	}

	return supplierData, nil
}

func GetMedicineEndpoint(w http.ResponseWriter, r *http.Request) {
	var supplierDataRequest SupplierDataRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&supplierDataRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	suppliers, err := GatherSupplierDataForMedicine(supplierDataRequest)
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

type Order struct {
	UUID         bson.ObjectId `bson:"_id" json:"id"`
	Type         MedicationID  `bson:"med_id" json:"med_id"`
	Supplier     bson.ObjectId `bson:"supplier" json:"supplier"`
	Target       bson.ObjectId `bson:"target" json:"target"`
	Time         time.Time     `bson:"time" json:"time"`
	Quantity     int           `bson:"quantity" json:"quantity"`
	PricePerUnit float32       `bson:"price_per_unit" json:"price_per_unit"`
}

func GetOrdersEndpoint(w http.ResponseWriter, r *http.Request) {
	var orders []Order
	err := db.C("orders").Find(bson.M{}).All(&orders)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, string(json))
}

type NewOrder struct {
	Type         MedicationID  `bson:"med_id" json:"med_id"`
	Supplier     bson.ObjectId `bson:"supplier" json:"supplier"`
	Target       bson.ObjectId `bson:"target" json:"target"`
	Quantity     int           `bson:"quantity" json:"quantity"`
	PricePerUnit float32       `bson:"price_per_unit" json:"price_per_unit"`
}

func CreateOrder(newOrder NewOrder) Order {
	return Order{bson.NewObjectId(), newOrder.Type, newOrder.Supplier, newOrder.Target, time.Now(), newOrder.Quantity, newOrder.PricePerUnit}
}

func InsertNewOrder(order Order) error {
	return db.C("orders").Insert(&order)
}

type OrderResponse struct {
	Response int `bson:"response" json:"response"`
}

func AddOrderEndpoint(w http.ResponseWriter, r *http.Request) {
	var newOrder NewOrder
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	order := CreateOrder(newOrder)

	supplier, err := GetUserForUUID(string(order.Supplier))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	valid := false

	for _, i := range supplier.Inventory {
		if i.ID == order.Type && i.Quantity >= order.Quantity {
			valid = true
			break
		}
	}

	if valid {
		json, err := json.Marshal(OrderResponse{1})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, string(json))
	} else {
		json, err := json.Marshal(OrderResponse{0})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, string(json))
	}
}

// APP ID: mBIRDYvCDBwBva0ZHTKh
// APP CODE: z5-x3UWZPP3NlwVtiZN42g

func GetCoordinatesFromAddress(address string) (float64, float64, error) {
	var client = &http.Client{Timeout: 10 * time.Second}
	url := "https://geocoder.api.here.com/6.2/geocode.json?app_id=mBIRDYvCDBwBva0ZHTKh&app_code=z5-x3UWZPP3NlwVtiZN42g&searchtext="
	add := strings.Replace(address, " ", "+", -1)
	add = strings.Replace(add, ",", "", -1)
	r, err := client.Get(url + add)
	if err != nil {
		return 0, 0, err
	}

	defer r.Body.Close()

	if r.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return 0, 0, err
		}
		str := string(bytes)

		i := strings.Index(str, "\"Latitude\":")
		s := ""
		for str[i] != ',' {

			if str[i] == '.' {
				s += string(str[i])
				i++
				continue
			}

			if str[i] == '-' {
				s += string(str[i])
				i++
				continue
			}

			if _, err := strconv.Atoi(string(str[i])); err == nil {
				s += string(str[i])
			}
			i++
		}

		latitude, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return 0, 0, err
		}

		i = strings.Index(str, "\"Longitude\":")
		s = ""
		for str[i] != ',' {

			if str[i] == '.' {
				s += string(str[i])
				i++
				continue
			}

			if str[i] == '-' {
				s += string(str[i])
				i++
				continue
			}

			if _, err := strconv.Atoi(string(str[i])); err == nil {
				s += string(str[i])
			}
			i++
		}
		longitude, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return 0, 0, err
		}

		return latitude, longitude, nil
	}

	return 0, 0, nil
}

func main() {

	InitializeServer("mongodb://htnadmin:htn2018@db.medwithoutborders.org:27017", "htn")

	router := mux.NewRouter()

	router.HandleFunc("/source/new_user", NewSourceEndpoint)

	// router.HandleFunc("/source/new_user", NewSourceEndpoint).Methods("POST")
	router.HandleFunc("/source/login_session", LoginSessionEndpoint).Methods("POST")
	router.HandleFunc("/source/login_normal", LoginNormalEndpoint).Methods("POST")

	router.HandleFunc("/medicine/", GetMedicineEndpoint).Methods("POST")

	router.HandleFunc("/orders", GetOrdersEndpoint).Methods("GET")
	router.HandleFunc("/orders/add", AddOrderEndpoint).Methods("POST")

	handler := cors.Default().Handler(router)

	http.ListenAndServe(":80", handler)
}
