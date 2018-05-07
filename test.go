package main

import (
	"log"
	"net/http"
	"strconv"

	. "./config"
	"./factories/response"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Car struct {
	Name string
	Type string
}

var Session mgo.Session

func main() {
	// Database connection
	Session, err := mgo.Dial(Config.Db)
	if nil != err {
		panic(err)
	}
	defer Session.Close()
	Session.SetMode(mgo.Monotonic, true)

	db := Session.DB("car-db")

	// Database name and collection name
	// car-db is database name car is collation name
	c := db.C("car")
	c.Insert(&Car{"Audi", "Luxury car"})

	// Get /car
	http.HandleFunc("/car", func(w http.ResponseWriter, r *http.Request) {
		result := Car{}
		err = c.Find(bson.M{"name": "Audi"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		responseFactory.Ok(result, w)
	})

	http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
}
