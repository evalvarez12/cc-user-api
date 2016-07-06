package ds

import (
	"github.com/arbolista-dev/cc-user-api/app/models"
	"log"
	"testing"
)

var (
	userTest = models.User{}
)

func TestUserSource(t *testing.T) {
	temp, err := userSource.Insert(userTest)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
	id := uint(temp.(int64))

	err = userSource.Find("user_id", id).One(&userTest)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}

	err = userSource.Find("user_id", id).Delete()
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
}
