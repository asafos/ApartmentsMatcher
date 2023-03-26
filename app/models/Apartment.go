package models

import (
	"time"
)

// Apartment model
type Apartment struct {
	CommonModelFields
	UserID         uint      `json:"user_id"`
	NumberOfRooms  int       `json:"numberOfRooms"`
	Price          int       `json:"price"`
	Balcony        bool      `json:"balcony"`
	Roof           bool      `json:"roof"`
	Parking        bool      `json:"parking"`
	Elevator       bool      `json:"elevator"`
	AnimalsAllowed bool      `json:"animalsAllowed"`
	Renovated      bool      `json:"renovated"`
	AvailableDate  time.Time `json:"availableDate"`
	Location       Location  `json:"location"`
}
