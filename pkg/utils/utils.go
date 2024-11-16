package utils

import (
	"encoding/json"
	"github.com/shoksin/go-contacts-REST-API-/pkg/logging"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logging.GetLogger().Error("Utils respond error: ", err)
	}
}