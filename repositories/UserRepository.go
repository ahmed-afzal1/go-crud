package repositories

import (
	"github.com/ahmed-afzal1/go-crud/config"
	"github.com/ahmed-afzal1/go-crud/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func FindUser(id uint) (models.User, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DeleteUser(id uint) error {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
