package models

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
	PetsAllowed    bool      `json:"petsAllowed"`
	Renovated      bool      `json:"renovated"`
	AvailableDates TimeSlice `json:"availableDates" gorm:"type:text"`
	Location       Location  `json:"location"`
}
