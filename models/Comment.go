package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Message string `gorm:"not null" validate:"required"`
	PhotoId uint   `json:"user_id"`
	UserId  uint   `json:"photo_id"`
}

type CommentUpdate struct {
	Message string `json:"message"`
}
