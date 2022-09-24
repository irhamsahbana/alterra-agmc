package controllers

import (
	"day02/lib/dto"
	"day02/lib/http_response"
	"day02/lib/jwt_handler"
	"day02/models"
	"day02/models/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login() func(c echo.Context) error {
	return func(c echo.Context) error {
		var body dto.Login
		c.Bind(&body)

		user := models.GetUserByUsername(body.Username)
		if user == nil {
			return http_response.Json(c, http.StatusNotFound, "User not found", nil)
		}

		valid := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if valid != nil {
			return http_response.Json(c, http.StatusUnauthorized, "Wrong password", nil)
		}
		// convert to string
		userId := strconv.Itoa(int(user.ID))
		token, err := jwt_handler.GenerateToken(userId)
		if err != nil {
			return http_response.Json(c, http.StatusInternalServerError, "Internal server error", nil)
		}

		data := map[string]string{
			"token": token,
		}

		return http_response.Json(c, http.StatusOK, "Login success", data)
	}
}

func Register() func(c echo.Context) error {
	return func(c echo.Context) error {
		var body dto.Register
		c.Bind(&body)

		user := models.GetUserByUsername(body.Username)
		if user != nil {
			if user.Username == body.Username {
				return http_response.Json(c, http.StatusConflict, "Username already exist", nil)
			}
		}

		user = &entity.User{
			Name:      body.Name,
			Username:  body.Username,
			Password:  body.Password,
			IsMarried: body.IsMarried,
			Phone:     body.Phone,
		}

		registeredUser := models.CreateUser(user)
		if registeredUser == nil {
			return http_response.Json(c, http.StatusInternalServerError, "Internal server error", nil)
		}

		//create token
		userId := strconv.Itoa(int(registeredUser.ID))
		token, err := jwt_handler.GenerateToken(userId)
		if err != nil {
			return http_response.Json(c, http.StatusInternalServerError, "Internal server error", nil)
		}

		data := map[string]interface{}{
			"user":  registeredUser,
			"token": token,
		}

		return http_response.Json(c, http.StatusOK, "Register success", data)
	}
}
