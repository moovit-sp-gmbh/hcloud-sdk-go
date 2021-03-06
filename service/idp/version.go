package idp

import (
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) Version() (*hcloud.Version, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + "/v1/version")

	if erro != nil {
		return nil, erro
	}

	v := &hcloud.Version{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}
