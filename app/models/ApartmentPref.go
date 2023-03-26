package models

// ApartmentPref model
type ApartmentPref struct {
	CommonModelFields
	UserID         uint
	NumberOfRooms  Range         `json:"numberOfRooms" gorm:"type:text"`
	Price          Range         `json:"price" gorm:"type:text"`
	Balcony        *bool         `json:"balcony" validate:"required"`
	Roof           *bool         `json:"roof" validate:"required"`
	Parking        *bool         `json:"parking" validate:"required"`
	Elevator       *bool         `json:"elevator" validate:"required"`
	PetsAllowed    *bool         `json:"petsAllowed" validate:"required"`
	Renovated      *bool         `json:"renovated" validate:"required"`
	AvailableDates TimeSlice     `json:"availableDates" gorm:"type:text"`
	Location       LocationSlice `json:"location" gorm:"type:text"`
}
