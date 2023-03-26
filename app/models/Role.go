package models

// Role model
type Role struct {
	CommonModelFields
	Name        string `gorm:"unique;not null" json:"name" xml:"name" form:"name" query:"name"`
	Description string `gorm:"type:varchar(100);" json:"description"`
}
