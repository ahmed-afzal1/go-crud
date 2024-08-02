package repositories

import (
	"github.com/ahmed-afzal1/go-crud/config"
	"github.com/ahmed-afzal1/go-crud/models"
)

func GetAllPost() ([]models.Post, error) {
	var posts []models.Post

	if err := config.DB.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func CreatePost(post models.Post) (models.Post, error) {
	if err := config.DB.Create(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

func FindPost(id uint) (models.Post, error) {
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		return post, err
	}

	return post, nil
}

func UpdatePost(post models.Post) (models.Post, error) {
	if err := config.DB.Save(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

func DeletePost(id uint) error {
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&models.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
