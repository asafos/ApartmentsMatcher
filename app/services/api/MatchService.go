package serviceApi

import (
	"context"
	"fiber-boilerplate/app/models"
	"sort"
	"strconv"
	"strings"

	"github.com/go-redis/cache/v8"
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

const (
	MATCHES_PREFIX              = "MATCHES:"
	ALL_MATCHES                 = "ALL"
	APARTMENT_PREF_MATCHES      = "APARTMENT_PREF:"
	APARTMENT_PREF_BY_APARTMENT = "APARTMENT_PREF_BY_APARTMENT"
	USER_BY_APARTMENT_PREF      = "APARTMENT_PREF_BY_APARTMENT"
)

func GenerateApartmentPrefMatchesCacheKey(apID uint) string {
	return MATCHES_PREFIX + APARTMENT_PREF_MATCHES + strconv.FormatUint(uint64(apID), 10)
}

func GenerateMatchingApartmentsPerPref(apartments []models.Apartment, apartmentPrefs []models.ApartmentPref, appCache *cache.Cache) (matches *map[uint][]uint, err error) {
	MatchingApartmentsPerPref := make(map[uint][]uint)
	ApartmentPrefIDByApartmentID := make(map[uint]uint)
	UserIDByApartmentPrefID := make(map[uint]uint)
	for _, ap := range apartmentPrefs {
		var MatchingApartmentsIDs []uint
		for _, a := range apartments {
			if ap.UserID == a.UserID {
				ApartmentPrefIDByApartmentID[a.ID] = ap.ID
			} else if IsApartmentMatchesPref(&a, &ap) {
				MatchingApartmentsIDs = append(MatchingApartmentsIDs, a.ID)
			}
		}
		MatchingApartmentsPerPref[ap.ID] = MatchingApartmentsIDs
		if err := appCache.Set(&cache.Item{
			Key:   GenerateApartmentPrefMatchesCacheKey(ap.ID),
			Value: MatchingApartmentsIDs,
		}); err != nil {
			return nil, err
		}
		UserIDByApartmentPrefID[ap.ID] = ap.UserID
	}

	if err := appCache.Set(&cache.Item{
		Key:   MATCHES_PREFIX + APARTMENT_PREF_BY_APARTMENT,
		Value: ApartmentPrefIDByApartmentID,
	}); err != nil {
		return nil, err
	}

	if err := appCache.Set(&cache.Item{
		Key:   MATCHES_PREFIX + USER_BY_APARTMENT_PREF,
		Value: UserIDByApartmentPrefID,
	}); err != nil {
		return nil, err
	}

	return &MatchingApartmentsPerPref, nil
}

type MatchItem struct {
	UserID              uint
	ApartmentID         uint
	ApartmentPrefID     uint
	MatchingApartmentID uint
}

type Match []MatchItem

func GetMatchingApartmentsByPref(ap models.ApartmentPref, appCache *cache.Cache) (matches []Match, err error) {
	var matches1 []uint
	var apartmentPrefIDByApartmentID map[uint]uint
	var userIDByApartmentPrefID map[uint]uint
	if err := appCache.Get(context.Background(), GenerateApartmentPrefMatchesCacheKey(ap.ID), &matches1); err != nil {
		return nil, err
	}
	if err := appCache.Get(context.Background(), MATCHES_PREFIX+APARTMENT_PREF_BY_APARTMENT, &apartmentPrefIDByApartmentID); err != nil {
		return nil, err
	}
	if err := appCache.Get(context.Background(), MATCHES_PREFIX+USER_BY_APARTMENT_PREF, &userIDByApartmentPrefID); err != nil {
		return nil, err
	}

	apID1 := ap.ID
	var aID1 uint
	for key, item := range apartmentPrefIDByApartmentID {
		if apID1 == item {
			aID1 = key
		}
	}
	matchesResult := []Match{}

	for _, aID2 := range matches1 {
		apID2 := apartmentPrefIDByApartmentID[aID2]
		var matches2 []uint
		if err := appCache.Get(context.Background(), GenerateApartmentPrefMatchesCacheKey(apID2), &matches2); err != nil {
			return nil, err
		}
		for _, aID3 := range matches2 {
			apID3 := apartmentPrefIDByApartmentID[aID3]
			var matches3 []uint
			if err := appCache.Get(context.Background(), GenerateApartmentPrefMatchesCacheKey(apID3), &matches3); err != nil {
				return nil, err
			}
			// check if threesome match
			for _, aID4 := range matches3 {
				apID4 := apartmentPrefIDByApartmentID[aID4]

				if apID1 == apID4 {
					matchesResult = append(matchesResult, Match{
						MatchItem{UserID: userIDByApartmentPrefID[apID1], ApartmentID: aID1, ApartmentPrefID: apID1, MatchingApartmentID: apID2},
						MatchItem{UserID: userIDByApartmentPrefID[apID2], ApartmentID: aID2, ApartmentPrefID: apID2, MatchingApartmentID: apID3},
						MatchItem{UserID: userIDByApartmentPrefID[apID3], ApartmentID: aID3, ApartmentPrefID: apID3, MatchingApartmentID: apID1},
					})
				}
			}
			// check if couple match
			if apID1 == apID3 {
				matchesResult = append(matchesResult, Match{
					MatchItem{UserID: userIDByApartmentPrefID[apID1], ApartmentID: aID1, ApartmentPrefID: apID1, MatchingApartmentID: apID2},
					MatchItem{UserID: userIDByApartmentPrefID[apID2], ApartmentID: aID2, ApartmentPrefID: apID2, MatchingApartmentID: apID1},
				})
			}
		}
	}
	return matchesResult, nil
}

func GetMatchingApartmentsByPrefs(apartmentPrefs []models.ApartmentPref, appCache *cache.Cache) (matches *map[uint][]Match, err error) {
	MatchingApartmentsByPrefs := make(map[uint][]Match)

	for _, ap := range apartmentPrefs {
		matchingApartments, err := GetMatchingApartmentsByPref(ap, appCache)
		if err != nil {
			return nil, err
		}
		MatchingApartmentsByPrefs[ap.ID] = matchingApartments
	}

	return &MatchingApartmentsByPrefs, nil
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
