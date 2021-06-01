package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// BaseResponse is basic JSON response form
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse is helper to send as JSON response
func SendResponse(w http.ResponseWriter, code int, payload interface{}) (interface{}, int) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Println(err)
	}
	return payload, code
}
