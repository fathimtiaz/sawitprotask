package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/domain"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helper/jwt"
	"github.com/labstack/echo/v4"
)

// Authenticate middleware adds a `Server` header to the response.
func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		swag, err := generated.GetSwagger()
		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			return err
		}

		var checkAuth bool
		switch c.Request().Method {
		case "GET":
			if swag.Paths.Find(c.Path()).Get.Security != nil {
				checkAuth = true
			}
		case "POST":
			if swag.Paths.Find(c.Path()).Post.Security != nil {
				checkAuth = true
			}
		}

		if checkAuth {
			authHdr := c.Request().Header.Get("Authorization")
			if authHdr == "" {
				c.JSON(http.StatusUnauthorized, nil)
				return errors.New("auth header doesn't exist")
			}

			prefix := "Bearer "
			if !strings.HasPrefix(authHdr, prefix) {
				c.JSON(http.StatusUnauthorized, nil)
				return errors.New("unauthorized")
			}

			token := strings.TrimPrefix(authHdr, prefix)

			claims, err := jwt.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, nil)
				return err
			}

			c.Set(domain.UserIdCtxKey, claims["user_id"])
		}

		return next(c)
	}
}
