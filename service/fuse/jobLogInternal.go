package fuse

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
	"github.com/sirupsen/logrus"
)

func (client *Client) CreateCronjobLog(orgName string, space string, cronjobId string, log entities.CronjobLogCreation) (*entities.CronjobLog, *http.Response, *hcloud.ErrorResponse) {
	logrus.Infof(client.getEndpoint() + fmt.Sprintf("/internal/v1/org/%s/spaces/%s/jobs/%s/logs", orgName, space, cronjobId))
	resp, body, err := client.client.Post(client.getEndpoint() + fmt.Sprintf("/internal/v1/org/%s/spaces/%s/jobs/%s/logs", orgName, space, cronjobId), log)
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
