package models

// Apartment model
type Apartment struct {
	CommonModelFields
	UserID         uint      `json:"user_id" validate:"required,number,min=1"`
	NumberOfRooms  int       `json:"numberOfRooms" validate:"required,number,min=1"`
	Price          int       `json:"price" validate:"required,number,min=1"`
	Balcony        *bool     `json:"balcony" validate:"required"`
	Roof           *bool     `json:"roof" validate:"required"`
	Parking        *bool     `json:"parking" validate:"required"`
	Elevator       *bool     `json:"elevator" validate:"required"`
	PetsAllowed    *bool     `json:"petsAllowed" validate:"required"`
	Renovated      *bool     `json:"renovated" validate:"required"`
	AvailableDates TimeSlice `json:"availableDates" gorm:"type:text" validate:"required"`
	Location       Location  `json:"location" validate:"required"`
}
