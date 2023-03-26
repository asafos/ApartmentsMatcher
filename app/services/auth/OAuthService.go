package serviceAuth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fiber-boilerplate/app/constants"
	"fiber-boilerplate/app/models"
	apiServices "fiber-boilerplate/app/services/api"
	configuration "fiber-boilerplate/config"
	"fiber-boilerplate/database"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/session/v2"

	"github.com/gofiber/fiber/v2"
)

type OAuthToken struct {
	Access_token string
	Id_token     string
}

type OAuthUser struct {
	Id    string
	Email string
	Name  string
}

func GetOauthToken(url string, body io.Reader) (*OAuthToken, error) {
	req, err := http.NewRequest("POST", url, body)
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

	var OAuthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &OAuthTokenRes); err != nil {
		return nil, err
	}
	var tokenBody *OAuthToken

	if _, ok := OAuthTokenRes["id_token"].(string); ok {
		tokenBody = &OAuthToken{
			Access_token: OAuthTokenRes["access_token"].(string),
			Id_token:     OAuthTokenRes["id_token"].(string),
		}
	} else {
		tokenBody = &OAuthToken{
			Access_token: OAuthTokenRes["access_token"].(string),
		}

	}

	return tokenBody, nil
}

func GetOAuthUser(rootUrl string, authHeader interface{}) (*OAuthUser, error) {
	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	header, ok := authHeader.(string)
	if ok {
		req.Header.Set("Authorization", header)
	}

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

	userBody := &OAuthUser{
		Id:    GoogleUserRes["id"].(string),
		Email: GoogleUserRes["email"].(string),
		Name:  GoogleUserRes["name"].(string),
	}

	return userBody, nil
}

type FacebookSignedRequest struct {
	Algorithm string `json:"algorithm"`
	Expires   int64  `json:"expires"`
	IssuedAt  int64  `json:"issued_at"`
	UserID    string `json:"user_id"`
	Code      string `json:"code"`
}

func ParseFacebookSignedRequest(signedRequest string, appSecret string) (*FacebookSignedRequest, error) {
	parts := strings.SplitN(signedRequest, ".", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid signed_request format")
	}

	sig := parts[0]
	encoded := parts[1]

	sigDecoded, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		return nil, err
	}

	mac := hmac.New(sha256.New, []byte(appSecret))
	_, err = mac.Write([]byte(encoded))
	if err != nil {
		return nil, err
	}

	if !hmac.Equal(sigDecoded, mac.Sum(nil)) {
		return nil, fmt.Errorf("invalid signature")
	}

	decoded, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	var fbSignedRequest FacebookSignedRequest
	err = json.Unmarshal(decoded, &fbSignedRequest)
	if err != nil {
		return nil, err
	}

	return &fbSignedRequest, nil
}

func GetFacebookOauthToken(code string, config configuration.Config) (*OAuthToken, error) {
	const rootURl = "https://graph.facebook.com/v16.0/oauth/access_token"

	values := url.Values{}
	values.Add("code", code)
	values.Add("client_id", config.GetString("OAUTH_FACEBOOK_CLIENT_ID"))
	values.Add("client_secret", config.GetString("OAUTH_FACEBOOK_SECRET"))
	values.Add("redirect_uri", config.GetString("OAUTH_FACEBOOK_REDIRECT_URI"))

	query := values.Encode()

	return GetOauthToken(rootURl, bytes.NewBufferString(query))
}

func GetFacebookUser(access_token string) (*OAuthUser, error) {
	rootUrl := fmt.Sprintf("https://graph.facebook.com/me?fields=name,email&access_token=%s", access_token)

	return GetOAuthUser(rootUrl, nil)
}

func AddOAuthUser(ctx *fiber.Ctx, db *database.Database, oauthUser *OAuthUser) (*models.User, error) {
	User := new(models.User)
	User.Email = oauthUser.Email
	User.Name = oauthUser.Name
	User.OAuthID = oauthUser.Id
	User.RoleID = uint(models.UserRole)
	if response := apiServices.AddUser(db, User); response.Error != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "an error occurred when storing the new user"+response.Error.Error())
	}

	return User, nil
}

func TryAddOAuthUser(ctx *fiber.Ctx, db *database.Database, session *session.Session, oauthUser *OAuthUser, jwtSecret string) (user *models.User, err error) {
	var User *models.User
	User, err = apiServices.FindUserByOAuthID(db, oauthUser.Id)
	if User.ID == 0 {
		User, err = AddOAuthUser(ctx, db, oauthUser)
	}
	if err != nil {
		return nil, err
	}
	err = SerializeUser(ctx, session, User, jwtSecret)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func SerializeUser(ctx *fiber.Ctx, session *session.Session, User *models.User, jwtSecret string) error {
	store := session.Get(ctx)
	store.Set(constants.USER_ID_SESSION_KEY, User.ID)
	defer store.Save()

	// sign a JWT cookie based on the session
	claims := jwt.MapClaims{
		constants.USER_ID_SESSION_KEY: User.ID,
		"exp":                         time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return err
	}
	cookie := fiber.Cookie{
		Name:     constants.JWT_COOKIE_NAME,
		Value:    signedToken,
		MaxAge:   86400,
		HTTPOnly: true,
		SameSite: "Strict",
	}

	// set the cookies in the response
	ctx.Cookie(&cookie)
	return nil
}

func GetGoogleOauthToken(code string, config configuration.Config) (*OAuthToken, error) {
	const rootURl = "https://oauth2.googleapis.com/token"

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", config.GetString("OAUTH_GOOGLE_CLIENT_ID"))
	values.Add("client_secret", config.GetString("OAUTH_GOOGLE_SECRET"))
	values.Add("redirect_uri", config.GetString("OAUTH_GOOGLE_REDIRECT_URI"))

	query := values.Encode()

	return GetOauthToken(rootURl, bytes.NewBufferString(query))
}

func GetGoogleUser(access_token string, id_token string) (*OAuthUser, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", access_token)

	return GetOAuthUser(rootUrl, fmt.Sprintf("Bearer %s", id_token))
}
