package fuse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetAll(limit int, page int) (*[]entities.Cronjob, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/internal/v1/jobs/all?limit=%d&page=%d", limit, page))
	if err != nil {
		return nil, 0, nil, err
	}

	v := &[]entities.Cronjob{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(resp.Header.Get("total"))
	if err2 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.total.header.missing", Message: err2.Error()}
	}

	return v, entities.Total(total), resp, nil
}
