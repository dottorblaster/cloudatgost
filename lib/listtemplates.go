package cloudatgost

import(
	"net/http"
	"net/url"
)

// A TemplateList represents an API response for machine templates
// currently available.
type TemplateList struct {
	Status string `json:"status"`
	Time int `json:"time"`
	Data []struct {
		ID string `json:"id"`
		Detail string `json:"detail"`
	} `json:"data"`
}

// ListTemplates formulates an HTTP request to the listtemplates.php
// endpoint and maps the JSON response through Do to a TemplateList
// structure.
func (c *Client) ListTemplates() (*TemplateList) {
	v := &TemplateList{}
	URL, err := url.Parse(c.BaseURL)
	if err != nil {
		panic("boom! Busted :F")
	}
	URL.Path += "listtemplates.php"
	parameters := url.Values{}
	parameters.Add("key", c.Token)
	parameters.Add("login", c.Login)
	URL.RawQuery = parameters.Encode()

	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return nil
	}

	c.Do(request, &v)
	return v
}
