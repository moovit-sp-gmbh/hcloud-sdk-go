package idp

import (
	"encoding/json"
	"fmt"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) PatchUser(user hcloud.PatchUser) (*hcloud.User, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/user"), user)
	if erro != nil {
		return nil, erro
	}

	v := &hcloud.User{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}
