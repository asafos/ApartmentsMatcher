package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	OAuthID        string `gorm:"column:oauth_id" json:"oauth_id"`
	Name           string `json:"name"`
	Email          string
	RoleID         uint               `gorm:"column:role_id" json:"role_id"`
	Role           Role               `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Apartments     ApartmentSlice     `gorm:"type:text"`
	ApartmentPrefs ApartmentPrefSlice `gorm:"type:text"`
}
