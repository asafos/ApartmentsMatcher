package serviceApi

import (
	"fiber-boilerplate/app/models"
	"sort"
	"strconv"
	"strings"
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

type MatchingResults struct {
	Couples    [][]uint `json:"couples"`
	Threesomes [][]uint `json:"threesomes"`
}

func GenerateMatches(apartments []models.Apartment, apartmentPrefs []models.ApartmentPref) *MatchingResults {
	MatchingApartmentsPerPref := make(map[uint][]uint)
	ApartmentPrefIDByApartmentID := make(map[uint]uint)
	UserIDByApartmentPrefID := make(map[uint]uint)
	for _, ap := range apartmentPrefs {
		var MatchingApartmentsIDs []uint
		for _, a := range apartments {
			if IsApartmentMatchesPref(&a, &ap) {
				MatchingApartmentsIDs = append(MatchingApartmentsIDs, a.ID)
			}
			if ap.UserID == a.UserID {
				ApartmentPrefIDByApartmentID[a.ID] = ap.ID
			}
		}
		MatchingApartmentsPerPref[ap.ID] = MatchingApartmentsIDs
		UserIDByApartmentPrefID[ap.ID] = ap.UserID
	}

	var MatchingUsersFirstRelation [][]uint
	var MatchingUsersSecondRelation [][]uint
	MatchesString := make(map[string]bool)

	for apID1, matches1 := range MatchingApartmentsPerPref {
		for _, aID2 := range matches1 {
			apID2 := ApartmentPrefIDByApartmentID[aID2]
			matches2 := MatchingApartmentsPerPref[apID2]
			if apID1 == apID2 {
				break
			}
			for _, aID3 := range matches2 {
				apID3 := ApartmentPrefIDByApartmentID[aID3]
				if apID2 == apID3 {
					break
				}
				matches3 := MatchingApartmentsPerPref[apID3]
				// check if threesome match
				for _, aID4 := range matches3 {
					apID4 := ApartmentPrefIDByApartmentID[aID4]
					if apID3 == apID4 {
						break
					}
					if apID1 == apID4 {
						MatchingUsers := []uint{UserIDByApartmentPrefID[apID1], UserIDByApartmentPrefID[apID2], UserIDByApartmentPrefID[apID3]}
						MatchString := FormatMatchString(MatchingUsers)
						if _, ok := MatchesString[MatchString]; !ok {
							MatchesString[MatchString] = true
							MatchingUsersSecondRelation = append(MatchingUsersSecondRelation, MatchingUsers)
						}
					}
				}
				// check if couple match
				if apID1 == apID3 {
					MatchingUsers := []uint{UserIDByApartmentPrefID[apID1], UserIDByApartmentPrefID[apID2]}
					MatchString := FormatMatchString(MatchingUsers)
					if _, ok := MatchesString[MatchString]; !ok {
						MatchesString[MatchString] = true
						MatchingUsersFirstRelation = append(MatchingUsersFirstRelation, MatchingUsers)
					}
				}
			}
		}
	}

	return &MatchingResults{
		Couples:    MatchingUsersFirstRelation,
		Threesomes: MatchingUsersSecondRelation,
	}
}

func FormatMatchString(slice []uint) string {
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	valuesText := []string{}
	for _, val := range slice {
		text := strconv.FormatUint(uint64(val), 10)
		valuesText = append(valuesText, text)
	}
	return strings.Join(valuesText[:], ",")
}
