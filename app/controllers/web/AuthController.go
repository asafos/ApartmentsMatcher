package web

import (
	utils "fiber-boilerplate/app/controllers/utils"
	"fiber-boilerplate/database"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	gf "github.com/shareed2k/goth_fiber"
	hashing "github.com/thomasvvugt/fiber-hashing"
)

func IsAuthenticated(session *session.Session, ctx *fiber.Ctx) (authenticated bool) {
	store := session.Get(ctx)
	// Get User ID from session store
	userID, correct := store.Get("userid").(int64)
	return correct && userID > 0
}

func OAuthLogin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if gothUser, err := gf.CompleteUserAuth(ctx); err == nil {
			ctx.JSON(gothUser)
		} else {
			gf.BeginAuthHandler(ctx)
		}
		return nil
	}
}

func OAuthLoginCallback(session *session.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, err := gf.CompleteUserAuth(ctx)
		if err != nil {
			return err
		}
		store := session.Get(ctx)
		defer store.Save()
		// Set the user ID in the session store
		store.Set("userid", user.UserID)
		ctx.JSON(user)
		return nil
	}
}

func OAuthLogout(session *session.Session) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		gf.Logout(ctx)
		store := session.Get(ctx)
		store.Delete("userid")
		if err := store.Save(); err != nil {
			ctx.Status(fiber.StatusInternalServerError).SendString("couldn't delete user from store" + err.Error())
		}
		return nil
	}
}

func ShowLoginForm() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Render("login", fiber.Map{})
		if err != nil {
			if err2 := ctx.Status(500).SendString(err.Error()); err2 != nil {
				return utils.SendError(ctx, err2.Error(), fiber.StatusInternalServerError)
			}
		}
		return err
	}
}

func PostLoginForm(hasher hashing.Driver, session *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		username := ctx.FormValue("username")
		// Find user
		user, err := FindUserByUsername(db, username)
		if err != nil {
			return utils.SendError(ctx, "User not found", fiber.StatusNotFound)
		}

		// Check if password matches hash
		if hasher != nil {
			password := ctx.FormValue("password")
			match, err := hasher.MatchHash(password, user.Password)
			if err != nil {
				return utils.SendError(ctx, "Password parsing failed", fiber.StatusInternalServerError)
			}
			if match {
				store := session.Get(ctx)
				defer store.Save()
				// Set the user ID in the session store
				store.Set("userid", user.ID)
				fmt.Printf("User set in session store with ID: %v\n", user.ID)
				if err := ctx.SendString("You should be logged in successfully!"); err != nil {
					return utils.SendError(ctx, err.Error(), fiber.StatusInternalServerError)
				}
			} else {
				if err := ctx.SendString("The entered details do not match our records."); err != nil {
					return utils.SendError(ctx, err.Error(), fiber.StatusInternalServerError)
				}
			}
		} else {
			panic("Hash provider was not set")
		}
		return nil
	}
}

func PostLogoutForm(sessionLookup string, session *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if IsAuthenticated(session, ctx) {
			store := session.Get(ctx)
			store.Delete("userid")
			if err := store.Save(); err != nil {
				panic(err.Error())
			}
			// Check if cookie needs to be unset
			split := strings.Split(sessionLookup, ":")
			if strings.ToLower(split[0]) == "cookie" {
				// Unset cookie on client-side
				ctx.Set("Set-Cookie", split[1]+"=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/; HttpOnly")
				if err := ctx.SendString("You are now logged out."); err != nil {
					return utils.SendError(ctx, err.Error(), fiber.StatusInternalServerError)
				}
				return nil
			}
			return nil
		}
		// TODO: Redirect?
		return nil
	}
}
