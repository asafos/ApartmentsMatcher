package api

import (
	"fiber-boilerplate/app/models"
	services "fiber-boilerplate/app/services/api"
	"fiber-boilerplate/database"
	"strconv"

	"github.com/go-redis/cache/v8"
	"github.com/gofiber/fiber/v2"
)

// Return all apartments as JSON
func GetUserMatchingApartments(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Apartments []models.Apartment
		var ApartmentsPref []models.ApartmentPref
		id := ctx.Params("id")
		uint64Id, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error occurred while parsing user id")
		}
		uintId := uint(uint64Id)
		if response := services.GetAllApartments(db, &Apartments); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving apartments from the database: "+response.Error.Error())
		}
		if response := services.GetAllApartmentsPrefByUserId(db, &ApartmentsPref, uintId); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving apartmentPrefs from the database: "+response.Error.Error())
		}

		matchingApartments := services.GetMatchingApartments(Apartments, ApartmentsPref)

		err = ctx.JSON(matchingApartments)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of apartments: "+err.Error())
		}
		return nil
	}
}

// Generate matches
func GenerateMatches(db *database.Database, appCache *cache.Cache) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Apartments []models.Apartment
		var ApartmentsPref []models.ApartmentPref
		if response := services.GetAllApartments(db, &Apartments); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving apartments from the database: "+response.Error.Error())
		}
		if response := services.GetAllApartmentPrefs(db, &ApartmentsPref); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving apartmentPrefs from the database: "+response.Error.Error())
		}

		matchingApartmentsPerPref, err := services.GenerateMatchingApartmentsPerPref(Apartments, ApartmentsPref, appCache)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when persisting matches: "+err.Error())
		}

		err = ctx.JSON(matchingApartmentsPerPref)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of matches: "+err.Error())
		}
		return nil
	}
}

// Return all matches
// func GetMatches(appCache *cache.Cache) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		var matchingApartments services.MatchingResults
// 		if err := appCache.Get(context.Background(), MATCHES_PREFIX+ALL_MATCHES, &matchingApartments); err != nil {
// 			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when fetching matches: "+err.Error())
// 		}

// 		err := ctx.JSON(matchingApartments)
// 		if err != nil {
// 			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of matches: "+err.Error())
// 		}
// 		return nil
// 	}
// }

func GetUserMatches(db *database.Database, appCache *cache.Cache) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		uint64Id, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error occurred while parsing user id")
		}
		var ApartmentsPref []models.ApartmentPref
		uintId := uint(uint64Id)
		if response := services.GetAllApartmentsPrefByUserId(db, &ApartmentsPref, uintId); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving apartmentPrefs from the database: "+response.Error.Error())
		}

		matchingApartmentsPerPref, err := services.GetMatchingApartmentsByPrefs(ApartmentsPref, appCache)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving matchingApartmentsPerPref: "+err.Error())
		}
		err = ctx.JSON(matchingApartmentsPerPref)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of matches: "+err.Error())
		}
		return nil
	}
}
