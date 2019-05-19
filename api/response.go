package api

import (
	"encoding/json"
	"net/http"
)

type (
	Response struct {
		BaseResponse
		Data interface{} `json:"data"`
	}
	BaseResponse struct {
		Errors []string `json:"errors,omitempty"`
	}
)

var (
	MessageUnauthorized = "Please login first or try re-login."
)

func RespondError(w http.ResponseWriter, message string, status int) {
	resp := Response{
		Data: nil,
		BaseResponse: BaseResponse{
			Errors: []string{
				message,
			},
		},
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		return
	}
}
