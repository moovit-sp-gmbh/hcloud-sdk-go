package auditor

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) GetAuditLogs(organizationId string, limit int, page int) (*[]hcloud.AuditLog, int, *hcloud.ErrorResponse) {
	url := fmt.Sprintf("%s/v1/logs?limit=%d&page=%d", c.getEndpoint(), limit, page)

	if organizationId != "" {
		url += "&organization=" + organizationId
	}

	response, body, erro := c.client.Get(url)
	if erro != nil {
		return nil, -1, erro
	}

	logs := &[]hcloud.AuditLog{}
	err := json.Unmarshal(body, logs)
	if err != nil {
		return nil, -1, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	total, err := strconv.Atoi(response.Header.Get("total"))
	if err != nil {
		return nil, -1, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.parse.total.header", Message: err.Error()}
	}

	return logs, total, nil
}
