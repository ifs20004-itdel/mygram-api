package models

type SocialMedia struct {
	Name           string `json:"name" gorm:"not null"`
	SocialMediaURL string `json:"social_media_url" gorm:"not null;type:text"`
	UserID         uint   `json:"user_id"`
}
