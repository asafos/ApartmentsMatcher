package serviceApi

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"gorm.io/gorm"
)

func GetAllUsers(db *database.Database, dest *[]models.User) *gorm.DB {
	return db.Find(&dest)
}

func GetUser(db *database.Database, dest *models.User, id string) *gorm.DB {
	return db.Find(&dest, id)
}

func AddUser(db *database.Database, dest *models.User) *gorm.DB {
	return db.Create(dest)
}

func EditUser(db *database.Database, dest *models.User) *gorm.DB {
	return db.Save(dest)
}

func DeleteUser(db *database.Database, dest *models.User) *gorm.DB {
	return db.Delete(dest)
}
