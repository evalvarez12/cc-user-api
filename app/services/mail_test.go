package services

import (
	// "github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBetaMail(t *testing.T) {
	data := map[string]string{"name": "prueba"}
	res, err := templateMail("new-user-beta", "eduardo@mazing.studio", data)
	if err != nil {
		log.Print(err)
	}
	log.Print(string(res[:]))
}

func TestSendMail(t *testing.T) {
	data := map[string]string{"name": "prueba"}
	err := SendMail("new-user-beta", "eduardo@mazing.studio", data)
	if err != nil {
		log.Print(err)
	}
	log.Print("success")
}
