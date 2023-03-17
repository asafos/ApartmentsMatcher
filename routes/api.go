package routes

import (
	Controller "fiber-boilerplate/app/controllers/api"
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(api fiber.Router, db *database.Database) {
	registerRoles(api, db)
	registerUsers(api, db)
	registerApartments(api, db)
	registerApartmentPrefs(api, db)
	registerMatching(api, db)
}

func registerRoles(api fiber.Router, db *database.Database) {
	roles := api.Group("/roles")

	roles.Get("/", Controller.GetAllRoles(db))
	roles.Get("/:id", Controller.GetRole(db))
	roles.Post("/", Controller.AddRole(db))
	roles.Put("/:id", Controller.EditRole(db))
	roles.Delete("/:id", Controller.DeleteRole(db))
}

func registerUsers(api fiber.Router, db *database.Database) {
	users := api.Group("/users")

	users.Get("/", Controller.GetAllUsers(db))
	users.Get("/:id", Controller.GetUser(db))
	users.Post("/", Controller.AddUser(db))
	users.Put("/:id", Controller.EditUser(db))
	users.Delete("/:id", Controller.DeleteUser(db))
}

func registerApartments(api fiber.Router, db *database.Database) {
	users := api.Group("/apartments")

	users.Get("/", Controller.GetAllApartments(db))
	users.Get("/:id", Controller.GetApartment(db))
	users.Post("/", Controller.AddApartment(db))
	users.Put("/:id", Controller.EditApartment(db))
	users.Delete("/:id", Controller.DeleteApartment(db))
}

func registerApartmentPrefs(api fiber.Router, db *database.Database) {
	users := api.Group("/apartmentPrefs")

	users.Get("/", Controller.GetAllApartmentPrefs(db))
	users.Get("/:id", Controller.GetApartmentPref(db))
	users.Post("/", Controller.AddApartmentPref(db))
	users.Put("/:id", Controller.EditApartmentPref(db))
	users.Delete("/:id", Controller.DeleteApartmentPref(db))
}

func registerMatching(api fiber.Router, db *database.Database) {
	users := api.Group("/match")

	users.Get("/user/:id", Controller.GetMatchingApartments(db))
}
