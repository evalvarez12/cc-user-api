package controllers

import (
	"encoding/json"
	"github.com/evalvarez12/cc-user-api/app/ds"
	"github.com/evalvarez12/cc-user-api/app/models"
	"github.com/evalvarez12/cc-user-api/app/services"
	"github.com/revel/revel"
	"io/ioutil"
	"net/url"
	"strconv"
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
	login, err := ds.Login(logRequest)
	if err != nil {
		return c.Error(err)
	}

	return c.Data(login)
}

func (c Users) Logout() revel.Result {
	userID, jti, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	err = ds.Logout(userID, jti)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}

func (c Users) LogoutAll() revel.Result {
	userID, _, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	err = ds.LogoutAll(userID)
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

	id, err := ds.Add(newUser)
	if err != nil {
		return c.Error(err)
	}

	data := map[string]string{"name": newUser.Name}
	err = services.SendMail("new-user-beta", newUser.Email, data)
	if err != nil {
		return c.Error(err)
	}

	return c.Data(id)
}

func (c Users) Delete() revel.Result {
	userID, _, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	err = ds.Delete(userID)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}

func (c Users) UpdateAnswers() revel.Result {
	userID, _, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.Error(err)
	}
	var bodyAnswers models.Answers
	err = json.Unmarshal(body, &bodyAnswers)
	if err != nil {
		return c.Error(err)
	}

	err = ds.UpdateAnswers(userID, bodyAnswers)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}

func (c Users) Update() revel.Result {
	userID, _, err := c.GetSession()
	if err != nil {
		return c.Error(err)
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.Error(err)
	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return c.Error(err)
	}

	err = ds.Update(userID, user)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}

func (c Users) PassResetRequest() revel.Result {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return c.Error(err)
	}

	var email models.Email
	err = json.Unmarshal(body, &email)
	if err != nil {
		return c.Error(err)
	}

	userID, token, err := ds.PassResetRequest(email.Email)
	if err != nil {
		return c.Error(err)
	}

	data := map[string]string{"link": PasswordResetURL(userID, token)}
	err = services.SendMail("passwords-reset", email.Email, data)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}

func (c Users) PassResetConfirm(userID uint, token, password string) revel.Result {
	err := ds.PassResetConfirm(userID, token, password)
	if err != nil {
		return c.Error(err)
	}
	return c.OK()
}

func PasswordResetURL(userID uint, token string) (uri string) {
	u := url.URL{}
	u.Scheme = "http"
	u.Host = "localhost:8082"
	u.Path = "/page/passreset"
	q := u.Query()
	q.Set("id", strconv.Itoa(int(userID)))
	q.Set("token", token)
	u.RawQuery = q.Encode()
	return u.String()
}
