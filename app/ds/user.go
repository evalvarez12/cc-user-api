package ds

import (
	"crypto/rand"
	"errors"
	"github.com/evalvarez12/cc-user-api/app/models"
	"golang.org/x/crypto/bcrypt"
	"time"
	"upper.io/db.v2"
)

func GetSession(token string) (userID uint, jti string, err error) {
	claims, err := ValidateToken(token)
	if err != nil {
		return
	}

	userID = uint(claims["id"].(float64))
	var user models.User
	err = userSource.Find(db.Cond{"user_id": userID}).One(&user)
	if err != nil {
		return
	}

	user.UnmarshalDB()

	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		err = errors.New("Session expired")
		user.RemoveJTI(claims["jti"].(string))
		user.MarshalDB()
		_ = userSource.Find(db.Cond{"user_id": userID}).Update(user)
		return
	}

	if !user.ContainsJTI(claims["jti"].(string)) {
		err = errors.New("Non existant session")
		return
	}

	jti = claims["jti"].(string)
	return
}

func Add(user models.User) (userID uint, err error) {
	hashPassword(&user)

	user.MarshalDB()
	temp, err := userSource.Insert(user)
	if err != nil {
		return
	}
	userID = uint(temp.(int64))
	return
}

func Delete(userID uint) (err error) {
	err = userSource.Find(db.Cond{"user_id": userID}).Delete()
	return
}

func Login(logRequest models.UserLogin) (login map[string]interface{}, err error) {
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
		"name":    user.FirstName,
		"email":   user.Email,
		"token":   sToken,
		"user_id": user.UserID,
		"answers": user.Answers,
	}
	return
}

func Logout(userID uint, jti string) (err error) {
	var user models.User
	err = userSource.Find(db.Cond{"user_id": userID}).One(&user)
	if err != nil {
		return
	}
	user.UnmarshalDB()

	user.RemoveJTI(jti)

	user.MarshalDB()
	err = userSource.Find(db.Cond{"user_id": userID}).Update(user)
	return
}

func LogoutAll(userID uint) (err error) {
	var user models.User
	err = userSource.Find(db.Cond{"user_id": userID}).One(&user)
	if err != nil {
		return
	}
	user.UnmarshalDB()

	user.ClearAllJTI()

	user.MarshalDB()
	err = userSource.Find(db.Cond{"user_id": userID}).Update(user)
	return
}

func UpdateAnswers(userID uint, answers models.Answers) (err error) {
	var user models.User
	err = userSource.Find(db.Cond{"user_id": userID}).One(&user)
	if err != nil {
		return
	}
	// user.UnmarshalDB()
	user.Answers = answers.Answers
	// user.MarshalDB()
	err = userSource.Find(db.Cond{"user_id": userID}).Update(user)
	return
}

func Update(userID uint, userNew models.User) (err error) {
	var user models.User
	err = userSource.Find(db.Cond{"user_id": userID}).One(&user)
	if err != nil {
		return
	}
	user.Update(userNew)
	err = userSource.Find(db.Cond{"user_id": userID}).Update(user)
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
