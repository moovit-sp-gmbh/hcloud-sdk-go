package hcloud

import (
	"net/http"
)

type Config struct {
	Api   string
	Token string
}
type Client struct {
	config     *Config
	httpClient *http.Client
}

func New(config *Config) *Client {
	return &Client{config: config, httpClient: &http.Client{}}
}

func (c *Client) SetApi(api string) {
	c.config.Api = api
}

func (c *Client) SetToken(token string) {
	c.config.Token = token
}
