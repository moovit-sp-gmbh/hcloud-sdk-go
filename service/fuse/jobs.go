package fuse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) PatchJobExpression(organization string, space string, jobId string, expression *entities.CronjobExpression) (*entities.Cronjob, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Put(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/jobs/%s", organization, space, jobId), expression)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Cronjob{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DeleteJob(organization string, space string, jobId string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/space/%s/jobs/%s", organization, space, jobId), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) UpdateJob(organization string, space string, jobId string, job *entities.CronjobCreation) (*entities.Cronjob, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Put(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/jobs/%s", organization, space, jobId), *job)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Cronjob{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) CreateJob(organization string, space string, job *entities.CronjobCreation) (*entities.Cronjob, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/space/%s/jobs", organization, space), *job)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.Cronjob{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) GetJobs(organization string, space string, limit int, page int) (*[]entities.Cronjob, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/space/%s/jobs?limit=%d&page=%d", organization, space, limit, page))
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

func (c *Client) GetJobById(organization string, space string, jobId string) (*entities.Cronjob, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/space/%s/jobs/%s", organization, space, jobId))
	if err != nil {
		return nil, nil, err
	}

	fuseSpace := &entities.Cronjob{}
	err1 := json.Unmarshal(body, fuseSpace)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return fuseSpace, resp, nil
}

func (c *Client) GetNextJobExecutions(organization string, space string, jobId string, amount int) (*[]int64, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/space/%s/jobs/%s/next?number=%d", organization, space, jobId, amount))
	if err != nil {
		return nil, nil, err
	}

	nextExecutions := &[]int64{}
	err1 := json.Unmarshal([]byte(body), &nextExecutions)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return nextExecutions, resp, nil
}

// GetAll requests all existing jobs form the backend
//
// INTERNAL ENDPOINT ONLY, REQUIRES MAIN 'hcloud' USER
func (c *Client) GetAllJobs(limit int, page int) (*[]entities.CronjobInternal, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/internal/v1/jobs/all?limit=%d&page=%d", limit, page))
	if err != nil {
		return nil, 0, nil, err
	}

	v := &[]entities.CronjobInternal{}
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
