package controllers

import (
	"day02/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllBooks() func(c echo.Context) error {
	return func(c echo.Context) error {
		data := models.GetAllBooks()

		return c.JSON(http.StatusOK, data)
	}
}

func GetBook() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		data := models.GetBook(idInt)

		return c.JSON(http.StatusOK, data)
	}
}

func CreateBook() func(c echo.Context) error {
	return func(c echo.Context) error {
		var body models.Book
		c.Bind(&body)

		data := models.CreateBook(&body)

		return c.JSON(http.StatusOK, data)
	}
}

func UpdateBook() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var body models.Book
		c.Bind(&body)

		data := models.UpdateBook(idInt, &body)

		return c.JSON(http.StatusOK, data)
	}
}

func DeleteBook() func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		data := models.DeleteBook(idInt)

		return c.JSON(http.StatusOK, data)
	}
}
