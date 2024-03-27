package service

import (
	"mygramapi/models"
	"mygramapi/repository"
)

type SocialMediasService interface {
	Create(u *models.User, sm models.SocialMedia) (models.SocialMedia, error)
	FindAll() ([]models.SocialMedia, error)
	FindUserById(id int) (models.User, error)
	FindSocialMediaById(ID int) (models.SocialMedia, error)
	Update(id int, newSocialMedia models.SocialMedia) (models.SocialMedia, error)
	Delete(ID int) error
}

type socialMediaService struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(repo repository.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{socialMediaRepository: repo}
}

func (c *socialMediaService) Create(user *models.User, sm models.SocialMedia) (models.SocialMedia, error) {
	createSM, err := c.socialMediaRepository.Create(user, sm)
	return createSM, err
}

func (c *socialMediaService) FindAll() ([]models.SocialMedia, error) {
	socialMedias, err := c.socialMediaRepository.FindAll()

	return socialMedias, err
}

func (c *socialMediaService) FindUserById(id int) (models.User, error) {
	user, err := c.socialMediaRepository.FindUserById(id)
	return user, err
}

func (c *socialMediaService) FindSocialMediaById(id int) (models.SocialMedia, error) {
	socialMedia, err := c.socialMediaRepository.FindSocialMediaById(id)
	return socialMedia, err
}

func (c *socialMediaService) Update(id int, newSocialMedia models.SocialMedia) (models.SocialMedia, error) {

	m, _ := c.socialMediaRepository.FindSocialMediaById(id)

	m.Name = newSocialMedia.Name
	m.SocialMediaURL = newSocialMedia.SocialMediaURL

	updateSM, err := c.socialMediaRepository.Update(m)
	return updateSM, err
}

func (c *socialMediaService) Delete(id int) error {
	return c.socialMediaRepository.Delete(id)
}
