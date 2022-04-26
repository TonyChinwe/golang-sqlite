package authentication

import (
	"github.com/dgrijalva/jwt-go"
)

func IsAuthorized(token string) bool {
	claim := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false
		}
		return false
	}
	if !tkn.Valid {
		return false
	}

	return true
}
