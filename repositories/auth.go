package repositories

import (
	"BE-foodways/models"

	"gorm.io/gorm"
)

type AuthRepositories interface {
	Register(user models.User) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
