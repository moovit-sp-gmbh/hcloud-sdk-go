package high5

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetWebhooks(organization string, space string, limit int, page int) (*[]entities.Webhook, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces/%s/webhooks?limit=%d&page=%d", organization, space, limit, page))
	if err != nil {
		return nil, 0, nil, err
	}

	e := &[]entities.Webhook{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	total, err2 := strconv.Atoi(resp.Header.Get("total"))
	if err2 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.total.header.missing", Message: err2.Error()}
	}

	return e, entities.Total(total), resp, nil
}

func (c *Client) CreateWebhook(organization string, space string, webhook *entities.WebhookCreation) (*entities.Webhook, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/webhooks", organization, space), *webhook)

	if err != nil {
		return nil, resp, err
	}

	e := &entities.Webhook{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}

func (c *Client) PatchWebhook(organization string, space string, webhookId string, webhook *entities.PatchWebhook) (*entities.Webhook, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/webhooks/%s", organization, space, webhookId), *webhook)

	if err != nil {
		return nil, resp, err
	}

	e := &entities.Webhook{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}

func (c *Client) DeleteWebhook(organization string, space string, webhookId string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/webhooks/%s", organization, space, webhookId), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RegenerateWebhookUrl(organization string, space string, webhookId string) (*entities.Webhook, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/spaces/%s/webhooks/%s/logs", organization, space, webhookId), "")
	if err != nil {
		return nil, nil, err
	}

	e := &entities.Webhook{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}

func (c *Client) GetWebhookLogs(organization string, space string, webhookId string, limit int, page int) (*[]entities.WebhookLog, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/spaces/%s/webhooks/%s/logs?limit=%d&page=%d", organization, space, webhookId, limit, page))
	if err != nil {
		return nil, nil, err
	}

	e := &[]entities.WebhookLog{}
	err1 := json.Unmarshal(body, e)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return e, resp, nil
}
