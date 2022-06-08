package idp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) EnableTOTP() (*entities.TOTP, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/user/settings/security/2fa/totp"), nil)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.TOTP{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) ActivateTOTP(token *entities.TOTPToken) (*entities.TOTPActivated, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/user/settings/security/2fa/totp"), token)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.TOTPActivated{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DisableTOTP(token *entities.TOTPToken) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/user/settings/security/2fa/totp"), token)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
