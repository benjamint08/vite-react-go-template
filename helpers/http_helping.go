package helpers

import (
	"encoding/json"
	"net/http"
)

func GetJsonFromBody(r *http.Request) map[string]interface{} {
	body := r.Body
	defer body.Close()
	var bodyJson map[string]interface{}
	err := json.NewDecoder(body).Decode(&bodyJson)
	if err != nil {
		return map[string]interface{}{"error": "Failed to decode body"}
	}
	return bodyJson
}
