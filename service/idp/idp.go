package idp

import (
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
)

// Client is a representative of the idp Client
type Client struct {
	HcloudClient *hcloud.Client
}

var httpClient *http.Client

func NewFromConfig(config *hcloud.Config) *Client {
	return &Client{HcloudClient: &hcloud.Client{Config: config, Token: config.Token, HttpClient: &http.Client{}}}
}

func (c *Client) SetToken(token string) {
	c.HcloudClient.Token = token
}
