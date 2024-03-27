package service

import (
	"mygramapi/models"
	"mygramapi/repository"
)

type CommentService interface {
	Create(u *models.User, p models.Comment) (models.Comment, error)
	FindAll() ([]models.Comment, error)
	FindUserById(id int) (models.User, error)
	FindCommentById(ID int) (models.Comment, error)
	Update(id int, comment models.CommentUpdate) (models.Comment, error)
	Delete(ID int) error
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *commentService {
	return &commentService{commentRepository: repo}
}

func (c *commentService) Create(user *models.User, comment models.Comment) (models.Comment, error) {
	createComment, err := c.commentRepository.Create(user, comment)
	return createComment, err
}

func (c *commentService) FindAll() ([]models.Comment, error) {
	Comments, err := c.commentRepository.FindAll()

	return Comments, err
}

func (c *commentService) FindUserById(id int) (models.User, error) {
	user, err := c.commentRepository.FindUserById(id)
	return user, err
}

func (c *commentService) FindCommentById(id int) (models.Comment, error) {
	Comment, err := c.commentRepository.FindCommentById(id)
	return Comment, err
}

func (c *commentService) Update(id int, newComment models.CommentUpdate) (models.Comment, error) {

	comment, _ := c.commentRepository.FindCommentById(id)

	comment.Message = newComment.Message

	updateComment, err := c.commentRepository.Update(comment)
	return updateComment, err
}

func (c *commentService) Delete(id int) error {
	return c.commentRepository.Delete(id)
}
