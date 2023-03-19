package serviceApi

import (
	"fiber-boilerplate/app/models"
	"time"
)

func isDateBetween(date, start, end time.Time) bool {
	return date.After(start) && date.Before(end)
}

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
	if !isDateBetween(a.AvailableDate, ap.AvailableDate[0], ap.AvailableDate[1]) {
		return false
	}

	return true
}

func GetMatchingApartments(apartments []models.Apartment, apartmentPrefs []models.ApartmentPref) *map[uint][]models.Apartment {
	MatchingApartmentsPerPref := make(map[uint][]models.Apartment)
	for _, ap := range apartmentPrefs {
		var MatchingApartments []models.Apartment
		for _, a := range apartments {
			if IsApartmentMatchesPref(&a, &ap) {
				MatchingApartments = append(MatchingApartments, a)
			}
		}
		MatchingApartmentsPerPref[ap.ID] = MatchingApartments
	}
	return &MatchingApartmentsPerPref
}
