package high5

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetSpaces(organization string, limit int, page int) (*[]entities.High5Space, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces?limit=%d&page=%d", organization, limit, page))

	if err != nil {
		return nil, 0, nil, err
	}

	a := &[]entities.High5Space{}
	err1 := json.Unmarshal(body, a)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(resp.Header.Get("total"))
	if err2 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.total.header.missing", Message: err2.Error()}
	}

	return a, entities.Total(total), resp, nil
}

func (c *Client) GetSpace(organization string, space string) (*entities.High5Space, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces/%s", organization, space))

	if err != nil {
		return nil, nil, err
	}

	a := &entities.High5Space{}
	err1 := json.Unmarshal(body, a)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return a, resp, nil
}

func (c *Client) CreateSpace(organization string, space *entities.SpaceCreation) (*entities.High5Space, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces", organization), space)

	if err != nil {
		return nil, nil, err
	}

	a := &entities.High5Space{}
	err1 := json.Unmarshal(body, a)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return a, resp, nil
}

func (c *Client) PatchSpacePermissions(organization string, space string, permission *entities.SpacePermissionPatch) (*entities.High5Space, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/permission", organization, space), permission)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.High5Space{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DeleteSpace(organization string, space string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s", organization, space), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
