package serviceApi

import (
	"errors"
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
func FindUserByOAuthID(db *database.Database, oAuthID string) (*models.User, error) {
	User := new(models.User)
	if response := db.Where("oauth_id = ?", oAuthID).First(&User); response.Error != nil {
		return nil, response.Error
	}
	if User.ID == 0 {
		return User, errors.New("user not found")
	}
	return User, nil
}
