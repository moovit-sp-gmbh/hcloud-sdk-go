package hcloud

import (
	"encoding/json"
	"fmt"
)

// ErrorReponse holds a code and a message parsed from an error resposne from the helmut.cloud platform
type ErrorResponse struct {
	Code    string `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func (e *ErrorResponse) ToString() string {
	return fmt.Sprintf("code: %s, error: %s, message: %s", e.Code, e.Error, e.Message)
}

func parseError(body []byte) *ErrorResponse {
	errResponse := &ErrorResponse{}
	err := json.Unmarshal(body, errResponse)
	if err != nil {
		return &ErrorResponse{Code: "000.000.000", Error: "sdk.parse.error", Message: err.Error()}
	}
	return errResponse
}
