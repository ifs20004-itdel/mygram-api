package models

import "github.com/jinzhu/gorm"

type Photo struct {
	gorm.Model
	Title    string    `json:"title" validate:"required"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `json:"photo_url" validate:"required"`
	UserID   uint      `json:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PhotoUpdate struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}
