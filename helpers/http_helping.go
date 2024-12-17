package helpers

import (
	"encoding/json"
	"net/http"
)

func GetJsonFromBody(r *http.Request) interface{} {
	body := r.Body
	defer body.Close()
	var bodyJson interface{}
	err := json.NewDecoder(body).Decode(&bodyJson)
	if err != nil {
		return map[string]interface{}{"error": "Failed to decode body"}
	}
	return bodyJson
}
