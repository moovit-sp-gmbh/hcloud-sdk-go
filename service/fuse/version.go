package fuse

import (
	"encoding/json"
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) Version() (*entities.Version, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + "/v1/version")
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Version{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}
