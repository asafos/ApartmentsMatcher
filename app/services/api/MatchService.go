package serviceApi

import (
	"fiber-boilerplate/app/models"
)

func dateRangesOverlap(dateSliceA, dateSliceB models.TimeSlice) bool {
	// check if either of the date ranges is invalid
	if dateSliceA[0].After(dateSliceA[1]) || dateSliceB[0].After(dateSliceB[1]) {
		return false
	}
	// check if the date ranges overlap
	return dateSliceA[0].Before(dateSliceB[1]) && dateSliceB[0].Before(dateSliceA[1])
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
	if !dateRangesOverlap(a.AvailableDates, ap.AvailableDates) {
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
