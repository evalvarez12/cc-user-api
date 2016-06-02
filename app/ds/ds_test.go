package ds

import (
	"financy/api/app/models"
	"log"
	"testing"
)

var (
	chargeTest   = models.Charge{}
	categoryTest = models.Category{}
	accountTest  = models.Account{}
	userTest     = models.User{}
)

func TestChargeSource(t *testing.T) {
	chargeTest.MarshalDB()
	temp, err := chargeSource.Insert(chargeTest)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
	id := uint(temp.(int64))

	err = chargeSource.Find("id", id).One(&chargeTest)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}

	err = chargeSource.Find("id", id).Delete()
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
}

func TestCategorySource(t *testing.T) {
	temp, err := categorySource.Insert(categoryTest)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
	id := uint(temp.(int64))

	err = categorySource.Find("category_id", id).One(&categoryTest)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}

	err = categorySource.Find("category_id", id).Delete()
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
}

func TestAccountSource(t *testing.T) {
	accountTest.MarshalDB()
	temp, err := accountSource.Insert(accountTest)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
	id := uint(temp.(int64))

	err = accountSource.Find("account_id", id).One(&accountTest)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}

	err = accountSource.Find("account_id", id).Delete()
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
}

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
