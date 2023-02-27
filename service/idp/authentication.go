package idp

import "github.com/moovit-sp-gmbh/hcloud-sdk-go"

func (c *Client) Login(email string, password string) (*hcloud.Token, *hcloud.ErrorResponse) {
	login := hcloud.Login{Email: email, Password: password}

	resp, _, erro := c.client.Post(c.getEndpoint()+"/v1/login", login)
	if erro != nil {
		return nil, erro
	}

	token := &hcloud.Token{Token: resp.Header.Get("authorization")}
	return token, nil
}
