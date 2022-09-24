package http_response

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	TimeStamps time.Time   `json:"timestamps"`
	Data       interface{} `json:"data"`
}

func Json(c echo.Context, code int, msg string, data interface{}) error {
	res := Response{
		StatusCode: code,
		Status:     http.StatusText(code),
		Message:    msg,
		TimeStamps: time.Now().UTC(),
		Data:       data,
	}

	return c.JSON(code, res)
}
