package ds

import (
	"crypto/rand"
	"errors"
	"github.com/arbolista-dev/cc-user-api/app/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
	"upper.io/db.v2"
	"github.com/lib/pq"
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
		err = errors.New(`{"session": "expired"}`)
		user.RemoveJTI(claims["jti"].(string))
		user.MarshalDB()
		_ = userSource.Find(db.Cond{"user_id": userID}).Update(user)
		return
	}

	if !user.ContainsJTI(claims["jti"].(string)) {
		err = errors.New(`{"session": "non-existent"}`)
		return
	}

	jti = claims["jti"].(string)
	return
}

func Add(user models.User) (login map[string]interface{}, err error) {
	hashPassword(&user)

	user.MarshalDB()
	temp, err := userSource.Insert(user)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr != nil {
			if pqErr.Code == "23505" {
				err = errors.New(`{"email": "non-unique"}`)
				return
			}
		}
		return
	}
	userID := uint(temp.(int64))

	token, err := newToken(userID)
	if err != nil {
		return
	}

	sToken, err := token.SignedString([]byte(os.Getenv("CC_JWTSIGN")))
	if err != nil {
		return
	}
	user.AddJTI(token.Claims["jti"].(string))

	user.MarshalDB()
	err = userSource.Find("user_id", userID).Update(user)
	if err != nil {
		return
	}

	login = map[string]interface{}{
		"name":    user.FirstName,
		"token":   sToken,
		"answers": user.Answers.String(),
	}
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
		err = errors.New(`{"email": "non-existent"}`)
		return
	}
	user.UnmarshalDB()

	err = bcrypt.CompareHashAndPassword(user.Hash, append([]byte(logRequest.Password), user.Salt...))
	if err != nil {
		err = errors.New(`{"password": "incorrect"}`)
		return
	}
	token, err := newToken(user.UserID)
	if err != nil {
		return
	}

	sToken, err := token.SignedString([]byte(os.Getenv("CC_JWTSIGN")))
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
		"token":   sToken,
		"answers": user.Answers.String(),
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


func UpdateTotalFootprint(userID uint, footprint models.TotalFootprint) (err error) {
	var user models.User
	err = userSource.Find(db.Cond{"user_id": userID}).One(&user)
	if err != nil {
		return
	}
	user.TotalFootprint = footprint.TotalFootprint
	err = userSource.Find(db.Cond{"user_id": userID}).Update(user)
	return
}

