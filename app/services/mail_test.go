package services

import (
	// "github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTemplateMail(t *testing.T) {
	res, err := templateMail("eduardo@mazing.studio", "prueba")
	if err != nil {
		log.Print(err)
	}
	log.Print(string(res[:]))
}

func TestSendMail(t *testing.T) {
	err := SendMail("eduardo@mazing.studio", "prueba")
	if err != nil {
		log.Print(err)
	}
	log.Print("success")
}
