package controllers

import (
	"day02/lib"
	"day02/lib/http_response"
	"day02/models"
	"day02/models/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsers() func(c echo.Context) error {
	return func(c echo.Context) error {
		data := models.GetUsers()

		return http_response.Json(c, http.StatusOK, "Users", data)
	}
}

func GetUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return http_response.Json(c, http.StatusBadRequest, err.Error(), nil)
		}

		data := models.GetUser(idInt)

		if data.ID == 0 {
			return http_response.Json(c, http.StatusNotFound, "User not found", nil)
		}

		return http_response.Json(c, http.StatusOK, "User", data)
	}
}

func CreateUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		var body entity.User
		c.Bind(&body)

		if err := lib.MinStringLength(body.Username, 5); err != nil {
			return http_response.Json(c, http.StatusBadRequest, err.Error(), nil)
		}

		data := models.CreateUser(&body)

		return http_response.Json(c, http.StatusOK, "User created", data)
	}
}

func UpdateUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return http_response.Json(c, http.StatusBadRequest, err.Error(), nil)
		}

		var body entity.User
		c.Bind(&body)

		data := models.UpdateUser(idInt, &body)

		if data.ID == 0 {
			return http_response.Json(c, http.StatusNotFound, "User not found", nil)
		}

		return http_response.Json(c, http.StatusOK, "User updated", data)
	}
}

func DeleteUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data := models.DeleteUser(idInt)

		if data.ID == 0 {
			return http_response.Json(c, http.StatusNotFound, "User not found", nil)
		}

		return http_response.Json(c, http.StatusOK, "User deleted", data)
	}
}
