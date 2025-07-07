// services/user_service.go

package services

import (
	"blog/database"
	"blog/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func LoginUser(username, password string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
