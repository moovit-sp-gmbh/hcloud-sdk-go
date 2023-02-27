package idp

import (
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

// Register registers a new user in the system
//
// entities.Register requires a captcha code which has go be generated manually
// Link: https://app.helmut.cloud/docs/register#captcha
// TODO: fix link
func (c *Client) Register(register *entities.Register) (*entities.User, *hcloud.ErrorResponse) {
	_, body, err := c.client.Post(c.getEndpoint()+"/v1/register", *register)
	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	err1 := json.Unmarshal(body, user)
	if err1 != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return user, nil
}

// RegisterVerify verifies a previous registration
//
// Registrations must be verified using the activationCode being sent to the E-Mail address provided during registration
// Link: https://app.helmut.cloud/docs/register#verify
// TODO: fix link
func (c *Client) VerifyRegistration(registerVerify *entities.RegisterVerify) (*entities.User, *hcloud.ErrorResponse) {
	_, body, err := c.client.Post(c.getEndpoint()+"/v1/register/verify", *registerVerify)
	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	err1 := json.Unmarshal(body, user)
	if err1 != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return user, nil
}
