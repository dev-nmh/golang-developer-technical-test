package config

import (
	"github/golang-developer-technical-test/internal/model"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type JwtGenerator struct {
	config *viper.Viper
}

func NewJwtGenerator(config *viper.Viper) *JwtGenerator {
	return &JwtGenerator{
		config: config,
	}
}
func (c JwtGenerator) CreateAccessTokenUser(user *model.UserBaseClaims) (accessToken string, err error) {
	expHour := c.config.GetInt("app.exp_token_inhour")
	secret := c.config.GetString("app.app_key")
	log.Println(expHour)
	exp := time.Now().Add(time.Hour * time.Duration(expHour)).Unix()
	claims := &model.JwtUserClaims{
		UserBaseClaims: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}
func (c JwtGenerator) CreateAccessTokenAdmin(user *model.BaseClaims) (accessToken string, err error) {
	expHour := c.config.GetInt("app.exp_token_inhour")
	secret := c.config.GetString("app.app_key")
	log.Println(expHour)
	exp := time.Now().Add(time.Hour * time.Duration(expHour)).Unix()
	claims := &model.JwtBaseClaims{
		BaseClaims: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}
func (c JwtGenerator) CreateRefreshToken(user_id string) (refreshToken string, err error) {
	expHour := c.config.GetInt("app.exp_refresh_token_inday")
	secret := c.config.GetString("app.app_key")
	claimsRefresh := &model.JwtCustomRefreshClaims{
		ID: user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(expHour)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}
