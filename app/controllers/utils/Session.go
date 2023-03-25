package controllerUtils

import (
	"fiber-boilerplate/app/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

func GetUserIDFromSession(session *session.Session, ctx *fiber.Ctx) (response uint, ok bool) {
	store := session.Get(ctx)
	userID := store.Get(constants.USER_ID_SESSION_KEY)
	if userID == nil {
		return 0, false
	}
	res, ok := userID.(uint)

	return res, ok
}
