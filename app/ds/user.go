package ds

import (
	"crypto/rand"
	"errors"
	"financy/api/app/models"
	"golang.org/x/crypto/bcrypt"
	"upper.io/db.v2"
)

func UserAdd(user models.User) (userID uint, err error) {
	hashPassword(&user)

	temp, err := userSource.Insert(user)
	if err != nil {
		return
	}
	userID = uint(temp.(int64))
	return
}

func UserLogin(logRequest models.UserLogin) (login map[string]interface{}, err error) {
	var user models.User
	err = userSource.Find("name", logRequest.Name).One(&user)
	if err != nil {
		err = errors.New("Incorrect Password or UserName")
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Hash, append([]byte(logRequest.Password), user.Salt...))
	if err != nil {
		err = errors.New("Incorrect Password or UserName")
		return
	}
	token, err := newToken(user.UserID)
	if err != nil {
		return
	}

	login = map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"token": token,
	}
	return
}

func UserLogout(userID uint) (err error) {
	
	err = sessionSource.Find("user_id", userID).Delete()
	return
}

func hashPassword(user *models.User) {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	user.Hash, err = bcrypt.GenerateFromPassword(append([]byte(user.Password), b...), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Salt = b
}
