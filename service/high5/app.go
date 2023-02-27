package high5

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetApps(organization string, limit int, page int) (*[]entities.App, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/apps?limit=%d&page=%d", organization, limit, page))

	if err != nil {
		return nil, 0, nil, err
	}

	a := &[]entities.App{}
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

func (c *Client) GetAppByName(organization string, app string) (*entities.App, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/apps/%s", organization, app))

	if err != nil {
		return nil, nil, err
	}

	a := &entities.App{}
	err1 := json.Unmarshal(body, a)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return a, resp, nil
}

func (c *Client) CreateApp(organization string, app *entities.AppCreation) (*entities.App, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/apps", organization), app)

	if err != nil {
		return nil, nil, err
	}

	a := &entities.App{}
	err1 := json.Unmarshal(body, a)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return a, resp, nil
}

func (c *Client) DeleteAppByName(organization string, app string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/apps/%s", organization, app))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
