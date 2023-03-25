package middleware

import (
	"fiber-boilerplate/app/constants"
	authServices "fiber-boilerplate/app/services/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

// auth requires user login via oauth
func Auth(session *session.Session, jwtSecret string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// jwtCookie := ctx.Cookies(constants.JWT_COOKIE_NAME)
		// store := session.Get(ctx)
		// // parse the JWT cookie
		// token, err := jwt.Parse(jwtCookie, func(token *jwt.Token) (interface{}, error) {
		// 	// validate the signing method
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fiber.ErrBadRequest
		// 	}

		// 	return jwtSecret, nil
		// })
		// if err != nil {
		// 	return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		// }

		// // validate the JWT claims
		// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 	// check if the session value matches the JWT claim
		// 	if store.Get(constants.USER_ID_SESSION_KEY) != claims[constants.USER_ID_SESSION_KEY] {
		// 		return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		// 	}

		// 	// set the session value
		// 	// store.Set("authenticated", true)
		// 	// store.Save()
		// 	return ctx.Next()
		// }

		// return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")

		// gf.GetState(ctx)

		userID, ok := authServices.IsAuthenticated(session, ctx, jwtSecret)
		if !ok {
			ctx.SendStatus(fiber.StatusUnauthorized)
			return nil
		}
		ctx.Locals(constants.USER_LOCALS_KEY, userID)
		return ctx.Next()
	}
}
