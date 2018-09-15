package main

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
)

// UserType defines an enum for differentiating between different user types.
type UserType int

const (
	// HospitalType is a Hospital User
	HospitalType UserType = 1
	// DoctorType is a Doctor User
	DoctorType UserType = 2
	// SupplierType is a Supplier User
	SupplierType UserType = 3
)

// AuthenticationData provides a way to transmit data to the frontend in a safe way.
type AuthenticationData struct {
	IsLoggedIn bool `bson:"is_logged_in" json:"is_logged_in"`
	AuthData   User `bson:"data" json:"data"`
}

// User holds information related to a user.
type User struct {
	UUID      bson.ObjectId   `bson:"_id" json:"id"`
	Username  string          `bson:"username" json:"username"`
	Password  string          `bson:"password" json:"password"`
	Email     string          `bson:"email" json:"email"`
	Token     string          `bson:"token" json:"token"`
	Type      UserType        `bson:"type" json:"type"`
	Address   string          `bson:"address" json:"address"`
	Owner     bson.ObjectId   `bson:"owner" json:"owner"`
	Inventory []InventoryItem `bson:"inventory" json:"inventory"`
}

// NewUser creates a new User object.
func NewUser(Username, Password, Email, Address string, Type UserType, Owner bson.ObjectId, Inventory []InventoryItem) User {
	return User{bson.NewObjectId(), Username, Password, Email, string(GenerateRandomBytes(16)), Type, Address, Owner, Inventory}
}

// InventoryItem represents medicine currently in a supplier's inventory.
type InventoryItem struct {
	Type         int     `bson:"type" json:"type"`
	Quantity     int     `bson:"quantity" json:"quantity"`
	PricePerUnit float32 `bson:"price_per_unit" json:"price_per_unit"`
}

// GetCollectionForType is a utility function which returns a corresponding string for a UserType.
func GetCollectionForType(t UserType) string {
	switch t {
	case HospitalType:
		return "hospitals"
	case DoctorType:
		return "doctors"
	case SupplierType:
		return "suppliers"
	default:
		logrus.Fatal("Invalid User Type")
		return ""
	}
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

// GetUserForUUID returns user data for a specified UUID.
func GetUserForUUID(t UserType, uuid string) (User, error) {

	collection := GetCollectionForType(t)
	var user User

	err := db.C(collection).Find(bson.M{"_id": uuid}).One(&user)

	if err != nil {
		return User{}, errors.New("user cannot be located")
	}

	return user, nil
}

// // DoesUserWithEmailExist returns user data for a specifed email
// func DoesUserWithEmailExist(t UserType, email string) bool {
// 	fmt.Println("w")
// 	collection := GetCollectionForType(t)
// 	fmt.Println("o")
// 	count, err := db.C(collection).Find(bson.M{"email": email}).Count()
// 	fmt.Println("w")
// 	return !(err != nil || count == 0)
// }

// GetAuthenticationDataForUser returns a set of authentication data for a user.
func GetAuthenticationDataForUser(cookies []*http.Cookie) AuthenticationData {

	cs := make(map[string]string)

	for _, c := range cookies {
		cs[c.Name] = c.Value
	}

	if cs["UUID"] != "" && cs["TOKEN"] != "" && cs["TYPE"] != "" {

		t, err := strconv.Atoi(cs["TYPE"])

		if err != nil {
			return AuthenticationData{false, User{}}
		}

		user, err := GetUserForUUID(UserType(t), cs["UUID"])
		if err != nil {
			return AuthenticationData{false, User{}}
		}

		if user.Token == cs["TOKEN"] {
			return AuthenticationData{true, user}
		}
	}

	return AuthenticationData{false, User{}}
}

// InsertNewUser adds a new user to the database
func InsertNewUser(user User) error {
	// logrus.Info("hi")
	// exists := DoesUserWithEmailExist(user.Type, user.Email)
	// logrus.Info("bue")
	// if exists {
	// 	logrus.Info("exists")
	// }

	return db.C(GetCollectionForType(user.Type)).Insert(user)
}
