package api

import (
	utils "fiber-boilerplate/app/controllers/utils"
	"fiber-boilerplate/app/models"
	services "fiber-boilerplate/app/services/api"
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

// Return all roles as JSON
func GetAllRoles(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Roles []models.Role
		if response := services.GetAllRoles(db, &Roles); response.Error != nil {
			utils.SendError(ctx, "Error occurred while retrieving roles from the database: "+response.Error.Error(), fiber.StatusInternalServerError)
		}
		err := ctx.JSON(Roles)
		if err != nil {
			utils.SendError(ctx, "Error occurred when returning JSON of roles: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Return a single role as JSON
func GetRole(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Role := new(models.Role)
		id := ctx.Params("id")
		if response := services.GetRole(db, Role, id); response.Error != nil {
			utils.SendError(ctx, "An error occurred when retrieving the role: "+response.Error.Error(), fiber.StatusBadRequest)
		}
		if Role.ID == 0 {
			// Send status not found
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				utils.SendError(ctx, "Cannot return status not found: "+err.Error(), fiber.StatusInternalServerError)
			}
			// Set ID
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				utils.SendError(ctx, "Error occurred when returning JSON of a role: "+err.Error(), fiber.StatusInternalServerError)
			}
			return err
		}
		err := ctx.JSON(Role)
		if err != nil {
			utils.SendError(ctx, "Error occurred when returning JSON of a role: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Add a single role to the database
func AddRole(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Role := new(models.Role)
		if err := ctx.BodyParser(Role); err != nil {
			utils.SendError(ctx, "An error occurred when parsing the new role: "+err.Error(), fiber.StatusBadRequest)
		}
		if response := services.AddRole(db, Role); response.Error != nil {
			utils.SendError(ctx, "An error occurred when storing the new role: "+response.Error.Error(), fiber.StatusInternalServerError)
		}
		err := ctx.JSON(Role)
		if err != nil {
			utils.SendError(ctx, "Error occurred when returning JSON of a role: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Edit a single role
func EditRole(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditRole := new(models.Role)
		Role := new(models.Role)
		if err := ctx.BodyParser(EditRole); err != nil {
			utils.SendError(ctx, "An error occurred when parsing the edited role: "+err.Error(), fiber.StatusBadRequest)
		}
		if response := services.GetRole(db, Role, id); response.Error != nil {
			utils.SendError(ctx, "An error occurred when retrieving the existing role: "+response.Error.Error(), fiber.StatusNotFound)
		}
		// Role does not exist
		if Role.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				utils.SendError(ctx, "Cannot return status not found: "+err.Error(), fiber.StatusInternalServerError)
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				utils.SendError(ctx, "Error occurred when returning JSON of a role: "+err.Error(), fiber.StatusInternalServerError)
			}
			return err
		}
		Role.Name = EditRole.Name
		Role.Description = EditRole.Description
		services.EditRole(db, Role)

		err := ctx.JSON(Role)
		if err != nil {
			utils.SendError(ctx, "Error occurred when returning JSON of a role: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Delete a single role
func DeleteRole(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var Role models.Role
		// services.Find(&Role, id)
		if response := services.GetRole(db, &Role, id); response.Error != nil {
			utils.SendError(ctx, "An error occurred when finding the role to be deleted: "+response.Error.Error(), fiber.StatusNotFound)
		}
		services.DeleteRole(db, &Role)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			utils.SendError(ctx, "Error occurred when returning JSON of a role: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}
