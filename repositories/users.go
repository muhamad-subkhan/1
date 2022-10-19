package repositories

import (
	"BE-foodways/models"

	"gorm.io/gorm"
)

type UserRepositories interface {
	FindUsers() ([]models.User, error)
  GetUser(ID int) (models.User, error)
  CreateUser(user models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoriesUser(db *gorm.DB) *repository{
	return &repository{db}

}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO users(name, email, phone, location, image, role) VALUES (?,?,?,?,?,?)", user.Name, user.Email, user.Phone, user.Location, user.Image, user.Role).Error

	return user, err 
}