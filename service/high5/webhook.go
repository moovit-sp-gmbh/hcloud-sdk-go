package high5

import (
	"encoding/json"
	"fmt"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

func (c *Client) GetWebhooks(appId string, limit int, page int) (*[]hcloud.Webhook, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/webhook/list/%s?page=%d&limit=%d", appId, page, limit))

	if erro != nil {
		return nil, erro
	}

	e := &[]hcloud.Webhook{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}

func (c *Client) CreateWebhook(appId string, eventId string, name string, token string, target string, securityHeaders *hcloud.WebhookSecurityHeaders) (*hcloud.Webhook, *hcloud.ErrorResponse) {
	wc := &hcloud.WebhookCreation{Name: name, Token: token, AppId: appId, EventId: eventId, Target: target}
	if securityHeaders != nil {
		wc.SecurityHeaders = *securityHeaders
	}
	_, body, erro := c.client.Post(c.getEndpoint()+"/v1/webhook", wc)

	if erro != nil {
		return nil, erro
	}

	e := &hcloud.Webhook{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}

func (c *Client) DeleteWebhookById(webhookId string) *hcloud.ErrorResponse {
	_, _, erro := c.client.Delete(c.getEndpoint() + "/v1/webhook/" + webhookId)

	if erro != nil {
		return erro
	}

	return nil
}

func (c *Client) GetWebhookLogs(webhookId string, limit int, page int) (*[]hcloud.WebhookLog, *hcloud.ErrorResponse) {
	_, body, erro := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/webhook/log/list/%s?page=%d&limit=%d", webhookId, page, limit))

	if erro != nil {
		return nil, erro
	}

	e := &[]hcloud.WebhookLog{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err.Error()}
	}

	return e, nil
}
