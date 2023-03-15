package models

import (
	"gorm.io/gorm"
)

// ApartmentPref model
type ApartmentPref struct {
	gorm.Model
	UserID         uint
	NumberOfRooms  Range     `json:"numberOfRooms" gorm:"type:text"`
	Price          Range     `json:"price" gorm:"type:text"`
	Balcony        BoolPref  `json:"balcony"`
	Roof           BoolPref  `json:"roof"`
	Parking        BoolPref  `json:"parking"`
	Elevator       BoolPref  `json:"elevator"`
	AnimalsAllowed BoolPref  `json:"animalsAllowed"`
	Renovated      BoolPref  `json:"renovated"`
	AvailableDate  TimeArray `json:"availableDate" gorm:"type:text"`
	Area           AreaArray `json:"area" gorm:"type:text"`
}
