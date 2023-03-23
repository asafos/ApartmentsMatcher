package routes

import (
	"fiber-boilerplate/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	hashing "github.com/thomasvvugt/fiber-hashing"
)

func RegisterWeb(web fiber.Router, session *session.Session, sessionLookup string, db *database.Database, hasher hashing.Driver) {
	// Homepage
	// web.Get("/", Controller.Index(session, db))

	// Make a new hash
	web.Get("/hash/*", func(ctx *fiber.Ctx) error {
		hash, err := hasher.CreateHash(ctx.Params("*"))
		if err != nil {
			log.Fatalf("Error when creating hash: %v", err)
		}
		if err := ctx.SendString(hash); err != nil {
			panic(err.Error())
		}
		return err
	})

	// Auth routes
	// web.Get("/login", Controller.ShowLoginForm())
	// web.Post("/login", Controller.PostLoginForm(hasher, session, db))
	// web.Post("/logout", Controller.PostLogoutForm(sessionLookup, session, db))
}
