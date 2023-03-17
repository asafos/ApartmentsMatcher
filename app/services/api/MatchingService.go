package serviceApi

import (
	"fiber-boilerplate/app/models"
)

func IsApartmentMatchesPref(a *models.Apartment, ap *models.ApartmentPref) bool {
	if a.Price < ap.Price[0] || a.Price > ap.Price[1] {
		return false
	}
	isLocationCompatible := false
	for _, l := range ap.Location {
		if l == a.Location {
			isLocationCompatible = true
		}
	}

	if !isLocationCompatible {
		return false
	}
	isAvailableDateCompatible := false
	for _, l := range ap.AvailableDate {
		if l == a.AvailableDate {
			isLocationCompatible = true
		}
	}

	return isAvailableDateCompatible

	// if !isAvailableDateCompatible {
	// 	return false
	// }

	// return true
}

func GetMatchingApartments(apartments []models.Apartment, apartmentPrefs []models.ApartmentPref) *[]models.Apartment {
	var MatchingApartments []models.Apartment
	for _, ap := range apartmentPrefs {
		for _, a := range apartments {
			if IsApartmentMatchesPref(&a, &ap) {
				MatchingApartments = append(MatchingApartments, a)
			}
		}
	}
	return &MatchingApartments
}
