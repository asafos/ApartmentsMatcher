package routes

import (
	Controller "fiber-boilerplate/app/controllers/web"
	configuration "fiber-boilerplate/config"

	"github.com/gofiber/session/v2"

	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func RegisterAuth(api fiber.Router, config configuration.Config, session *session.Session, db *database.Database, sessionLookup string) {
	goth.UseProviders(
		google.New(config.GetString("OAUTH_GOOGLE_CLIENT_ID"), config.GetString("OAUTH_GOOGLE_SECRET"), config.GetString("OAUTH_GOOGLE_REDIRECT_URI")),
		facebook.New(config.GetString("OAUTH_FACEBOOK_CLIENT_ID"), config.GetString("OAUTH_FACEBOOK_SECRET"), config.GetString("OAUTH_FACEBOOK_REDIRECT_URI")),
	)

	api.Get("/:provider/callback", Controller.OAuthLoginCallback(session, db))

	api.Get("/logout/:provider", Controller.OAuthLogout(session, sessionLookup))

	api.Get("/:provider", Controller.OAuthLogin())

	api.Get("/isAuthenticated", Controller.IsAuthenticatedHandler(session, db))
}
