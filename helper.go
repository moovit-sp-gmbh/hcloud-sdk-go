package hcloud

import "encoding/json"

// ErrorReponse holds a code and a message parsed from an error resposne from the helmut.cloud platform
type ErrorResponse struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func parseError(body []byte) *ErrorResponse {
	errResponse := &ErrorResponse{}
	err := json.Unmarshal(body, errResponse)
	if err != nil {
		return &ErrorResponse{Code: -1, Message: err.Error()}
	}
	return errResponse
}

func (c *Client) GetEndpoint(endpoint string) string {
	return c.config.Api + endpoint
}
