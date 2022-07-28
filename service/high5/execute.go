package high5

import (
	b64 "encoding/base64"
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) ExecuteStreamById(streamId string, target string, data []byte, timeout int, waitForResult bool) (*map[string]interface{}, *hcloud.ErrorResponse) {
	b64data := b64.StdEncoding.EncodeToString(data)
	_, body, erro := c.client.Post(c.getEndpoint()+"/v1/execute/stream/id", hcloud.StreamExecutionRequest{StreamId: streamId, Target: target, Timeout: timeout, WaitForResult: waitForResult, Data: b64data})

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
