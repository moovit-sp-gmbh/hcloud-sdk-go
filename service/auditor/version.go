package auditor

import (
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) Version() (*entities.Version, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + "/v1/version")

	if erro != nil {
		return nil, erro
	}

	v := &entities.Version{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}
