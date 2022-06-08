package idp

import (
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) Version() (*hcloud.Version, *hcloud.ErrorResponse) {
	_, body, erro := c.HcloudClient.Get(c.getEndpoint() + "/v1/version")

	if erro != nil {
		return nil, erro
	}

	v := &hcloud.Version{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: -1, Message: err.Error()}
	}

	return v, nil
}
