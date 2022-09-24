package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	token.Claims = jwt.MapClaims{
		"user_id":    userId,
		"exp":        time.Now().Add(time.Hour * 3).Unix(),
		"authorized": true,
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractTokenData(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(int))

	return userId
}

type TokenData struct {
	UserId       int
	IsAuthorized bool
}
