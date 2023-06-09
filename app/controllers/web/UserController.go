package web

import (
	"errors"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"
)

// Return a single user as JSON
func FindUserByUsername(db *database.Database, username string) (*models.User, error) {
	User := new(models.User)
	if response := db.Where("name = ?", username).First(&User); response.Error != nil {
		return nil, response.Error
	}
	if User.ID == 0 {
		return User, errors.New("user not found")
	}
	return User, nil
}

// Return a single user as JSON
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

// Return a single user as JSON
func FindUserByID(db *database.Database, id uint) (*models.User, error) {
	User := new(models.User)
	if response := db.Where("id = ?", id).First(&User); response.Error != nil {
		return nil, response.Error
	}
	if User.ID == 0 {
		return User, errors.New("user not found")
	}
	return User, nil
}
