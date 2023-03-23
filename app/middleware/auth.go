package middleware

import (
	Controller "fiber-boilerplate/app/controllers/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"

	gf "github.com/shareed2k/goth_fiber"
)

// auth requires user login via oauth
func Auth(session *session.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		gf.GetState(ctx)

		userID, ok := Controller.IsAuthenticated(session, ctx)
		if !ok {
			ctx.SendStatus(fiber.StatusUnauthorized)
			return nil
		}
		ctx.Locals("userID", userID)
		return ctx.Next()
	}
}
