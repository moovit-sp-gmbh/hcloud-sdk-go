package idp

import (
	"encoding/json"
	"fmt"
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
	user := &entities.User{}
	jsonError := json.Unmarshal(respBody, user)
	if jsonError != nil {
		return nil, nil, err
	}
	fmt.Println(resp)
	fmt.Println(respBody)
	fmt.Println(user)
	if user == nil {
		err := &hcloud.ErrorResponse{
			Code: "001.001.000",
			Error: "unauthorized",
			Message: "You're not authorized",
		}
		return nil, nil, err
	}

	token := &entities.Token{Token: resp.Header.Get("authorization")}
	return token, user, nil
}
