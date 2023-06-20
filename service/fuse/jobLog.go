package fuse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (client *Client) GetJobLogs(orgName string, space string, cronjobId string, limit int, page int) (*[]entities.CronjobLog, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := client.client.Get(client.getEndpoint() + fmt.Sprintf("/v1/org/%s/space/%s/jobs/%s/logs?limit=%d&page=%d", orgName, space, cronjobId, limit, page))
	if err != nil {
		return nil, 0, nil, err
	}

	cronjobLogs := &[]entities.CronjobLog{}
	err1 := json.Unmarshal(body, cronjobLogs)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(resp.Header.Get("total"))
	if err2 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.total.header.missing", Message: err2.Error()}
	}

	return cronjobLogs, entities.Total(total), resp, nil
}

func (client *Client) GetJobLogById(organization string, space string, cronJobId string, cronJobLogId string) (*entities.CronjobLog, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := client.client.Get(client.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces/%s/jobs/%s/logs/%s", organization, space, cronJobId, cronJobLogId))
	if err != nil {
		return nil, nil, err
	}

	cronjobLog := &entities.CronjobLog{}
	err1 := json.Unmarshal(body, cronjobLog)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return cronjobLog, resp, nil
}
