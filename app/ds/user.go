package ds

import (
	"crypto/rand"
	"errors"
	"financy/api/app/models"
	"golang.org/x/crypto/bcrypt"
	"upper.io/db.v2"
)

func GetSession(token string) (userID uint, err error) {
	var sess models.Session
	err = sessionSource.Find("token", token).One(&sess)
	if err != nil {
		err = errors.New("Non existent session")
		return
	}
	userID = sess.UserID
	return
}

func UserAdd(user models.User) (userID uint, err error) {
	hashPassword(&user)

	temp, err := userSource.Insert(user)
	if err != nil {
		return
	}
	userID = uint(temp.(int64))
	defaultAcc := models.DefaultAccount()
	defaultAcc.UserID = userID
	id, err := AccountsAdd(defaultAcc)
	if err != nil {
		err = errors.New("Error setting up default account")
		return
	}
	user.DefaultAccount = id
	err = userSource.Find(db.Cond{"user_id": userID}).Update(user)
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
	token := createToken(10)

	userSession := models.Session{
		UserID: user.UserID,
		Token:  token,
	}

	_, err = sessionSource.Insert(userSession)
	if err != nil {
		return
	}

	login = map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"default-account": user.DefaultAccount,
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
