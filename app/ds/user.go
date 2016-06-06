package ds

import (
	"crypto/rand"
	"errors"
	"github.com/evalvarez12/cc-user-api/app/models"
	"golang.org/x/crypto/bcrypt"
)

func UserAdd(user models.User) (userID uint, err error) {
	hashPassword(&user)

	user.MarshalDB()
	temp, err := userSource.Insert(user)
	if err != nil {
		return
	}
	userID = uint(temp.(int64))
	return
}

func UserDelete(userID uint) (err error) {
	err = userSource.Find("user_id", userID).Delete()
	return
}

func UserLogin(logRequest models.UserLogin) (login map[string]interface{}, err error) {
	var user models.User
	err = userSource.Find("email", logRequest.Email).One(&user)
	if err != nil {
		err = errors.New("Incorrect Password or UserName")
		return
	}
	user.UnmarshalDB()

	err = bcrypt.CompareHashAndPassword(user.Hash, append([]byte(logRequest.Password), user.Salt...))
	if err != nil {
		err = errors.New("Incorrect Password or UserName")
		return
	}
	token, err := newToken(user.UserID)
	if err != nil {
		return
	}

	sToken, err := token.SignedString([]byte("pl8IKa8Wz5tu64JuV3ksSQ7YVyDDjet17jE5YXS37lIasCxjhYlHjYYGnNT9Gzs"))
	if err != nil {
		return
	}
	user.AddJTI(token.Claims["jti"].(string))

	user.MarshalDB()
	err = userSource.Find("user_id", user.UserID).Update(user)
	if err != nil {
		return
	}

	login = map[string]interface{}{
		"name":  user.FirstName,
		"email": user.Email,
		"token": sToken,
	}
	return
}

func UserLogout(userID uint, jti string) (err error) {
	var user models.User
	err = userSource.Find("user_id", userID).One(&user)
	if err != nil {
		err = errors.New("Incorrect Password or UserName")
		return
	}
	user.UnmarshalDB()

	user.RemoveJTI(jti)

	user.MarshalDB()
	err = userSource.Find("user_id", userID).Update(user)
	if err != nil {
		return
	}

	return
}

func UserLogoutAll(userID uint) (err error) {
	var user models.User
	err = userSource.Find("user_id", userID).One(&user)
	if err != nil {
		err = errors.New("Incorrect Password or UserName")
		return
	}
	user.UnmarshalDB()

	user.ClearAllJTI()

	user.MarshalDB()
	err = userSource.Find("user_id", userID).Update(user)
	if err != nil {
		return
	}

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
