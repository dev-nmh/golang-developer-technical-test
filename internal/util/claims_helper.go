package util

import (
	"fmt"
	"github/golang-developer-technical-test/internal/constant"
	"strconv"

	"braces.dev/errtrace"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ClaimsUtil struct {
	claims jwt.MapClaims
}

func NewClaimUtil(e echo.Context) (*ClaimsUtil, error) {
	Auth := e.Get("Authorization")
	if Auth == nil {
		return nil, errtrace.Wrap(echo.ErrUnauthorized)
	}
	claims, ok := Auth.(jwt.MapClaims)
	if !ok {
		return nil, errtrace.Wrap(echo.ErrUnauthorized)
	}
	return &ClaimsUtil{claims: claims}, nil

}
func (c ClaimsUtil) GetRole() (int, error) {

	if id, ok := c.claims["RoleId"].(string); ok {
		if idInt, err := strconv.Atoi(id); err != nil {
			return 0, errtrace.Wrap(err)
		} else {
			return idInt, nil

		}
	} else {
		return 0, errtrace.Wrap(fmt.Errorf("format is not supported"))
	}
}

func (c ClaimsUtil) GetId() (uuid.UUID, error) {
	if id, ok := c.claims["ID"].(string); ok {
		if idUUID, err := uuid.Parse(id); err != nil {
			return uuid.Nil, errtrace.Wrap(err)
		} else {
			return idUUID, nil

		}
	} else {
		return uuid.Nil, errtrace.Wrap(fmt.Errorf("format is not supported"))
	}
}

func (c ClaimsUtil) GetUserId() (uuid.UUID, error) {
	if roleId, err := c.GetRole(); err != nil {
		return uuid.Nil, errtrace.Wrap(err)
	} else if roleId != constant.USER_ROLES_USER_INT {
		return uuid.Nil, errtrace.Wrap(fmt.Errorf("Not Allow To Access"))
	}

	if id, ok := c.claims["UserID"].(string); ok {
		if idUUID, err := uuid.Parse(id); err != nil {
			return uuid.Nil, errtrace.Wrap(err)
		} else {
			return idUUID, nil

		}
	} else {
		return uuid.Nil, errtrace.Wrap(fmt.Errorf("format is not supported"))
	}
}
