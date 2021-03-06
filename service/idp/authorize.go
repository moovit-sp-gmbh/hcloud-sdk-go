package idp

import (
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) Authorize() (*hcloud.User, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + "/v1/authorize")
	if erro != nil {
		return nil, erro
	}

	user := &hcloud.User{}
	err := json.Unmarshal(body, user)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return user, nil
}
