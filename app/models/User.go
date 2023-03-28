package models

// User model
type User struct {
	CommonModelFields
	OAuthID string `gorm:"column:oauth_id" json:"oauth_id" validate:"required"`
	Name    string `json:"name"`
	Email   string `json:"email" validate:"required,email,min=6,max=64"`
	RoleID  uint   `gorm:"column:role_id" json:"role_id" validate:"required"`
	// Apartments     ApartmentSlice     `gorm:"type:text"`
	// ApartmentPrefs ApartmentPrefSlice `gorm:"type:text"`
}
