package ds

import (
	"log"
	"os"
	"upper.io/db.v2"
	"upper.io/db.v2/postgresql"
)

var sess postgresql.Database

var (
	chargeSource   db.Collection
	userSource     db.Collection
	categorySource db.Collection
	accountSource  db.Collection
	sessionSource  db.Collection
	plannedSource  db.Collection
)

func init() {
	log.Println("CCUSER", os.Getenv("CC_DBUSER"))
	settings := postgresql.ConnectionURL{
		Database: os.Getenv("CC_DBNAME"),
		Host:     os.Getenv("CC_DBADDRESS"),
		User:     os.Getenv("CC_DBUSER"),
		Password: os.Getenv("CC_DBPASS"),
	}

	// Conexion a la DB y comunicarse con las tables
	var err error
	sess, err = postgresql.Open(settings)
	// sess, err = db.Open(postgresql.Adapter, settings)
	if err != nil {
		log.Fatal("Session Open(): ", err)
	}

	userSource = sess.Collection("users")
}
