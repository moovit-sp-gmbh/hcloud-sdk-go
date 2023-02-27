package fuse

import (
	"encoding/json"
	"fmt"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) GetAll() (*[]hcloud.Cronjob, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/internal/v1/jobs/all"))
	if erro != nil {
		return nil, erro
	}

	v := &[]hcloud.Cronjob{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}
