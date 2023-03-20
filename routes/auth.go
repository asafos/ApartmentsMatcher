package routes

import (
	Controller "fiber-boilerplate/app/controllers/web"
	configuration "fiber-boilerplate/config"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func RegisterAuth(api fiber.Router, config configuration.Config) {
	goth.UseProviders(
		google.New(config.GetString("OAUTH_GOOGLE_CLIENT_ID"), config.GetString("OAUTH_GOOGLE_SECRET"), config.GetString("OAUTH_GOOGLE_REDIRECT_URI")),
		facebook.New(config.GetString("OAUTH_FACEBOOK_CLIENT_ID"), config.GetString("OAUTH_FACEBOOK_SECRET"), config.GetString("OAUTH_FACEBOOK_REDIRECT_URI")),
	)
	registerOAuth(api)
}

func registerOAuth(api fiber.Router) {
	api.Get("/:provider/callback", Controller.OAuthLoginCallback())

	api.Get("/logout/:provider", Controller.OAuthLogout())

	api.Get("/:provider", Controller.OAuthLogin())
}
