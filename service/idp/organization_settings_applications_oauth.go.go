package idp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moovit-sp-gmbh/hcloud-sdk-go"
	"github.com/moovit-sp-gmbh/hcloud-sdk-go/entities"
)

func (c *Client) GetOAuthApplications(organization string, limit int, page int) (*[]entities.OAuthApplication, entities.Total, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/settings/applications/oauth?limit=%d&page=%d", organization, limit, page))
	if err != nil {
		return nil, 0, nil, err
	}

	v := &[]entities.OAuthApplication{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, 0, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, 0, resp, nil
}

func (c *Client) CreateOAuthApplication(organization string, oauthApplication *entities.OAuthApplicationCreation) (*entities.OAuthApplication, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/settings/applications/oauth", organization), oauthApplication)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.OAuthApplication{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) CreateOAuthApplicationSecret(organization string, oauthApplicationId string, secret *entities.OAuthApplicationClientSecretCreation) (*entities.OAuthApplication, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Post(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/settings/applications/oauth/%s/secret", organization, oauthApplicationId), secret)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.OAuthApplication{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DeleteOAuthApplicationSecret(organization string, oauthApplicationId string, secret string) (*entities.OAuthApplication, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/settings/applications/oauth/%s/secret/%s", organization, oauthApplicationId, secret), nil)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.OAuthApplication{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) PatchOAuthApplicationSecret(organization string, oauthApplicationId string, secret string, secretName *entities.OAuthApplicationClientSecretCreation) (*entities.OAuthApplication, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Patch(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/settings/applications/oauth/%s/secret/%s", organization, oauthApplicationId, secret), nil)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.OAuthApplication{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) UpdateOAuthApplication(organization string, oauthApplicationId string, oauthApplication *entities.OAuthApplicationUpdate) (*entities.OAuthApplication, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Put(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/settings/applications/oauth/%s", organization, oauthApplicationId), oauthApplication)
	if err != nil {
		return nil, nil, err
	}

	v := &entities.OAuthApplication{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) GetOAuthApplication(organization string, oauthApplicationId string) (*entities.OAuthApplication, *http.Response, *hcloud.ErrorResponse) {
	resp, body, err := c.client.Get(c.getEndpoint() + fmt.Sprintf("/v1/org/%s/settings/applications/oauth/%s", organization, oauthApplicationId))
	if err != nil {
		return nil, nil, err
	}

	v := &entities.OAuthApplication{}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		return nil, nil, &hcloud.ErrorResponse{Code: "000.000.000", Error: "sdk.body.unmarshal", Message: err1.Error()}
	}

	return v, resp, nil
}

func (c *Client) DeleteOAuthApplication(organization string, oauthApplicationId string) (*http.Response, *hcloud.ErrorResponse) {
	resp, _, err := c.client.Delete(c.getEndpoint()+fmt.Sprintf("/v1/org/%s/settings/applications/oauth/%s", organization, oauthApplicationId), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
