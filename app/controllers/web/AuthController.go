package web

import (
	"fiber-boilerplate/app/models"
	apiServices "fiber-boilerplate/app/services/api"
	authServices "fiber-boilerplate/app/services/auth"
	configuration "fiber-boilerplate/config"
	"fiber-boilerplate/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

func IsAuthenticatedHandler(session *session.Session, db *database.Database, jwtSecret string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID, ok := authServices.IsAuthenticated(session, ctx, jwtSecret)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "user not logged in")
		}
		User := new(models.User)

		if response := apiServices.GetUser(db, User, fmt.Sprint(userID)); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "An error occurred when retrieving the user: "+response.Error.Error())
		}
		return ctx.JSON(User)
	}
}

func AuthorizeGoogle(config configuration.Config, session *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		code := ctx.FormValue("code")
		oauthToken, err := authServices.GetGoogleOauthToken(code, config)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		oauthUser, err := authServices.GetGoogleUser(oauthToken.Access_token, oauthToken.Id_token)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		User, err := authServices.AddOAuthUser(ctx, session, db, oauthUser, config.GetString("JWT_SECRET"))
		if err != nil {
			return err
		}

		ctx.JSON(User)
		return nil
	}
}

func AuthorizeFacebook(config configuration.Config, session *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		code := ctx.FormValue("code")

		facebookOAuthCode, err := authServices.ParseFacebookSignedRequest(code, config.GetString("OAUTH_FACEBOOK_SECRET"))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		oauthToken, err := authServices.GetFacebookOauthToken(facebookOAuthCode.Code, config)
		fmt.Println(oauthToken)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		oauthUser, err := authServices.GetFacebookUser(oauthToken.Access_token)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		User, err := authServices.AddOAuthUser(ctx, session, db, oauthUser, config.GetString("JWT_SECRET"))
		if err != nil {
			return err
		}

		ctx.JSON(User)
		return nil
	}
}

// ---------------- Old login ----------------

// func ShowLoginForm() fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		err := ctx.Render("login", fiber.Map{})
// 		if err != nil {
// 			if err2 := ctx.Status(500).SendString(err.Error()); err2 != nil {
// 				return utils.SendError(ctx, err2.Error(), fiber.StatusInternalServerError)
// 			}
// 		}
// 		return err
// 	}
// }

// func PostLoginForm(hasher hashing.Driver, session *session.Session, db *database.Database) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		username := ctx.FormValue("username")
// 		// Find user
// 		user, err := FindUserByUsername(db, username)
// 		if err != nil {
// 			return utils.SendError(ctx, "User not found", fiber.StatusNotFound)
// 		}

// 		// Check if password matches hash
// 		if hasher != nil {
// 			// password := ctx.FormValue("password")
// 			// match, err := hasher.MatchHash(password, user.Password)
// 			// if err != nil {
// 			// 	return utils.SendError(ctx, "Password parsing failed", fiber.StatusInternalServerError)
// 			// }
// 			// if match {
// 			store := session.Get(ctx)
// 			defer store.Save()
// 			// Set the user ID in the session store
// 			store.Set("userid", user.ID)
// 			fmt.Printf("User set in session store with ID: %v\n", user.ID)
// 			if err := ctx.SendString("You should be logged in successfully!"); err != nil {
// 				return utils.SendError(ctx, err.Error(), fiber.StatusInternalServerError)
// 			}
// 			// } else {
// 			// 	if err := ctx.SendString("The entered details do not match our records."); err != nil {
// 			// 		return utils.SendError(ctx, err.Error(), fiber.StatusInternalServerError)
// 			// 	}
// 			// }
// 		} else {
// 			panic("Hash provider was not set")
// 		}
// 		return nil
// 	}
// }

// func PostLogoutForm(sessionLookup string, session *session.Session, db *database.Database) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		if IsAuthenticated(session, ctx) {
// 			store := session.Get(ctx)
// 			store.Delete("userid")
// 			if err := store.Save(); err != nil {
// 				panic(err.Error())
// 			}
// 			// Check if cookie needs to be unset
// 			split := strings.Split(sessionLookup, ":")
// 			if strings.ToLower(split[0]) == "cookie" {
// 				// Unset cookie on client-side
// 				ctx.Set("Set-Cookie", split[1]+"=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/; HttpOnly")
// 				if err := ctx.SendString("You are now logged out."); err != nil {
// 					return utils.SendError(ctx, err.Error(), fiber.StatusInternalServerError)
// 				}
// 				return nil
// 			}
// 			return nil
// 		}
// 		// TODO: Redirect?
// 		return nil
// 	}
// }
