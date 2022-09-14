package controllers

import (
	"day02/lib"
	"day02/models"
	"day02/models/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsers() func(c echo.Context) error {
	return func(c echo.Context) error {
		data := models.GetUsers()

		return c.JSON(http.StatusOK, data)
	}
}

func GetUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		data := models.GetUser(idInt)

		if data.ID == 0 {
			return c.JSON(http.StatusNotFound, "User not found")
		}

		return c.JSON(http.StatusOK, data)
	}
}

func CreateUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		var body entity.User
		c.Bind(&body)

		if err := lib.MinStringLength(body.Username, 5); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data := models.CreateUser(&body)

		return c.JSON(http.StatusOK, data)
	}
}

func UpdateUser() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var body entity.User
		c.Bind(&body)

		data := models.UpdateUser(idInt, &body)

		return c.JSON(http.StatusOK, data)
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
			return c.JSON(http.StatusNotFound, "User not found")
		}

		return c.JSON(http.StatusOK, data)
	}
}
