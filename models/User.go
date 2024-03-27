package models

import (
	"errors"
	"mygramapi/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string        `gorm:"not null;uniqueIndex;size:255" json:"username" form:"username" valid:"required~Your username is required"`
	Email        string        `gorm:"not null;uniqueIndex;size:255" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password     string        `gorm:"not null;" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          int           `gorm:"not null;" json:"age" valid:"required"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserUpdate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserRequest struct {
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errResp := govalidator.ValidateStruct(u)

	if errResp != nil {
		err = errResp
		return
	}

	if u.Age <= 8 {
		err = errors.New("must be at least 8 years old")
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
