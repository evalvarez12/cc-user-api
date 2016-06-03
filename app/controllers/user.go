package controllers

import (
	"encoding/json"
	"financy/api/app/ds"
	"financy/api/app/models"
	"financy/api/app/services"
	"financy/api/app/jobs"
	"github.com/revel/revel"
	"io/ioutil"
)

type Users struct {
	App
}

func (c Users) Login() revel.Result {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.Error(err)
	}
	var logRequest models.UserLogin
	err = json.Unmarshal(body, &logRequest)
	if err != nil {
		return c.Error(err)
	}
	login, err := ds.UserLogin(logRequest)
	if err != nil {
		return c.Error(err)
	}

	return c.Data(login)
}

func (c Users) Logout() revel.Result {
	userID, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	err = ds.UserLogout(userID)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}

func (c Users) Add() revel.Result {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.Error(err)
	}
	var newUser models.User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		return c.Error(err)
	}
	newUser.Validate(c.Validation)
	if c.Validation.HasErrors() {
		errors := c.Validation.ErrorMap()
		return c.ErrorData(errors)
	}
	newUser.NewFill()

	id, err := ds.UserAdd(newUser)
	if err != nil {
		return c.Error(err)
	}
	err = categorySamples(id)
	if err != nil {
		return c.Error(err)
	}

	err = services.SendMail(newUser.Email, newUser.Name)
	if err != nil {
		return c.Error(err)
	}

	return c.OK()
}

func (c Users) SampleUser(n uint) revel.Result {
	user := models.GenerateUser()
	userID, err := ds.UserAdd(user)
	if err != nil {
		return c.Error(err)
	}

	err = categorySamples(userID)
	if err != nil {
		return c.Error(err)
	}

	err = accountSamples(userID)
	if err != nil {
		return c.Error(err)
	}

	if n == 0 {
		n = 100
	}
	err = chargesSamples(n, userID)
	if err != nil {
		return c.Error(err)
	}
	err = plannedSamples(userID)
	if err != nil {
		return c.Error(err)
	}

	return c.OK()

}
