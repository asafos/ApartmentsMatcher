package api

import (
	"fiber-boilerplate/app/constants"
	utils "fiber-boilerplate/app/controllers/utils"
	"fiber-boilerplate/app/models"
	services "fiber-boilerplate/app/services/api"
	"fiber-boilerplate/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Return all apartmentPrefs as JSON
func GetAllApartmentPrefs(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var ApartmentPrefs []models.ApartmentPref
		if response := services.GetAllApartmentPrefs(db, &ApartmentPrefs); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving apartmentPrefs from the database: "+response.Error.Error())
		}
		err := ctx.JSON(ApartmentPrefs)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of apartmentPrefs: "+err.Error())
		}
		return err
	}
}

// Return apartmentPrefs as JSON
func GetUserApartmentPrefs(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		intID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		userID := ctx.Locals(constants.USER_LOCALS_KEY)
		if userID != uint(intID) {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized user")
		}
		var ApartmentPref []models.ApartmentPref
		if response := services.GetApartmentPrefsByUserID(db, &ApartmentPref, id); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving apartments from the database: "+response.Error.Error())
		}
		err = ctx.JSON(ApartmentPref)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of apartments: "+err.Error())
		}
		return err
	}
}

// Return a single apartmentPref as JSON
func GetApartmentPref(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ApartmentPref := new(models.ApartmentPref)
		id := ctx.Params("id")
		if response := services.GetApartmentPref(db, ApartmentPref, id); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "An error occurred when retrieving the apartmentPref: "+response.Error.Error())
		}
		if ApartmentPref.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Cannot return status not found: "+err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a role: "+err.Error())
			}
			return err
		}
		userID := ctx.Locals(constants.USER_LOCALS_KEY)
		if userID != ApartmentPref.UserID {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized user")
		}
		err := ctx.JSON(ApartmentPref)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a apartmentPref: "+err.Error())
		}
		return err
	}
}

// Add a single apartmentPref to the database
func AddApartmentPref(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ApartmentPref := new(models.ApartmentPref)
		if err := ctx.BodyParser(ApartmentPref); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "an error occurred when parsing the new apartmentPref")
		}
		errors := utils.ValidateStruct(*ApartmentPref)
		if errors != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(errors)
		}
		if response := services.AddApartmentPref(db, ApartmentPref); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "an error occurred when storing the new apartmentPref"+response.Error.Error())
		}
		err := ctx.JSON(ApartmentPref)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "error occurred when returning JSON of a apartmentPref")
		}
		return err
	}
}

// Edit a single apartmentPref
func EditApartmentPref(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditApartmentPref := new(models.ApartmentPref)
		ApartmentPref := new(models.ApartmentPref)
		if err := ctx.BodyParser(EditApartmentPref); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "An error occurred when parsing the edited apartmentPref: "+err.Error())
		}
		errors := utils.ValidateStruct(*EditApartmentPref)
		if errors != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(errors)
		}
		if response := services.GetApartmentPref(db, ApartmentPref, id); response.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, "An error occurred when retrieving the existing apartmentPref: "+response.Error.Error())
		}
		// ApartmentPref does not exist
		if ApartmentPref.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Cannot return status not found: "+err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a apartmentPref: "+err.Error())
			}
			return err
		}
		userID := ctx.Locals(constants.USER_LOCALS_KEY)
		if userID != ApartmentPref.UserID {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized user")
		}
		ApartmentPref.NumberOfRooms = EditApartmentPref.NumberOfRooms
		ApartmentPref.Price = EditApartmentPref.Price
		ApartmentPref.Balcony = EditApartmentPref.Balcony
		ApartmentPref.Roof = EditApartmentPref.Roof
		ApartmentPref.Parking = EditApartmentPref.Parking
		ApartmentPref.Elevator = EditApartmentPref.Elevator
		ApartmentPref.PetsAllowed = EditApartmentPref.PetsAllowed
		ApartmentPref.Renovated = EditApartmentPref.Renovated
		ApartmentPref.AvailableDates = EditApartmentPref.AvailableDates
		ApartmentPref.Location = EditApartmentPref.Location

		// Save apartmentPref
		services.EditApartmentPref(db, ApartmentPref)

		err := ctx.JSON(ApartmentPref)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a apartmentPref: "+err.Error())
		}
		return err
	}
}

// Delete a single apartmentPref
func DeleteApartmentPref(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var ApartmentPref models.ApartmentPref
		// services.Find(&ApartmentPref, id)
		if response := services.GetApartmentPref(db, &ApartmentPref, id); response.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, "An error occurred when finding the apartmentPref to be deleted"+response.Error.Error())
		}
		userID := ctx.Locals(constants.USER_LOCALS_KEY)
		if userID != ApartmentPref.UserID {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized user")
		}
		services.DeleteApartmentPref(db, &ApartmentPref)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a apartmentPref: "+err.Error())
		}
		return err
	}
}
