package idp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetTeam(organization string, teamId string) (*entities.Team, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/teams/%s", organization, teamId))
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Team{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) GetTeams(organization string, limit int, page int) (*[]entities.Team, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/teams?limit=%d&page%d", organization, limit, page))
	if err != nil {
		return nil, 0, nil, err
	}

	v := &[]entities.Team{}
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

func (c *Client) CreateTeam(organization string, team *entities.TeamCreation) (*entities.Team, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/teams", organization), team)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Team{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) PatchTeam(organization string, teamId string, team *entities.TeamPatchUserOperation) (*entities.Team, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/teams/%s", organization, teamId), team)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Team{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DeleteTeam(organization string, teamId string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/teams/%s", organization, teamId), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
