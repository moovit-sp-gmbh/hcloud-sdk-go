package high5

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetStreams(organization string, space string, event string, limit int, page int) (*[]entities.Stream, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces/%s/events/%s/streams?limit=%d&page=%d", organization, space, event, limit, page))

	if err != nil {
		return nil, 0, nil, err
	}

	e := &[]entities.Stream{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(resp.Header.Get("total"))
	if err2 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.total.header.missing", Message: err2.Error()}
	}

	return e, entities.Total(total), nil, nil
}

func (c *Client) GetStream(organization string, space string, event string, streamId string) (*entities.Stream, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces/%s/events/%s/streams/%s", organization, space, event, streamId))
	if err != nil {
		return nil, nil, err
	}

	e := &entities.Stream{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}

func (c *Client) DeleteStream(organization string, space string, event string, streamId string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/events/%s/stream/%s", organization, space, event, streamId), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateStream(organization string, space string, event string, stream *entities.StreamCreation) (*entities.Stream, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/events/%s/streams", organization, space, event), stream)
	if err != nil {
		return nil, nil, err
	}

	e := &entities.Stream{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}

func (c *Client) ChangeStreamOrder(organization string, space string, event string, streamOrder *[]entities.StreamOrder) (*[]entities.Stream, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/events/%s/streams", organization, space, event), streamOrder)
	if err != nil {
		return nil, nil, err
	}

	e := &[]entities.Stream{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}
