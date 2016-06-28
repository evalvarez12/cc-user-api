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
	log.Println("DB ADDRESS: ", os.Getenv("CC_DBADDRESS"))
	host := os.Getenv("CC_DBADDRESS")
	settings := postgresql.ConnectionURL{
		Database: os.Getenv("CC_DBNAME"),
		Host:     host,
		User:     os.Getenv("CC_DBUSER"),
		Password: os.Getenv("CC_DBPASS"),
	}

	// Conexion a la DB y comunicarse con las tables
	var err error
	sess, err = postgresql.Open(settings)
	if err != nil {
		log.Printf("Session Open Error: ", err)
		settings.Host = os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
		sess, err = postgresql.Open(settings)
		if err != nil {
			log.Fatal("Session Open Error: ", err)
		}
	}

	userSource = sess.Collection("users")
}
