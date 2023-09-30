package controllers

import (
	"net/http"
	"tiketAPI/models"
	"tiketAPI/services"

	"github.com/labstack/echo/v4"
)

// login
func Login(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	result, err := services.Login(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func Register(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	result, err := services.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
