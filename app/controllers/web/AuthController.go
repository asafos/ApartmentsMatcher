package web

import (
	"bytes"
	"encoding/json"
	"errors"
	utils "fiber-boilerplate/app/controllers/utils"
	"fiber-boilerplate/app/models"
	services "fiber-boilerplate/app/services/api"
	configuration "fiber-boilerplate/config"
	"fiber-boilerplate/database"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

func IsAuthenticated(session *session.Session, ctx *fiber.Ctx) (userID uint, authenticated bool) {
	return utils.GetUserIDFromSession(session, ctx)
}

func IsAuthenticatedHandler(session *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID, ok := utils.GetUserIDFromSession(session, ctx)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "user not logged in")
		}
		User := new(models.User)
		if response := services.GetUser(db, User, string(userID)); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "An error occurred when retrieving the user: "+response.Error.Error())
		}
		return ctx.JSON(User)
	}
}

func IsAdmin(session *session.Session, ctx *fiber.Ctx, db *database.Database) (authenticated bool) {
	userID, ok := utils.GetUserIDFromSession(session, ctx)
	if !ok {
		return false
	}
	user, err := FindUserByID(db, userID)
	if err != nil || user == nil {
		return false
	}
	return models.RoleEnum(user.RoleID) == models.AdminRole
}

type GoogleOauthToken struct {
	Access_token string
	Id_token     string
}

func GetGoogleOauthToken(code string, config configuration.Config) (*GoogleOauthToken, error) {
	const rootURl = "https://oauth2.googleapis.com/token"

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", config.GetString("OAUTH_GOOGLE_CLIENT_ID"))
	values.Add("client_secret", config.GetString("OAUTH_GOOGLE_SECRET"))
	values.Add("redirect_uri", config.GetString("OAUTH_GOOGLE_REDIRECT_URI"))

	query := values.Encode()

	req, err := http.NewRequest("POST", rootURl, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		resBody, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New("could not retrieve token" + string(resBody))
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleOauthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleOauthTokenRes); err != nil {
		return nil, err
	}

	tokenBody := &GoogleOauthToken{
		Access_token: GoogleOauthTokenRes["access_token"].(string),
		Id_token:     GoogleOauthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

type GoogleUserResult struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
}

func GetGoogleUser(access_token string, id_token string) (*GoogleUserResult, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", access_token)

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", id_token))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return nil, err
	}

	userBody := &GoogleUserResult{
		Id:             GoogleUserRes["id"].(string),
		Email:          GoogleUserRes["email"].(string),
		Verified_email: GoogleUserRes["verified_email"].(bool),
		Name:           GoogleUserRes["name"].(string),
		Given_name:     GoogleUserRes["given_name"].(string),
		Picture:        GoogleUserRes["picture"].(string),
		Locale:         GoogleUserRes["locale"].(string),
	}

	return userBody, nil
}

func AuthorizeGoogle(config configuration.Config, session *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		code := ctx.FormValue("code")
		googleOauthToken, err := GetGoogleOauthToken(code, config)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		googleOauthUser, err := GetGoogleUser(googleOauthToken.Access_token, googleOauthToken.Id_token)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		User := new(models.User)
		User.Email = googleOauthUser.Email
		User.Name = googleOauthUser.Name
		User.OAuthID = googleOauthUser.Id
		User.RoleID = uint(models.UserRole)
		if response := services.AddUser(db, User); response.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "an error occurred when storing the new user"+response.Error.Error())
		}

		store := session.Get(ctx)
		defer store.Save()
		store.Set("userid", User.ID)

		ctx.JSON(User)
		return nil
	}
}

// func OAuthLogin() fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		if gothUser, err := gf.CompleteUserAuth(ctx); err == nil {
// 			ctx.JSON(gothUser)
// 		} else {
// 			gf.BeginAuthHandler(ctx)
// 		}
// 		return nil
// 	}
// }

// func OAuthLoginCallback(session *session.Session, db *database.Database) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		oAuthUser, err := gf.CompleteUserAuth(ctx)
// 		if err != nil {
// 			return err
// 		}

// 		user, err := FindUserByOAuthID(db, oAuthUser.UserID)
// 		if err == nil && user != nil {
// 			ctx.JSON(user)
// 			return nil
// 		}

// 		User := new(models.User)
// 		User.Email = oAuthUser.Email
// 		User.Name = oAuthUser.Name
// 		User.OAuthID = oAuthUser.UserID
// 		User.RoleID = uint(models.UserRole)
// 		if response := services.AddUser(db, User); response.Error != nil {
// 			return fiber.NewError(fiber.StatusInternalServerError, "an error occurred when storing the new user"+response.Error.Error())
// 		}

// 		store := session.Get(ctx)
// 		defer store.Save()
// 		store.Set("userid", User.ID)

// 		ctx.JSON(User)
// 		return nil
// 	}
// }

// func OAuthLogout(session *session.Session, sessionLookup string) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		gf.Logout(ctx)
// 		store := session.Get(ctx)
// 		store.Delete("userid")
// 		if err := store.Save(); err != nil {
// 			ctx.Status(fiber.StatusInternalServerError).SendString("couldn't delete user from store" + err.Error())
// 			return fiber.NewError(fiber.StatusInternalServerError, "couldn't delete user from store"+err.Error())
// 		}
// 		// Check if cookie needs to be unset
// 		split := strings.Split(sessionLookup, ":")
// 		if strings.ToLower(split[0]) == "cookie" {
// 			// Unset cookie on client-side
// 			ctx.Set("Set-Cookie", split[1]+"=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/; HttpOnly")
// 			if err := ctx.SendString("You are now logged out."); err != nil {
// 				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 			}
// 			return nil
// 		}
// 		return nil
// 	}
// }

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
