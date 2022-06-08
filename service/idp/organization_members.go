package idp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetOrganizationMembers(organization string, limit int, page int) (*[]entities.OrganizationMember, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/members?limit=%d&page=%d", organization, limit, page))
	if err != nil {
		return nil, 0, nil, err
	}

	v := &[]entities.OrganizationMember{}
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

func (c *Client) InviteOrganizationMember(organization string, member *entities.OrganizationMemberCreation) (*entities.OrganizationMember, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/members/invitations", organization), member)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.OrganizationMember{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) PatchOrganizationMemberPermission(organization string, userId string, permission *entities.PatchOrganizationMemberPermission) (*entities.OrganizationMember, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/members/%s/permissions", organization, userId), permission)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.OrganizationMember{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DeleteOrganizationMember(organization string, userId string) (*http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/members/%s", organization, userId), nil)
	if err != nil {
		return nil, err
	}

	v := &entities.OrganizationMember{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return resp, nil
}
