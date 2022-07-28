package high5

import (
	"encoding/json"
	"fmt"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) GetEvents(appId string, limit int, page int) (*[]hcloud.Event, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/event/list/%s?page=%d&limit=%d", appId, page, limit))

	if erro != nil {
		return nil, erro
	}

	e := &[]hcloud.Event{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}

func (c *Client) GetEventById(eventId string) (*hcloud.Event, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/event/%s", eventId))

	if erro != nil {
		return nil, erro
	}

	e := &hcloud.Event{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}

func (c *Client) CreateEvent(name string, appId string) (*hcloud.Event, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/event/%s", appId), hcloud.Event{Name: name})

	if erro != nil {
		return nil, erro
	}

	e := &hcloud.Event{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}

func (c *Client) DeleteEventById(eventId string) *hcloud.ErrorResponse {
	_, _, erro := c.client.Delete(c.getEndpoint() + fmt.Sprintf("/v1/event/%s", eventId))

	if erro != nil {
		return erro
	}

	return nil
}
