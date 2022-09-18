package jwt_handler

import (
	"day02/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	UserUUID string `json:"user_uuid"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte(config.App.JwtSecret)
var lifetime = 600

func GenerateToken(userId string) (token string, err error) {
	tokenLifetime := time.Duration(lifetime)

	// Create the claims
	claims := &MyCustomClaims{
		UserUUID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			Issuer:    "Authenticator",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(tokenLifetime * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(signedToken string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	return claims, nil
}
