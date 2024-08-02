package services

import (
	"github.com/ahmed-afzal1/go-crud/models"
	"github.com/ahmed-afzal1/go-crud/repositories"
)

func GetAllPost() ([]models.Post, error) {
	return repositories.GetAllPost()
}

func CreatePost(post models.Post) (models.Post, error) {
	return repositories.CreatePost(post)
}

func FindPost(id uint) (models.Post, error) {
	return repositories.FindPost(id)
}

func UpdatePost(post models.Post) (models.Post, error) {
	return repositories.UpdatePost(post)
}

func DeletePost(id uint) error {
	return repositories.DeletePost(id)
}
