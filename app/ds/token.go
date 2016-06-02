package ds

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "crypto/rand"
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

func newToken(userID uint) (sToken string, err error) {
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims["id"] = userID
    token.Claims["iat"] = time.Now().Unix()
    token.Claims["exp"] = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
    token.Claims["jti"] = randString(5)
    sToken, err = token.SignedString([]byte("ccsignature"))
    return
}


func validateToken(sToken string, user models.User) (pass bool, err error) {
    token, err := jwt.Parse(sToken, "ccsignature")
    if err != nil {
        return
    }
    if user.ContainsJTI(token.Claims["jti"]) {
        pass = true
    } else {
        err = jwt.ErrNoTokenInRequest
    }
    return
}
