package hcloud

import (
	_ "embed"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

//go:embed .version
var version string

func (c *Client) Get(url string) (*http.Response, []byte, *ErrorResponse) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	setGlobalHeaders(req)

	if c.config.Token != "" {
		req.Header.Set("Authorization", c.config.Token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	if resp.StatusCode >= 400 {
		return nil, nil, parseError(body)
	}

	return resp, body, nil
}

func (c *Client) Post(url string, payload interface{}) (*http.Response, []byte, *ErrorResponse) {
	var b []byte
	var err error
	if payload != nil {
		b, err = json.Marshal(payload)
		if err != nil {
			return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
		}
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(b)))
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	setGlobalHeaders(req)

	req.Header.Set("Content-Type", "application/json")
	if c.config.Token != "" {
		req.Header.Set("Authorization", c.config.Token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	if resp.StatusCode >= 400 {
		return nil, nil, parseError(body)
	}

	return resp, body, nil
}

func (c *Client) Delete(url string) (*http.Response, []byte, *ErrorResponse) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	setGlobalHeaders(req)

	if c.config.Token != "" {
		req.Header.Set("Authorization", c.config.Token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, &ErrorResponse{Code: "000.000.000", Error: "sdk.http", Message: err.Error()}
	}

	if resp.StatusCode >= 400 {
		return nil, nil, parseError(body)
	}

	return resp, body, nil
}

func setGlobalHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "hcloud-sdk-go/v"+version)
}
