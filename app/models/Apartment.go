package models

import (
	"time"

	"gorm.io/gorm"
)

// Apartment model
type Apartment struct {
	gorm.Model
	UserID         uint
	NumberOfRooms  int       `json:"numberOfRooms"`
	Price          int       `json:"price"`
	Balcony        bool      `json:"balcony"`
	Roof           bool      `json:"roof"`
	Parking        bool      `json:"parking"`
	Elevator       bool      `json:"elevator"`
	AnimalsAllowed bool      `json:"animalsAllowed"`
	Renovated      bool      `json:"renovated"`
	AvailableDate  time.Time `json:"availableDate"`
	Area           Area      `json:"area"`
}
