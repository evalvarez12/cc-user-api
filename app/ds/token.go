package ds

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "crypto/rand"
    "github.com/evalvarez12/cc-user-api/app/models"
    "fmt"
)

var (
	randomStringCharset = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
)

// String() returns a random string of the given size.
func randString(size uint) string {
	buf := make([]byte, size)
	rand.Read(buf)

	lr := uint(len(randomStringCharset))

	for i := uint(0); i < size; i++ {
		buf[i] = randomStringCharset[buf[i]%byte(lr)]
	}
	return string(buf)
}

func newToken(userID uint) (token *jwt.Token, err error) {
    token = jwt.New(jwt.SigningMethodHS256)
    token.Claims["id"] = userID
    token.Claims["iat"] = time.Now().Unix()
    token.Claims["exp"] = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
    token.Claims["jti"] = randString(5)
    return
}


func validateToken(sToken string, user models.User) (pass bool, err error) {
    token, err := jwt.Parse(sToken, func(token *jwt.Token) (interface{}, error) {
    // Don't forget to validate the alg is what you expect:
    if _, ok := token.Method.(jwt.SigningMethodHS256); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }
    return myLookupKey("ccsignature"), nil
    })

    if err == nil && token.Valid {
        pass = true
    }
    return
}
