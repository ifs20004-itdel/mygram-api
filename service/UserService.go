package service

import (
	"mygramapi/models"
	"mygramapi/repository"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserService interface {
	Create(user models.User) (models.UserResponse, error)
	Login(user models.UserRequest) (models.User, error)
	FindAll() ([]models.UserResponse, error)
	FindById(ID int) (models.User, error)
	Update(id int, newUser models.UserUpdate) (models.UserResponse, error)
	Delete(ID int) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{userRepository: repo}
}

func (s *userService) Create(user models.User) (models.UserResponse, error) {
	createUser, err := s.userRepository.Create(user)

	var userCreated = models.UserResponse{}
	userCreated.ID = createUser.ID
	userCreated.Age = createUser.Age
	userCreated.Email = createUser.Email
	userCreated.Username = createUser.Username

	return userCreated, err
}

func (s *userService) Login(u models.UserRequest) (models.User, error) {
	return s.userRepository.Login(u)
}

func (s *userService) FindAll() ([]models.UserResponse, error) {
	user, err := s.userRepository.FindAll()
	userResponses := make([]models.UserResponse, 0, len(user))
	for _, item := range user {
		userResponse := models.UserResponse{
			ID:       item.ID,
			Username: item.Username,
			Email:    item.Email,
			Age:      item.Age,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, err
}

func (s *userService) FindById(id int) (models.User, error) {
	user, err := s.userRepository.FindById(id)

	return user, err
}

func (s *userService) Update(id int, newUser models.UserUpdate) (models.UserResponse, error) {

	var userResponse models.UserResponse
	user, _ := s.userRepository.FindById(id)

	user.Email = newUser.Email
	user.Username = newUser.Username

	updateUser, err := s.userRepository.Update(user)

	userResponse.ID = updateUser.ID
	userResponse.Age = updateUser.Age
	userResponse.Email = updateUser.Email
	userResponse.Username = updateUser.Username

	return userResponse, err
}

func (s *userService) Delete(id int) error {
	return s.userRepository.Delete(id)
}
