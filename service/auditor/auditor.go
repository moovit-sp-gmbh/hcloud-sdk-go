package auditor

import "github.com/moovit-sp-gmbh/hcloud-sdk-go"

type Client struct {
	client *hcloud.Client
}

func New(client *hcloud.Client) *Client {
	return &Client{client: client}
}

func (c *Client) getEndpoint() string {
	return c.client.GetEndpoint("/api/auditor")
}
