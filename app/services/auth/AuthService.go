package serviceAuth

import (
	"errors"
	"fiber-boilerplate/app/constants"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/session/v2"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(session *session.Session, ctx *fiber.Ctx, jwtSecret string) (userID uint, authenticated bool) {
	jwtCookie := ctx.Cookies(constants.JWT_COOKIE_NAME)
	store := session.Get(ctx)
	// parse the JWT cookie
	token, err := jwt.Parse(jwtCookie, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrBadRequest
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		return 0, false
	}

	// validate the JWT claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFromSession, userIDFromSessionOk := store.Get(constants.USER_ID_SESSION_KEY).(int64)
		userIDFromCookie, userIDFromCookieOk := claims[constants.USER_ID_SESSION_KEY].(float64)
		// check if the session value matches the JWT claim
		if !userIDFromSessionOk || !userIDFromCookieOk || uint(userIDFromSession) != uint(userIDFromCookie) {
			return 0, false
		}

		// set the session value
		// store.Set("authenticated", true)
		// store.Save()
		return uint(userIDFromCookie), true

	}
	return 0, false
}

func IsAdmin(session *session.Session, ctx *fiber.Ctx, db *database.Database, jwtSecret string) (authenticated bool) {
	userID, ok := IsAuthenticated(session, ctx, jwtSecret)
	if !ok {
		return false
	}
	user, err := FindUserByID(db, userID)
	if err != nil || user == nil {
		return false
	}
	return models.RoleEnum(user.RoleID) == models.AdminRole
}

func FindUserByID(db *database.Database, id uint) (*models.User, error) {
	User := new(models.User)
	if response := db.Where("id = ?", id).First(&User); response.Error != nil {
		return nil, response.Error
	}
	if User.ID == 0 {
		return User, errors.New("user not found")
	}
	// Match role to user
	if User.RoleID != 0 {
		Role := new(models.Role)
		if res := db.Find(&Role, User.RoleID); res.Error != nil {
			return User, errors.New("error when retrieving the role of the user")
		}
		if Role.ID != 0 {
			User.Role = *Role
		}
	}
	return User, nil
}
