// Package cloudatgost is a Go wrapper for CloudAtCost API.
package cloudatgost

import (
	"encoding/json"
	"net/http"
)

// A Client represents a CloudAtCost API client. It does HTTP requests
// and returns properly populated structures.
type Client struct {
	client *http.Client
	BaseURL string
	Login string
	Token string
}

// NewClient instantiates and returns a new API client. This function
// accepts an email and a token as its parameters, as well as an optional
// HTTP client. If httpClient is nil, the function creates a new
// http.DefaultClient and assigns it to the newly created client.
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

// Do is a shorthand for HTTP requests in the cloudatgost client domain.
// The function takes an HTTP request, sends it and returns the API response,
// decoded to the v, usually passed passed by reference.
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
