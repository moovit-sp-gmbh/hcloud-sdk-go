package high5

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetEvents(organization string, space string, limit int, page int) (*[]entities.Event, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces/%s/events?limit=%d&page=%d", organization, space, limit, page))

	if err != nil {
		return nil, 0, nil, err
	}

	e := &[]entities.Event{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(resp.Header.Get("total"))
	if err2 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.total.header.missing", Message: err2.Error()}
	}

	return e, entities.Total(total), resp, nil
}

func (c *Client) GetEvent(organization string, space string, event string) (*entities.Event, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces/%s/events/%s", organization, space, event))
	if err != nil {
		return nil, nil, err
	}

	e := &entities.Event{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}

func (c *Client) CreateEvent(organization string, space string, event *entities.EventCreation) (*entities.Event, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/events", organization, space), event)
	if err != nil {
		return nil, nil, err
	}

	e := &entities.Event{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}

func (c *Client) DeleteEvent(organization string, space string, event string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/events/%s", organization, space, event), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
