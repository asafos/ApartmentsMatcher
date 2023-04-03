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

// Return all apartments as JSON
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

		matchingApartments := services.GenerateMatches(Apartments, ApartmentsPref)

		if err := appCache.Set(&cache.Item{
			Key:   "matches",
			Value: matchingApartments,
			// TTL:   time.Hour,
		}); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when persisting matches: "+err.Error())
		}

		err := ctx.JSON(matchingApartments)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of matches: "+err.Error())
		}
		return nil
	}
}
