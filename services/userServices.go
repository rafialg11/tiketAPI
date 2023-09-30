package services

import (
	"tiketAPI/configs"
	"tiketAPI/models"
)

// User Login
func Login(email string, password string) (interface{}, error) {
	var user models.User
	if err := configs.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// User Register
func Register(user models.User) (interface{}, error) {
	if err := configs.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
