package auditor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetLogs(organization string, limit int, page int) (*[]entities.AuditLog, int, *http.Response, *hcloud.ErrorResponse) {
	response, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/logs?limit=%d&page=%d", organization, limit, page))
	if err != nil {
		return nil, -1, nil, err
	}

	logs := &[]entities.AuditLog{}
	err1 := json.Unmarshal(body, logs)
	if err1 != nil {
		return nil, -1, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(response.Header.Get("total"))
	if err2 != nil {
		return nil, -1, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.parse.total.header", Message: err2.Error()}
	}

	return logs, total, response, nil
}
