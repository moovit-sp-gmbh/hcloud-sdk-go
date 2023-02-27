package high5

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetStreams(organization string, app string, event string, limit int, page int) (*[]entities.Stream, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/apps/%s/events/%s/streams?limit=%d&page=%d", organization, app, event, limit, page))

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

func (c *Client) GetStreamById(organization string, app string, event string, streamId string) (*entities.Stream, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/apps/%s/events/%s/streams/%s", organization, app, event, streamId))
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

func (c *Client) DeleteStreamById(organization string, app string, event string, streamId string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/apps/%s/events/%s/stream/%s", organization, app, event, streamId))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateStream(organization string, app string, event string, stream *entities.StreamCreation) (*entities.Stream, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/apps/%s/events/%s/streams", organization, app, event), stream)
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

func (c *Client) ChangeStreamOrder(organization string, app string, event string, streamOrder *[]entities.StreamOrder) (*[]entities.Stream, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/apps/%s/events/%s/streams", organization, app, event), streamOrder)
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
