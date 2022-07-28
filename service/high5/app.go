package high5

import (
	"encoding/json"
	"fmt"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) GetApps(limit int, page int) (*[]hcloud.App, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/app?page=%d&limit=%d", page, limit))

	if erro != nil {
		return nil, erro
	}

	a := &[]hcloud.App{}
	err := json.Unmarshal(body, a)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return a, nil
}

func (c *Client) GetAppById(appId string) (*hcloud.App, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/app/%s", appId))

	if erro != nil {
		return nil, erro
	}

	a := &hcloud.App{}
	err := json.Unmarshal(body, a)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return a, nil
}

func (c *Client) CreateApp(name string) (*hcloud.App, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Post(c.getEndpoint()+"/v1/app", hcloud.App{Name: name})

	if erro != nil {
		return nil, erro
	}

	fmt.Print(string(body))
	a := &hcloud.App{}
	err := json.Unmarshal(body, a)

	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return a, nil
}

func (c *Client) DeleteAppById(appId string) *hcloud.ErrorResponse {
	_, _, erro := c.client.Delete(c.getEndpoint() + fmt.Sprintf("/v1/app/%s", appId))

	if erro != nil {
		return erro
	}

	return nil
}
