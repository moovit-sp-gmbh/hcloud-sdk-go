package idp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

// CreateOrganization returns the new organization
func (c *Client) CreateOrganization(organization *entities.Organization) (*entities.Organization, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+"/v1/org", organization)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Organization{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

// GetOrganizationByName returns a given organization if found
func (c *Client) GetOrganization(name string) (*entities.Organization, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + "/v1/org/" + name)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Organization{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

// DeleteOrganizationByName deletes an organization by name
func (c *Client) DeleteOrganization(name string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s", name), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
