package api

import (
	utils "fiber-boilerplate/app/controllers/utils"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

// Return all apartmentPrefs as JSON
func GetAllApartmentPrefs(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var ApartmentPrefs []models.ApartmentPref
		if response := db.Preload("ApartmentPrefPrefs").Preload("ApartmentPrefs").Find(&ApartmentPrefs); response.Error != nil {
			return utils.SendError(ctx, "Error occurred while retrieving apartmentPrefs from the database: "+response.Error.Error(), fiber.StatusInternalServerError)
		}
		err := ctx.JSON(ApartmentPrefs)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of apartmentPrefs: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Return a single apartmentPref as JSON
func GetApartmentPref(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ApartmentPref := new(models.ApartmentPref)
		id := ctx.Params("id")
		if response := db.Find(&ApartmentPref, id); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when retrieving the apartmentPref: "+response.Error.Error(), fiber.StatusInternalServerError)
		}
		if ApartmentPref.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				return utils.SendError(ctx, "Cannot return status not found: "+err.Error(), fiber.StatusInternalServerError)
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				return utils.SendError(ctx, "Error occurred when returning JSON of a role: "+err.Error(), fiber.StatusInternalServerError)
			}
			return err
		}
		err := ctx.JSON(ApartmentPref)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a apartmentPref: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Add a single apartmentPref to the database
func AddApartmentPref(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ApartmentPref := new(models.ApartmentPref)
		if err := ctx.BodyParser(ApartmentPref); err != nil {
			return utils.SendError(ctx, "an error occurred when parsing the new apartmentPref", fiber.StatusBadRequest)
		}
		if response := db.Create(&ApartmentPref); response.Error != nil {
			return utils.SendError(ctx, "an error occurred when storing the new apartmentPref"+response.Error.Error(), fiber.StatusInternalServerError)
		}
		err := ctx.JSON(ApartmentPref)
		if err != nil {
			return utils.SendError(ctx, "error occurred when returning JSON of a apartmentPref", fiber.StatusInternalServerError)
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
			return utils.SendError(ctx, "An error occurred when parsing the edited apartmentPref: "+err.Error(), fiber.StatusBadRequest)
		}
		if response := db.Find(&ApartmentPref, id); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when retrieving the existing apartmentPref: "+response.Error.Error(), fiber.StatusNotFound)
		}
		// ApartmentPref does not exist
		if ApartmentPref.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				return utils.SendError(ctx, "Cannot return status not found: "+err.Error(), fiber.StatusInternalServerError)
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				return utils.SendError(ctx, "Error occurred when returning JSON of a apartmentPref: "+err.Error(), fiber.StatusInternalServerError)
			}
			return err
		}
		ApartmentPref.NumberOfRooms = EditApartmentPref.NumberOfRooms
		ApartmentPref.Price = EditApartmentPref.Price
		ApartmentPref.Balcony = EditApartmentPref.Balcony
		ApartmentPref.Roof = EditApartmentPref.Roof
		ApartmentPref.Parking = EditApartmentPref.Parking
		ApartmentPref.Elevator = EditApartmentPref.Elevator
		ApartmentPref.AnimalsAllowed = EditApartmentPref.AnimalsAllowed
		ApartmentPref.Renovated = EditApartmentPref.Renovated
		ApartmentPref.AvailableDate = EditApartmentPref.AvailableDate
		ApartmentPref.Location = EditApartmentPref.Location

		// Save apartmentPref
		db.Save(&ApartmentPref)

		err := ctx.JSON(ApartmentPref)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a apartmentPref: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Delete a single apartmentPref
func DeleteApartmentPref(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var ApartmentPref models.ApartmentPref
		db.Find(&ApartmentPref, id)
		if response := db.Find(&ApartmentPref); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when finding the apartmentPref to be deleted"+response.Error.Error(), fiber.StatusNotFound)
		}
		db.Delete(&ApartmentPref)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a apartmentPref: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}
