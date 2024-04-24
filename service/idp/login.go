package idp

import (
	"encoding/json"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) Login(login *entities.Login) (*entities.Token, *entities.User, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+"/v1/login", *login)
	if err != nil {
		return nil, nil, err
	}

	user := &entities.User{}
	jsonError := json.Unmarshal(body, user)
	if jsonError != nil {
		return nil, nil, err
	}

	token := &entities.Token{Token: resp.Header.Get("authorization")}
	return token, user, nil
}
