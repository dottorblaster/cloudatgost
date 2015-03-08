package cloudatgost

import(
	"net/http"
	"net/url"
)

type TemplateList struct {
	Status string `json:"status"`
	Time int `json:"time"`
	Data []struct {
		ID string `json:"id"`
		Detail string `json:"detail"`
	} `json:"data"`
}

func (c *Client) ListTemplates() (*TemplateList) {
	v := &TemplateList{}
	Url, err := url.Parse(c.BaseURL)
	if err != nil {
		panic("boom! Busted :F")
	}
	Url.Path += "listtemplates.php"
	parameters := url.Values{}
	parameters.Add("key", c.Token)
	parameters.Add("login", c.Login)
	Url.RawQuery = parameters.Encode()
	request, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		return nil
	}

	c.Do(request, &v)
	return v
}
