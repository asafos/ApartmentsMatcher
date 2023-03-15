package api

import (
	utils "fiber-boilerplate/app/controllers/utils"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

// Return all users as JSON
func GetAllUsers(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Users []models.User
		if response := db.Preload("ApartmentPrefs").Preload("Apartments").Find(&Users); response.Error != nil {
			return utils.SendError(ctx, "Error occurred while retrieving users from the database: "+response.Error.Error(), fiber.StatusInternalServerError)
		}
		// Match roles to users
		for index, User := range Users {
			if User.RoleID != 0 {
				Role := new(models.Role)
				if response := db.Find(&Role, User.RoleID); response.Error != nil {
					return utils.SendError(ctx, "An error occurred when retrieving the role: "+response.Error.Error(), fiber.StatusInternalServerError)
				}
				if Role.ID != 0 {
					Users[index].Role = *Role
				}
			}
		}
		err := ctx.JSON(Users)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of users: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Return a single user as JSON
func GetUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(models.User)
		id := ctx.Params("id")
		if response := db.Preload("ApartmentPrefs").Preload("Apartments").Find(&User, id); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when retrieving the user: "+response.Error.Error(), fiber.StatusInternalServerError)
		}
		if User.ID == 0 {
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
		// Match role to user
		if User.RoleID != 0 {
			Role := new(models.Role)
			if response := db.Find(&Role, User.RoleID); response.Error != nil {
				return utils.SendError(ctx, "An error occurred when retrieving the role: "+response.Error.Error(), fiber.StatusInternalServerError)
			}
			if Role.ID != 0 {
				User.Role = *Role
			}
		}
		err := ctx.JSON(User)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a user: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Add a single user to the database
func AddUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(models.User)
		if err := ctx.BodyParser(User); err != nil {
			return utils.SendError(ctx, "an error occurred when parsing the new user", fiber.StatusBadRequest)
		}
		if response := db.Create(&User); response.Error != nil {
			return utils.SendError(ctx, "an error occurred when storing the new user"+response.Error.Error(), fiber.StatusInternalServerError)
		}
		// Match role to user
		if User.RoleID != 0 {
			Role := new(models.Role)
			if response := db.Find(&Role, User.RoleID); response.Error != nil {
				return utils.SendError(ctx, "an error occurred when retrieving the role"+response.Error.Error(), fiber.StatusInternalServerError)
			}
			if Role.ID != 0 {
				User.Role = *Role
			}
		}
		err := ctx.JSON(User)
		if err != nil {
			return utils.SendError(ctx, "error occurred when returning JSON of a user", fiber.StatusInternalServerError)
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
			return utils.SendError(ctx, "An error occurred when parsing the edited user: "+err.Error(), fiber.StatusBadRequest)
		}
		if response := db.Find(&User, id); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when retrieving the existing user: "+response.Error.Error(), fiber.StatusNotFound)
		}
		// User does not exist
		if User.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				return utils.SendError(ctx, "Cannot return status not found: "+err.Error(), fiber.StatusInternalServerError)
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				return utils.SendError(ctx, "Error occurred when returning JSON of a user: "+err.Error(), fiber.StatusInternalServerError)
			}
			return err
		}
		User.Name = EditUser.Name
		User.Email = EditUser.Email
		User.RoleID = EditUser.RoleID
		// Match role to user
		if User.RoleID != 0 {
			Role := new(models.Role)
			if response := db.Find(&Role, User.RoleID); response.Error != nil {
				return utils.SendError(ctx, "An error occurred when retrieving the role"+response.Error.Error(), fiber.StatusBadRequest)
			}
			if Role.ID != 0 {
				User.Role = *Role
			}
		}
		// Save user
		db.Save(&User)

		err := ctx.JSON(User)
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a user: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}

// Delete a single user
func DeleteUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var User models.User
		db.Find(&User, id)
		if response := db.Find(&User); response.Error != nil {
			return utils.SendError(ctx, "An error occurred when finding the user to be deleted"+response.Error.Error(), fiber.StatusNotFound)
		}
		db.Delete(&User)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			return utils.SendError(ctx, "Error occurred when returning JSON of a user: "+err.Error(), fiber.StatusInternalServerError)
		}
		return err
	}
}
