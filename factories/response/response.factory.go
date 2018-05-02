package responseFactory

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func Ok(data interface{}, w http.ResponseWriter) Response {
	r := Response{
		Success: true,
		Data:    data}

	if w != nil {
		json.NewEncoder(w).Encode(r)
	}

	return r
}

func Fail(code int, message string, w http.ResponseWriter) Response {
	r := Response{
		Success: false,
		Code:    code,
		Message: message}

	if w != nil {
		json.NewEncoder(w).Encode(r)
	}

	return r
}
