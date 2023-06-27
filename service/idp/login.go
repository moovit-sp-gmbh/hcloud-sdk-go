package idp

import (
	"encoding/json"
	"io/ioutil"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) Login(login *entities.Login) (*entities.Token, *entities.User, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Post(c.getEndpoint()+"/v1/login", *login)
	if err != nil {
		return nil, nil, err
	}

	respBody, ioErr := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if ioErr != nil {
		return nil, nil, err
	}
	var user entities.User
	jsonError := json.Unmarshal(respBody, &user)
	if jsonError != nil {
		return nil, nil, err
	}

	token := &entities.Token{Token: resp.Header.Get("authorization")}
	return token, &user, nil
}
