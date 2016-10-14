package ds

import (
	"log"
	"os"
	"upper.io/db.v2"
	"upper.io/db.v2/postgresql"
)

var sess postgresql.Database

var (
	userSource     					db.Collection
	leadersSource 					db.Collection
	leadersFoodSource  			db.Collection
	leadersHousingSource  	db.Collection
	leadersShoppingSource  	db.Collection
	leadersTransportSource  db.Collection
	query										db.Result
)

func init() {
	settings := postgresql.ConnectionURL{
		Database: os.Getenv("CC_DBNAME"),
		Host:     "postgres",
		User:     os.Getenv("CC_DBUSER"),
		Password: os.Getenv("CC_DBPASS"),
	}

	// Conexion a la DB y comunicarse con las tables
	var err error
	sess, err = postgresql.Open(settings)
	if err != nil {
		log.Fatal("Session Open Error: ", err)
	}

	userSource = sess.Collection("users")
	leadersSource = sess.Collection("leaders_public_footprint")
	leadersFoodSource = sess.Collection("leaders_public_food_footprint")
	leadersHousingSource = sess.Collection("leaders_public_housing_footprint")
	leadersShoppingSource = sess.Collection("leaders_public_shopping_footprint")
	leadersTransportSource = sess.Collection("leaders_public_transport_footprint")
}
