package repository

import (
	"mygramapi/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(u *models.User, p models.Photo) (models.Photo, error)
	FindAll() ([]models.Photo, error)
	FindUserById(id int) (models.User, error)
	FindPhotoById(ID int) (models.Photo, error)
	Update(u models.Photo) (models.Photo, error)
	Delete(Id int) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Create(u *models.User, p models.Photo) (models.Photo, error) {
	err := r.db.Model(u).Association("Photos").Append(&p)
	return p, err
}

func (r *photoRepository) FindAll() ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.Find(&photos).Error
	return photos, err
}

func (r *photoRepository) FindUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, id).Error
	return user, err
}

func (r *photoRepository) FindPhotoById(id int) (models.Photo, error) {
	var photo models.Photo
	err := r.db.Find(&photo, id).Error
	return photo, err
}

func (r *photoRepository) Update(p models.Photo) (models.Photo, error) {
	err := r.db.Save(&p).Error
	return p, err
}

func (r *photoRepository) Delete(id int) error {
	err := r.db.Delete(&models.Photo{}, id).Error
	return err
}
