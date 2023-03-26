package serviceApi

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"gorm.io/gorm"
)

func GetAllApartments(db *database.Database, dest *[]models.Apartment) *gorm.DB {
	return db.Find(&dest)
}

func GetApartment(db *database.Database, dest *models.Apartment, id string) *gorm.DB {
	return db.Find(&dest, id)
}

func GetApartmentsByUserID(db *database.Database, dest *[]models.Apartment, id string) *gorm.DB {
	return db.Find(&dest, "user_id = ?", id)
}

func AddApartment(db *database.Database, dest *models.Apartment) *gorm.DB {
	return db.Create(dest)
}

func EditApartment(db *database.Database, dest *models.Apartment) *gorm.DB {
	return db.Save(dest)
}

func DeleteApartment(db *database.Database, dest *models.Apartment) *gorm.DB {
	return db.Delete(dest)
}
