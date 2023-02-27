package hcloud

import (
	"net/http"
)

type Config struct {
	Server string
	Token  string
}
type Client struct {
	config     *Config
	httpClient *http.Client
}

func New(config *Config) *Client {
	return &Client{config: config, httpClient: &http.Client{}}
}

func (c *Client) SetServer(server string) {
	c.config.Server = server
}

func (c *Client) SetToken(token string) {
	c.config.Token = token
}
