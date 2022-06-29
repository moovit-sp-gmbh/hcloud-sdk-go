package idp

import (
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) Register(name string, company string, email string, password string) (*hcloud.User, *hcloud.ErrorResponse) {
	register := hcloud.Register{Name: name, Company: company, Email: email, Password: password}

	_, body, erro := c.client.Post(c.getEndpoint()+"/v1/registration", register)
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
