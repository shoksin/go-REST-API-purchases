package utils

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
)

func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

func Respond(c echo.Context, statusCode int, message map[string]interface{}) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Status = statusCode
	if err := json.NewEncoder(c.Response()).Encode(message); err != nil {
		logging.GetLogger().Fatal("API respond isn't encoded")
		return err
	}
	return nil
}
