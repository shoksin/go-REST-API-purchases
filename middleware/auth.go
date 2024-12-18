package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-REST-API-purchases/internal/models"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"net/http"
	"os"
	"strings"
)

func GetToken(c echo.Context) (*models.Token, error) {
	tokenHeader := c.Request().Header.Get("Authorization")

	if tokenHeader == "" {
		logging.GetLogger().Error("No Authorization header found")
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "No Authorization header found")
	}

	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		logging.GetLogger().Error("Invalid Authorization header")
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header")
	}

	tokenPart := splitted[1]
	tk := &models.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		logging.GetLogger().Error("Error while parsing token ", err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error while parsing token")
	}

	if !token.Valid {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	return tk, nil
}

var JWTAuth = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		nonAuth := []string{"/register", "/login"}
		requestPath := c.Request().URL.Path

		for _, value := range nonAuth {
			if value == requestPath {
				if err := next(c); err != nil {
					c.Error(err)
				}
				return nil
			}
		}

		tk, err := GetToken(c)
		if err != nil {
			logging.GetLogger().Error("JWTAuth error")
			return err
		}

		ctx := context.WithValue(c.Request().Context(), "user_id", tk.UserId)
		ctx = context.WithValue(c.Request().Context(), "role", tk.Role)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
