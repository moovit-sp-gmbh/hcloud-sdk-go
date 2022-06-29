package idp

import (
	"encoding/json"
	"fmt"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) ListOrganizations(page int, limit int) (*[]hcloud.Organization, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/organization?page=%d&limit=%d", page, limit))
	if erro != nil {
		return nil, erro
	}

	v := &[]hcloud.Organization{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}

func (c *Client) CreateOrganization(name string, company string) (*hcloud.Organization, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Post(c.getEndpoint()+"/v1/organization", hcloud.Organization{Name: name, Company: company})
	if erro != nil {
		return nil, erro
	}

	v := &hcloud.Organization{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}

func (c *Client) GetOrganizationById(id string) (*hcloud.Organization, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + "/v1/organization/" + id)
	if erro != nil {
		return nil, erro
	}

	v := &hcloud.Organization{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}

func (c *Client) DeleteOrganizationById(id string) *hcloud.ErrorResponse {
	_, _, erro := c.client.Delete(c.getEndpoint() + "/v1/organization/" + id)
	if erro != nil {
		return erro
	}

	return nil
}

func (c *Client) ListOrganizationMembersById(id string, page int, limit int) (*[]hcloud.OrganizationMember, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/organization/"+id+"/member?page=%d&limit=%d", page, limit))
	if erro != nil {
		return nil, erro
	}

	fmt.Print(string(body))

	v := &[]hcloud.OrganizationMember{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}

func (c *Client) AddOrganizationMemberById(id string, userId string) (*hcloud.OrganizationMember, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Post(c.getEndpoint()+"/v1/organization/"+id+"/member/"+userId, hcloud.OrganizationMember{Roles: []string{"admin"}})
	if erro != nil {
		return nil, erro
	}

	v := &hcloud.OrganizationMember{}
	err := json.Unmarshal(body, v)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return v, nil
}

func (c *Client) DeleteOrganizationMemberById(id string) *hcloud.ErrorResponse {
	_, _, erro := c.client.Delete(c.getEndpoint() + "/v1/organization/member/" + id)
	if erro != nil {
		return erro
	}

	return nil
}
