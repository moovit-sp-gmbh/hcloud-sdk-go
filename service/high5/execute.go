package high5

import (
	b64 "encoding/base64"
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) ExecuteStreamById(streamId string, target string, dataType string, data []byte, timeout int, waitForResult bool) (*map[string]interface{}, *hcloud.ErrorResponse) {
	b64data := b64.StdEncoding.EncodeToString(data)

	_, body, erro := c.client.Post(c.getEndpoint()+"/v1/execute/stream/id/"+streamId, hcloud.StreamExecutionRequest{Target: target, Timeout: timeout, WaitForResult: waitForResult, Payload: hcloud.StreamExecutionPayload{Type: dataType, Data: b64data}})
	if erro != nil {
		return nil, erro
	}

	var t map[string]interface{}
	err := json.Unmarshal(body, &t)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return &t, nil
}

func (c *Client) ExecuteEventByName(appId string, eventName string, target string, dataType string, data []byte, timeout int, waitForResult bool) (*[]interface{}, *hcloud.ErrorResponse) {
	b64data := b64.StdEncoding.EncodeToString(data)
	_, body, erro := c.client.Post(c.getEndpoint()+"/v1/execute/event/name/"+appId, hcloud.EventExecutionRequest{EventName: eventName, Target: target, Timeout: timeout, WaitForResult: waitForResult, Payload: hcloud.StreamExecutionPayload{Type: dataType, Data: b64data}})
	if erro != nil {
		return nil, erro
	}

	var t []interface{}
	err := json.Unmarshal(body, &t)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return &t, nil
}
