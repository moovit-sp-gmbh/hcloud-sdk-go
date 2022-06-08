package fuse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetApps(organization string, limit int, page int) (*[]entities.FuseApp, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/apps?limit=%d&page=%d", organization, limit, page))
	if err != nil {
		return nil, 0, nil, err
	}

	v := &[]entities.FuseApp{}
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

func (c *Client) GetApp(organization string, app string) (*[]entities.FuseApp, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/apps/%s", organization, app))
	if err != nil {
		return nil, nil, err
	}

	v := &[]entities.FuseApp{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) CreateApp(organization string, app *entities.AppCreation) (*entities.FuseApp, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/apps", organization), app)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.FuseApp{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) PatchAppPermissions(organization string, app string, permission *entities.AppPermissionPatch) (*entities.FuseApp, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/apps/%s/permission", organization, app), permission)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.FuseApp{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) UpdateApp(organization string, app string, updatedApp *entities.AppCreation) (*entities.FuseApp, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Put(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/apps/%s", organization, app), updatedApp)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.FuseApp{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DeleteApp(organization string, app string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/apps/%s", organization, app), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
