package repositories

import (
	"BE-foodways/models"

	"gorm.io/gorm"
)

type UserRepositories interface {
	FindUsers() ([]models.User, error)
  GetUser(ID int) (models.User, error)
}

type repositories struct {
	db *gorm.DB
}

func RepositoriesUser(db *gorm.DB) *repositories{
	return &repositories{db}
}

func (r *repositories) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}