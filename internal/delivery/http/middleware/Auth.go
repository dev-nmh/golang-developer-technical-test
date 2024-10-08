package middleware

import (
	"github/golang-developer-technical-test/internal/constant"
	"github/golang-developer-technical-test/internal/model"
	"github/golang-developer-technical-test/internal/util"
	"net/http"
	"strconv"
	"strings"
	"time"

	"braces.dev/errtrace"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Middleware struct {
	config *viper.Viper
}

func NewMiddleware(cfg *viper.Viper) *Middleware {
	return &Middleware{config: cfg}
}

func (m Middleware) AuthApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := m.config.GetString("app.api_key")
		apiKey := c.Request().Header.Get("X-API-KEY")
		if apiKey != key {
			var response model.JSONResponse
			response.Message = http.StatusText(http.StatusUnauthorized)
			response.StatusCode = http.StatusUnauthorized
			return errtrace.Wrap(c.JSON(http.StatusUnauthorized, response))
		} else {
			return errtrace.Wrap(next(c))

		}
	}
}

func (m Middleware) BaseAuth(e echo.Context) error {

	if authorization := e.Request().Header.Get("Authorization"); authorization != "" {
		authorizationToken := strings.Split(authorization, " ")
		if len(authorizationToken) != 2 {
			return errtrace.Wrap(echo.ErrUnauthorized)
		}

		if authorizationToken[0] != "Bearer" {
			return errtrace.Wrap(echo.ErrUnauthorized)
		}

		e.Set("RawToken", authorizationToken[1])

		token, err := jwt.Parse(authorizationToken[1], func(token *jwt.Token) (interface{}, error) {
			_, _ = token.Method.(*jwt.SigningMethodHMAC)

			return []byte(m.config.GetString("app.app_key")), nil
		})

		if err != nil {
			return errtrace.Wrap(echo.ErrUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			return errtrace.Wrap(echo.ErrUnauthorized)
		}
		times := strconv.FormatFloat(claims["exp"].(float64), 'f', 0, 64)
		timeint64, _ := strconv.ParseInt(times, 10, 64)
		expired := time.Unix(timeint64, 0).In(time.Local)

		if time.Now().After(expired) {
			return errtrace.Wrap(echo.ErrUnauthorized)
		}

		e.Set("Authorization", claims)
		return errtrace.Wrap(nil)
	}
	return errtrace.Wrap(echo.ErrUnauthorized)
}
func (m Middleware) AuthBaseAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		err := m.BaseAuth(e)
		if err != nil {
			return errtrace.Wrap(err)
		} else {
			return errtrace.Wrap(next(e))

		}
	}

}

func (m Middleware) AuthAdminJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		err := m.BaseAuth(e)
		if err != nil {
			response := new(model.JSONResponse)
			response.StatusCode = http.StatusUnauthorized
			if httpError, ok := err.(*echo.HTTPError); ok {
				response.StatusCode = httpError.Code
				response.Message = httpError.Message.(string)
				return errtrace.Wrap(e.JSON(response.StatusCode, response))
			}
			return errtrace.Wrap(e.JSON(response.StatusCode, response))
		} else {
			claims := e.Get("Authorization").(jwt.MapClaims)
			if role, ok := claims["RoleId"].(string); !ok || role == constant.USER_ROLES_USER {
				response := util.CreateResponse(http.StatusForbidden, "Forbidden: Admin access required", nil)
				return errtrace.Wrap(e.JSON(response.StatusCode, response))
			} else {
				return errtrace.Wrap(next(e))

			}
		}

	}
}

func (m Middleware) AuthUserJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		err := m.BaseAuth(e)
		if err != nil {
			response := new(model.JSONResponse)
			response.StatusCode = http.StatusUnauthorized
			response.Message = http.StatusText(http.StatusUnauthorized)
			if httpError, ok := err.(*echo.HTTPError); ok {
				response.StatusCode = httpError.Code
				response.Message = httpError.Message.(string)
				return errtrace.Wrap(e.JSON(response.StatusCode, response))
			}
			return errtrace.Wrap(e.JSON(response.StatusCode, response))
		} else {
			claims := e.Get("Authorization").(jwt.MapClaims)

			if role, ok := claims["RoleId"].(string); !ok || role != constant.USER_ROLES_USER {
				response := util.CreateResponse(http.StatusForbidden, "Forbidden: User access required", nil)
				return errtrace.Wrap(e.JSON(response.StatusCode, response))
			} else {
				return errtrace.Wrap(next(e))
			}
		}

	}
}
