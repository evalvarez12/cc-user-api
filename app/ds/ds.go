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
	// Name of the linked repo
	log.Printf("Database Container Name: %s", os.Getenv("POSTGRES_NAME"))

	// Obtener la direccion de la DBs
	dbAddress := os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	if dbAddress == "" {
		log.Printf("ERROR LINKING CONTAINERS")
		// For testing on local machine
		dbAddress = "localhost"
	}
	log.Printf("Database Container Address: %s", dbAddress)

	settings := postgresql.ConnectionURL{
		Database: `cc_users`,
		Host:     dbAddress,
		User:     `cc`,
		Password: `pass`,
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
