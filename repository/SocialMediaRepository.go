package repository

import (
	"mygramapi/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(u *models.User, c models.SocialMedia) (models.SocialMedia, error)
	FindAll() ([]models.SocialMedia, error)
	FindUserById(id int) (models.User, error)
	FindSocialMediaById(ID int) (models.SocialMedia, error)
	Update(u models.SocialMedia) (models.SocialMedia, error)
	Delete(Id int) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) Create(u *models.User, c models.SocialMedia) (models.SocialMedia, error) {
	err := r.db.Model(u).Association("SocialMedias").Append(&c)
	return c, err
}

func (r *socialMediaRepository) FindAll() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := r.db.Find(&socialMedias).Error
	return socialMedias, err
}

func (r *socialMediaRepository) FindUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, id).Error
	return user, err
}

func (r *socialMediaRepository) FindSocialMediaById(id int) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := r.db.Find(&socialMedia, id).Error
	return socialMedia, err
}

func (r *socialMediaRepository) Update(sm models.SocialMedia) (models.SocialMedia, error) {
	err := r.db.Save(&sm).Error
	return sm, err
}

func (r *socialMediaRepository) Delete(id int) error {
	err := r.db.Delete(&models.SocialMedia{}, id).Error
	return err
}
