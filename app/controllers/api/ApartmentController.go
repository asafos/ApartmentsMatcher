package api

import (
	utils "fiber-boilerplate/app/controllers/utils"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

// Return all apartments as JSON
func GetAllApartments(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Apartments []models.Apartment
		if response := db.Preload("ApartmentPrefs").Preload("Apartments").Find(&Apartments); response.Error != nil {
			return utils.SendError(ctx, "Error occurred while retrieving apartments from the database: "+response.Error.Error(), fiber.StatusInternalServerError)
		}
		err := ctx.JSON(Apartments)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of apartments: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Return a single apartment as JSON
func GetApartment(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Apartment := new(models.Apartment)
		id := ctx.Params("id")
		if response := db.Find(&Apartment, id); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when retrieving the apartment: "+response.Error.Error(), fiber.StatusInternalServerError)
		}
		if Apartment.ID == 0 {
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
		err := ctx.JSON(Apartment)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a apartment: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Add a single apartment to the database
func AddApartment(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Apartment := new(models.Apartment)
		if err := ctx.BodyParser(Apartment); err != nil {
			return utils.SendError(ctx, "an error occurred when parsing the new apartment", fiber.StatusBadRequest)
		}
		if response := db.Create(&Apartment); response.Error != nil {
			return utils.SendError(ctx, "an error occurred when storing the new apartment"+response.Error.Error(), fiber.StatusInternalServerError)
		}
		err := ctx.JSON(Apartment)
		if err != nil {
			return utils.SendError(ctx, "error occurred when returning JSON of a apartment", fiber.StatusInternalServerError)
		}
		return err
	}
}

// Edit a single apartment
func EditApartment(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditApartment := new(models.Apartment)
		Apartment := new(models.Apartment)
		if err := ctx.BodyParser(EditApartment); err != nil {
			return utils.SendError(ctx, "An error occurred when parsing the edited apartment: "+err.Error(), fiber.StatusBadRequest)
		}
		if response := db.Find(&Apartment, id); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when retrieving the existing apartment: "+response.Error.Error(), fiber.StatusNotFound)
		}
		// Apartment does not exist
		if Apartment.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				return utils.SendError(ctx, "Cannot return status not found: "+err.Error(), fiber.StatusInternalServerError)
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				return utils.SendError(ctx, "Error occurred when returning JSON of a apartment: "+err.Error(), fiber.StatusInternalServerError)
			}
			return err
		}
		Apartment.NumberOfRooms = EditApartment.NumberOfRooms
		Apartment.Price = EditApartment.Price
		Apartment.Balcony = EditApartment.Balcony
		Apartment.Roof = EditApartment.Roof
		Apartment.Parking = EditApartment.Parking
		Apartment.Elevator = EditApartment.Elevator
		Apartment.AnimalsAllowed = EditApartment.AnimalsAllowed
		Apartment.Renovated = EditApartment.Renovated
		Apartment.AvailableDate = EditApartment.AvailableDate
		Apartment.Location = EditApartment.Location

		// Save apartment
		db.Save(&Apartment)

		err := ctx.JSON(Apartment)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a apartment: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Delete a single apartment
func DeleteApartment(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var Apartment models.Apartment
		db.Find(&Apartment, id)
		if response := db.Find(&Apartment); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when finding the apartment to be deleted"+response.Error.Error(), fiber.StatusNotFound)
		}
		db.Delete(&Apartment)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a apartment: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}
