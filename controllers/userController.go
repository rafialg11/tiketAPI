package controllers

import (
	"net/http"
	"tiketAPI/helpers"
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

	// Generate token
	token, err := helpers.GenerateToken(result.(models.User).ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Return token
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
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
