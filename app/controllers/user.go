package controllers

import (
	"github.com/evalvarez12/cc-user-api/app/ds"
	"github.com/evalvarez12/cc-user-api/app/models"
	"github.com/revel/revel"
	"encoding/json"
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
	claims, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	userID := claims["id"].(uint)
	jti := claims["jti"].(string)

	err = ds.UserLogout(userID, jti)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}

func (c Users) LogoutAll() revel.Result {
	claims, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	userID := claims["id"].(uint)

	err = ds.UserLogoutAll(userID)
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

	id, err := ds.UserAdd(newUser)
	if err != nil {
		return c.Error(err)
	}
	return c.Data(id)
}

func (c Users) Delete() revel.Result {
	claims, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	userID := claims["id"].(uint)

	err = ds.UserDelete(userID)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}
