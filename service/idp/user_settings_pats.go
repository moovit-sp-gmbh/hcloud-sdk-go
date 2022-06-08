package idp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetPats(limit int, page int) (*[]entities.UserPat, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/user/settings/pats?limit=%d&page=%d", limit, page))
	if err != nil {
		return nil, nil, err
	}

	v := &[]entities.UserPat{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) CreatePat(pat *entities.UserPatCreation) (*entities.UserPat, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/user/settings/pats"), pat)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.UserPat{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DeletePats() (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/user/settings/pats"), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeletePat(patId string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/user/settings/pats/%s", patId), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) UpdatePat(patId string, pat *entities.UserPatPatch) (*entities.UserPat, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/user/settings/pats/%s", patId), pat)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.UserPat{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}
