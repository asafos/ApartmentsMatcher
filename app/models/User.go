package models

// User model
type User struct {
	CommonModelFields
	OAuthID string `gorm:"column:oauth_id" json:"oauth_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	RoleID  uint   `gorm:"column:role_id" json:"role_id"`
	Role    Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"role"`
	// Apartments     ApartmentSlice     `gorm:"type:text"`
	// ApartmentPrefs ApartmentPrefSlice `gorm:"type:text"`
}
