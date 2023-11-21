package authentication

import (
	"os"
	"github.com/golang-jwt/jwt/v4"

)

func ValidateJWT(s string) (claims *TokenClaims, msg string) {
	token, err := jwt.ParseWithClaims(s, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("COOKIE_SECRET_JWT_AUTH1")), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		msg = "token is invalid"
		return
	}

	return claims, msg

}

func ValidateRefreshJWT(s string) (claims *TokenClaims, msg string) {
	token, err := jwt.ParseWithClaims(s, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("COOKIE_SECRET_JWT_AUTH2")), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		msg = "token is invalid"
		return
	}

	return claims, msg

}