func SetLocation(userID uint, location models.Location) (err error) {
	var user models.User
	err = userSource.Find(db.Cond{"user_id": userID}).One(&user)
	if err != nil {
		return
	}
	user.City = location.City
	user.State = location.State
	user.County = location.County
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

func PassResetRequest(email string) (userID uint, token string, err error) {
	var user models.User
	err = userSource.Find(db.Cond{"email": email}).One(&user)
	if err != nil {
		return
	}
	token = hashReset(&user)
	userID = user.UserID
	err = userSource.Find(db.Cond{"user_id": user.UserID}).Update(user)
	return
}

func PassResetConfirm(userID uint, token, password string) (err error) {
	var user models.User
	err = userSource.Find(db.Cond{"user_id": userID}).One(&user)
	if err != nil {
		return
	}
	user.UnmarshalDB()
	if user.ResetExpiration.After(time.Now()) {
		user.ResetHash = []byte{}
		user.ResetExpiration = time.Time{}
		err = errors.New(`{"password-reset": "expired"}`)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.ResetHash, []byte(token))
	if err != nil {
		err = errors.New(`{"reset-token": "corrupt"}`)
		return err
	}

	user.Password = password
	hashPassword(&user)

	user.ClearAllJTI()
	user.ResetHash = []byte{}
	user.ResetExpiration = time.Time{}

	user.MarshalDB()
	err = userSource.Find(db.Cond{"user_id": user.UserID}).Update(user)
	if err != nil {
		return
	}
	return
}

func ListLeaders(limit int, offset int, category string, city string, state string) (leaders models.PaginatedLeaders, err error) {

	switch {
	case category=="food":
		if len(city) == 0 && len(state) == 0 {
			query = leadersFoodSource.Find().OrderBy("footprint")
		} else if len(city) > 0 && len(state) > 0 {
			query = leadersFoodSource.Find(db.Cond{"city": city},db.Cond{"state": state}).OrderBy("footprint")
		} else if len(city) > 0  && len(state) == 0 {
			query = leadersFoodSource.Find(db.Cond{"city": city}).OrderBy("footprint")
		} else if len(city) == 0 && len(state) > 0{
			query = leadersFoodSource.Find(db.Cond{"state": state}).OrderBy("footprint")
		}
		break
	case category=="housing":
		if len(city) == 0 && len(state) == 0 {
		  query = leadersHousingSource.Find().OrderBy("footprint")
		} else if len(city) > 0 && len(state) > 0 {
		  query = leadersHousingSource.Find(db.Cond{"city": city},db.Cond{"state": state}).OrderBy("footprint")
		} else if len(city) > 0  && len(state) == 0 {
		  query = leadersHousingSource.Find(db.Cond{"city": city}).OrderBy("footprint")
		} else if len(city) == 0 && len(state) > 0{
		  query = leadersHousingSource.Find(db.Cond{"state": state}).OrderBy("footprint")
		}
		break
	case category=="shopping":
		if len(city) == 0 && len(state) == 0 {
		  query = leadersShoppingSource.Find().OrderBy("footprint")
		} else if len(city) > 0 && len(state) > 0 {
		  query = leadersShoppingSource.Find(db.Cond{"city": city},db.Cond{"state": state}).OrderBy("footprint")
		} else if len(city) > 0  && len(state) == 0 {
		  query = leadersShoppingSource.Find(db.Cond{"city": city}).OrderBy("footprint")
		} else if len(city) == 0 && len(state) > 0{
		  query = leadersShoppingSource.Find(db.Cond{"state": state}).OrderBy("footprint")
		}
		break
	case category=="transport":
		if len(city) == 0 && len(state) == 0 {
		  query = leadersTransportSource.Find().OrderBy("footprint")
		} else if len(city) > 0 && len(state) > 0 {
		  query = leadersTransportSource.Find(db.Cond{"city": city},db.Cond{"state": state}).OrderBy("footprint")
		} else if len(city) > 0  && len(state) == 0 {
		  query = leadersTransportSource.Find(db.Cond{"city": city}).OrderBy("footprint")
		} else if len(city) == 0 && len(state) > 0{
		  query = leadersTransportSource.Find(db.Cond{"state": state}).OrderBy("footprint")
		}
		break
	default:
		if len(city) == 0 && len(state) == 0 {
		  query = leadersSource.Find().OrderBy("footprint")
		} else if len(city) > 0 && len(state) > 0 {
		  query = leadersSource.Find(db.Cond{"city": city},db.Cond{"state": state}).OrderBy("footprint")
		} else if len(city) > 0  && len(state) == 0 {
		  query = leadersSource.Find(db.Cond{"city": city}).OrderBy("footprint")
		} else if len(city) == 0 && len(state) > 0{
		  query = leadersSource.Find(db.Cond{"state": state}).OrderBy("footprint")
		}
		break
	}

	count, err := query.Count()
	if err != nil {
		return
	}
	leaders.TotalCount = count

	list := query.Limit(limit).Offset(offset)
	err = list.All(&leaders.List)
	if err != nil {
		return
	}

	return
}

func ListLocations() (locations []models.Location, err error) {

	q := leadersSource.Find().Select("city", "state", "county").Group("city", "state", "county")
	err = q.All(&locations)
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

func hashReset(user *models.User) (token string) {
	token = randString(10)
	var err error
	user.ResetHash, err = bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.ResetExpiration = time.Now().Add(time.Minute * 5)
	return
}
