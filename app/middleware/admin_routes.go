package middleware

import (
	Controller "fiber-boilerplate/app/controllers/web"
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"

	gf "github.com/shareed2k/goth_fiber"
)

// admin routes requires user login via oauth and admin role
func AdminRole(session *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		gf.GetState(ctx)
		if !Controller.IsAdmin(session, ctx, db) {
			ctx.SendStatus(fiber.StatusUnauthorized)
			return nil
		}
		return ctx.Next()
	}
}
