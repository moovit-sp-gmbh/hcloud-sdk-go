package high5

import (
	"encoding/json"
	"fmt"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) GetStreams(eventId string, limit int, page int) (*[]hcloud.Stream, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/stream/list/%s?page=%d&limit=%d", eventId, page, limit))

	if erro != nil {
		return nil, erro
	}

	e := &[]hcloud.Stream{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}

func (c *Client) GetStreamById(streamId string) (*hcloud.Stream, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/stream/%s", streamId))

	if erro != nil {
		return nil, erro
	}

	e := &hcloud.Stream{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}

func (c *Client) DeleteStreamById(streamId string) *hcloud.ErrorResponse {
	_, _, erro := c.client.Delete(c.getEndpoint() + fmt.Sprintf("/v1/stream/%s", streamId))

	if erro != nil {
		return erro
	}

	return nil
}

func (c *Client) CreateStream(name string, eventId string) (*hcloud.Stream, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/stream/%s", eventId), hcloud.Stream{Name: name})

	if erro != nil {
		return nil, erro
	}

	e := &hcloud.Stream{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}
