package services

import (
	"github.com/ahmed-afzal1/go-crud/models"
	"github.com/ahmed-afzal1/go-crud/repositories"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func CreateUser(user models.User) (models.User, error) {
	return repositories.CreateUser(user)
}

func FindUser(id uint) (models.User, error) {
	return repositories.FindUser(id)
}

func UpdateUser(user models.User) (models.User, error) {
	return repositories.UpdateUser(user)
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}
