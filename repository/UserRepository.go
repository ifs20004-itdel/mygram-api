package repository

import (
	"mygramapi/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(u models.User) (models.User, error)
	Login(u models.UserRequest) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(id int) (models.User, error)
	Update(u models.User) (models.User, error)
	Delete(Id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(u models.User) (models.User, error) {
	err := r.db.Create(&u).Error
	return u, err
}

func (r *userRepository) Login(u models.UserRequest) (models.User, error) {
	userRes := models.User{}
	err := r.db.First(&userRes, "email = ?", u.Email).Take(&userRes).Error
	return userRes, err
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindById(id int) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, id).Error
	return user, err
}

func (r *userRepository) Update(u models.User) (models.User, error) {
	err := r.db.Save(&u).Error
	return u, err
}

func (r *userRepository) Delete(id int) error {
	err := r.db.Delete(&models.User{}, id).Error
	return err
}
