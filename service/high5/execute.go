package high5

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) ExecuteStreamById(organization string, space string, streamId string, streamExecutionRequest *entities.StreamExecutionRequest) (*entities.StreamResult, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/execute/stream/id/%s", organization, space, streamId), streamExecutionRequest)
	if err != nil {
		return nil, nil, err
	}

	var t = &entities.StreamResult{}
	err1 := json.Unmarshal(body, &t)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return t, resp, nil
}

func (c *Client) ExecuteEventByName(organization string, space string, event string, streamExecutionRequest *entities.StreamExecutionRequest) (*entities.StreamResult, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/execute/event/name/%s", organization, space, event), streamExecutionRequest)
	if err != nil {
		return nil, nil, err
	}

	var t = &entities.StreamResult{}
	err1 := json.Unmarshal(body, t)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return t, resp, nil
}
