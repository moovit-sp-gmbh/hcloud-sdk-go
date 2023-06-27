package idp

import (
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) Login(login *entities.Login) (*entities.Token, *http.Response,*hcloud.ErrorResponse) {
	resp, _, err := c.client.Post(c.getEndpoint()+"/v1/login", *login)
	if err != nil {
		return nil, resp, err
	}

	token := &entities.Token{Token: resp.Header.Get("authorization")}
	return token, resp, nil
}
