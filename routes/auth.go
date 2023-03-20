package routes

import (
	Controller "fiber-boilerplate/app/controllers/web"
	configuration "fiber-boilerplate/config"

	"github.com/gofiber/session/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func RegisterAuth(api fiber.Router, config configuration.Config, session *session.Session) {
	goth.UseProviders(
		google.New(config.GetString("OAUTH_GOOGLE_CLIENT_ID"), config.GetString("OAUTH_GOOGLE_SECRET"), config.GetString("OAUTH_GOOGLE_REDIRECT_URI")),
		facebook.New(config.GetString("OAUTH_FACEBOOK_CLIENT_ID"), config.GetString("OAUTH_FACEBOOK_SECRET"), config.GetString("OAUTH_FACEBOOK_REDIRECT_URI")),
	)

	api.Get("/:provider/callback", Controller.OAuthLoginCallback(session))

	api.Get("/logout/:provider", Controller.OAuthLogout(session))

	api.Get("/:provider", Controller.OAuthLogin())
}
