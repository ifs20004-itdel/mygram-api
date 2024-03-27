package service

import (
	"mygramapi/models"
	"mygramapi/repository"
)

type PhotoService interface {
	Create(u *models.User, p models.Photo) (models.Photo, error)
	FindAll() ([]models.Photo, error)
	FindUserById(id int) (models.User, error)
	FindPhotoById(ID int) (models.Photo, error)
	Update(id int, photo models.PhotoUpdate) (models.Photo, error)
	Delete(ID int) error
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(repo repository.PhotoRepository) *photoService {
	return &photoService{photoRepository: repo}
}

func (s *photoService) Create(user *models.User, photo models.Photo) (models.Photo, error) {
	createPhoto, err := s.photoRepository.Create(user, photo)
	return createPhoto, err
}

func (s *photoService) FindAll() ([]models.Photo, error) {
	photos, err := s.photoRepository.FindAll()

	return photos, err
}

func (s *photoService) FindUserById(id int) (models.User, error) {
	user, err := s.photoRepository.FindUserById(id)
	return user, err
}

func (s *photoService) FindPhotoById(id int) (models.Photo, error) {
	photo, err := s.photoRepository.FindPhotoById(id)
	return photo, err
}

func (s *photoService) Update(id int, newPhoto models.PhotoUpdate) (models.Photo, error) {

	photo, _ := s.photoRepository.FindPhotoById(id)

	photo.Title = newPhoto.Title
	photo.Caption = newPhoto.Caption
	photo.PhotoUrl = newPhoto.PhotoUrl

	updatePhoto, err := s.photoRepository.Update(photo)
	return updatePhoto, err
}

func (s *photoService) Delete(id int) error {
	return s.photoRepository.Delete(id)
}
