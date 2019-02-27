package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	user := new(models.User)
	var err error
	return c.JSON(http.StatusCreated, user)
}
