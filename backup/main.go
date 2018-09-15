package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

// InitializeDBConnection initializes the database connection upon application start.
func InitializeDBConnection(address, database string) {
	session, err := mgo.Dial(address)
	if err != nil {
		logrus.WithError(err).Fatal("Could not initialize mongodb client")
	}
	db = session.DB(database)
}

// HandleIndex handles the root of the domain.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("tmpl/index.html")
	if err != nil {
		logrus.WithError(err).Fatal()
	}
	tmpl.Execute(w, GetAuthenticationDataForUser(r.Cookies()))
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		data := GetAuthenticationDataForUser(r.Cookies())
		if data.IsLoggedIn {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		tmpl, err := template.ParseFiles("tmpl/signup.html")
		if err != nil {
			logrus.WithError(err).Fatal()
		}
		tmpl.Execute(w, data)
	} else {
		r.ParseForm()

		t := SupplierType
		// TODO: HASH PASSWORD
		user := NewUser(r.Form["username"][0], r.Form["password"][0], r.Form["email"][0], r.Form["address"][0], t, bson.NewObjectId(), []InventoryItem{})

		err := InsertNewUser(user)
		if err != nil {
			fmt.Fprint(w, err)
		}

		r.AddCookie(&http.Cookie{Name: "UUID", Value: string(user.UUID)})
		r.AddCookie(&http.Cookie{Name: "TOKEN", Value: user.Token})
		r.AddCookie(&http.Cookie{Name: "TYPE", Value: string(user.Type)})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func main() {
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/signup", HandleSignup)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.WithError(err).Fatal("web server crashed")
	}
}
