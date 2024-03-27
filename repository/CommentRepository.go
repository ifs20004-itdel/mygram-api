package repository

import (
	"mygramapi/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(u *models.User, c models.Comment) (models.Comment, error)
	FindAll() ([]models.Comment, error)
	FindUserById(id int) (models.User, error)
	FindCommentById(ID int) (models.Comment, error)
	Update(u models.Comment) (models.Comment, error)
	Delete(Id int) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(u *models.User, c models.Comment) (models.Comment, error) {
	err := r.db.Model(u).Association("Comments").Append(&c)
	return c, err
}

func (r *commentRepository) FindAll() ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Find(&comments).Error
	return comments, err
}

func (r *commentRepository) FindUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, id).Error
	return user, err
}

func (r *commentRepository) FindCommentById(id int) (models.Comment, error) {
	var comment models.Comment
	err := r.db.Find(&comment, id).Error
	return comment, err
}

func (r *commentRepository) Update(p models.Comment) (models.Comment, error) {
	err := r.db.Save(&p).Error
	return p, err
}

func (r *commentRepository) Delete(id int) error {
	err := r.db.Delete(&models.Comment{}, id).Error
	return err
}
