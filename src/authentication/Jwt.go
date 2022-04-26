package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var SigningKey = []byte("secret_key")

func GenerateJwt(credential Credential) (string, error) {
	expireDate := time.Now().Add(time.Hour * 24)
	claim := &Claims{
		Username: credential.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireDate.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(SigningKey)

	return tokenString, err

}
