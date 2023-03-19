package models

import (
	"gorm.io/gorm"
)

// ApartmentPref model
type ApartmentPref struct {
	gorm.Model
	UserID         uint
	NumberOfRooms  Range         `json:"numberOfRooms" gorm:"type:text"`
	Price          Range         `json:"price" gorm:"type:text"`
	Balcony        bool          `json:"balcony"`
	Roof           bool          `json:"roof"`
	Parking        bool          `json:"parking"`
	Elevator       bool          `json:"elevator"`
	AnimalsAllowed bool          `json:"animalsAllowed"`
	Renovated      bool          `json:"renovated"`
	AvailableDate  TimeSlice     `json:"availableDate" gorm:"type:text"`
	Location       LocationSlice `json:"location" gorm:"type:text"`
}
