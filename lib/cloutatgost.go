// CloudAtCost API wrapper written in Go.
package cloudatgost

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	client *http.Client
	BaseURL string
	Login string
	Token string
}

func NewClient(email string, token string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL := "https://panel.cloudatcost.com/api/v1/"

	c := &Client{
		Login: email,
		Token: token,
		client: httpClient,
		BaseURL: baseURL,
	}

	return c
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	if v != nil {
		json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
