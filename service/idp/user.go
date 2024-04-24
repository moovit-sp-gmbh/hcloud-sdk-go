package idp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

// GetUserEmail returns the email of the user for the active token
func (c *Client) GetUserEmail() (*entities.UserEmail, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprint("/v1/user/email"))
	if err != nil {
		return nil, nil, err
	}

	v := &entities.UserEmail{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

// GetUser returns the user for the active token
func (c *Client) GetUser() (*entities.User, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/user"))
	if err != nil {
		return nil, nil, err
	}

	v := &entities.User{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

// GetUserNats returns the allowed nats permissions for the active token
func (c *Client) GetUserNats() (*entities.NatsPermissions, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/user/nats"))
	if err != nil {
		return nil, nil, err
	}

	v := &entities.NatsPermissions{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

// DeleteUser deletes the user behind the current active token
//
// THIS OPERATION CAN NOT BE UNDONE
func (c *Client) DeleteUser() (*entities.User, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/user"), nil)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.User{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

// DeleteUserSessions deletes all active sessions of the user for the active token
//
// Deleting all sessions will forcibly logout all devices and invalidate all existing JWT Tokens for this account
func (c *Client) DeleteUserSessions() (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprint("/v1/user/sessions"), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PatchUser returns the user patched, given user
func (c *Client) PatchUser(user *entities.PatchUser) (*entities.User, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprint("/v1/user"), *user)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.User{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

// GetUserOrgs returns all organizations of the user for the active token
func (c *Client) GetUserOrganizations(limit int, page int, search entities.Search) (*[]entities.Organization, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/user/orgs/search?limit=%d&page=%d", limit, page), search)
	if err != nil {
		return nil, 0, nil, err
	}

	v := &[]entities.Organization{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(resp.Header.Get("total"))
	if err2 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.total.header.missing", Message: err2.Error()}
	}

	return v, entities.Total(total), resp, nil
}

// SearchUserOrgs returns all organizations of the user for the active token for the given search filter
func (c *Client) SearchUserOrganizations(filter *entities.SearchOrganizationFilter, limit int, page int) (*[]entities.Organization, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/user/orgs/search?limit=%d&page=%d", limit, page), *filter)
	if err != nil {
		return nil, 0, nil, err
	}

	v := &[]entities.Organization{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(resp.Header.Get("total"))
	if err2 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.total.header.missing", Message: err2.Error()}
	}

	return v, entities.Total(total), resp, nil
}
