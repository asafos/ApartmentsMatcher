package models

// ApartmentPref model
type ApartmentPref struct {
	CommonModelFields
	UserID         uint
	NumberOfRooms  Range         `json:"numberOfRooms" gorm:"type:text"`
	Price          Range         `json:"price" gorm:"type:text"`
	Balcony        bool          `json:"balcony"`
	Roof           bool          `json:"roof"`
	Parking        bool          `json:"parking"`
	Elevator       bool          `json:"elevator"`
	PetsAllowed    bool          `json:"petsAllowed"`
	Renovated      bool          `json:"renovated"`
	AvailableDates TimeSlice     `json:"availableDates" gorm:"type:text"`
	Location       LocationSlice `json:"location" gorm:"type:text"`
}
