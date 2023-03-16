package serviceApi

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"gorm.io/gorm"
)

func GetAllRoles(db *database.Database, dest *[]models.Role) *gorm.DB {
	return db.Preload("ApartmentPrefs").Preload("Apartments").Find(&dest)
}

func GetRole(db *database.Database, dest *models.Role, id interface{}) *gorm.DB {
	return db.Find(&dest, id)
}

func AddRole(db *database.Database, dest *models.Role) *gorm.DB {
	return db.Create(dest)
}

func EditRole(db *database.Database, dest *models.Role) *gorm.DB {
	return db.Save(dest)
}

func DeleteRole(db *database.Database, dest *models.Role) *gorm.DB {
	return db.Delete(dest)
}
