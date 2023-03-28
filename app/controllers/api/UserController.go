package api

import (
	utils "fiber-boilerplate/app/controllers/utils"
	"fiber-boilerplate/app/models"
	services "fiber-boilerplate/app/services/api"
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

// Return all users as JSON
func GetAllUsers(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Users []models.User
		if response := services.GetAllUsers(db, &Users); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred while retrieving users from the database: "+response.Error.Error())
		}
		err := ctx.JSON(Users)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of users: "+err.Error())
		}
		return err
	}
}

// Return a single user as JSON
func GetUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(models.User)
		id := ctx.Params("id")
		if response := services.GetUser(db, User, id); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "An error occurred when retrieving the user: "+response.Error.Error())
		}
		if User.ID == 0 {
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
		// userID := ctx.Locals(constants.USER_LOCALS_KEY)
		// if userID != User.ID {
		// 	return fiber.fiber.StatusUnauthorized, NewError("Unauthorized user")
		// }
		err := ctx.JSON(User)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a user: "+err.Error())
		}
		return err
	}
}

// Add a single user to the database
func AddUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(models.User)
		if err := ctx.BodyParser(User); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "an error occurred when parsing the new user")
		}
		if response := services.AddUser(db, User); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "an error occurred when storing the new user"+response.Error.Error())
		}
		errors := utils.ValidateStruct(*User)
		if errors != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(errors)
		}
		err := ctx.JSON(User)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "error occurred when returning JSON of a user")
		}
		return err
	}
}

// Edit a single user
func EditUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditUser := new(models.User)
		User := new(models.User)
		if err := ctx.BodyParser(EditUser); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "An error occurred when parsing the edited user: "+err.Error())
		}
		errors := utils.ValidateStruct(*EditUser)
		if errors != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(errors)
		}
		if response := services.GetUser(db, User, id); response.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, "An error occurred when retrieving the existing user: "+response.Error.Error())
		}
		// User does not exist
		if User.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Cannot return status not found: "+err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a user: "+err.Error())
			}
			return err
		}
		// userID := ctx.Locals(constants.USER_LOCALS_KEY)
		// if userID != User.ID {
		// 	return fiber.fiber.StatusUnauthorized, NewError("Unauthorized user")
		// }
		User.Name = EditUser.Name
		User.Email = EditUser.Email
		User.RoleID = EditUser.RoleID
		// Save user
		services.EditUser(db, User)

		err := ctx.JSON(User)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a user: "+err.Error())
		}
		return err
	}
}

// Delete a single user
func DeleteUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var User models.User
		// services.Find(&User, id)
		if response := services.GetUser(db, &User, id); response.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, "An error occurred when finding the user to be deleted"+response.Error.Error())
		}
		// userID := ctx.Locals(constants.USER_LOCALS_KEY)
		// if userID != User.ID {
		// 	return fiber.fiber.StatusUnauthorized, NewError("Unauthorized user")
		// }
		services.DeleteUser(db, &User)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred when returning JSON of a user: "+err.Error())
		}
		return err
	}
}
