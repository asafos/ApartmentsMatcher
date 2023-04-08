package serviceApi

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"gorm.io/gorm"
)

func GetAllApartmentPrefs(db *database.Database, dest *[]models.ApartmentPref) *gorm.DB {
	return db.Find(&dest)
}

func GetAllApartmentsPrefByUserId(db *database.Database, dest *[]models.ApartmentPref, id uint) *gorm.DB {
	return db.Where(&models.ApartmentPref{UserID: id}).Find(&dest)
}

func GetApartmentPref(db *database.Database, dest *models.ApartmentPref, id string) *gorm.DB {
	return db.Find(&dest, id)
}

func GetApartmentPrefsByUserID(db *database.Database, dest *[]models.ApartmentPref, id string) *gorm.DB {
	return db.Find(&dest, "user_id = ?", id)
}

func AddApartmentPref(db *database.Database, dest *models.ApartmentPref) *gorm.DB {
	return db.Create(dest)
}

func EditApartmentPref(db *database.Database, dest *models.ApartmentPref) *gorm.DB {
	return db.Save(dest)
}

func DeleteApartmentPref(db *database.Database, dest *models.ApartmentPref) *gorm.DB {
	return db.Delete(dest)
}
