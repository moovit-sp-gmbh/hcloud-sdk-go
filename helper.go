package hcloud

func (c *Client) GetEndpoint(endpoint string) string {
	return c.config.Server + endpoint
}
