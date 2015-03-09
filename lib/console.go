package cloudatgost

import(
	"bytes"
	"net/http"
	"net/url"
)

type CacConsole struct {
	Status string `json:"status"`
	Time int `json:"time"`
	API string `json:"api"`
	Serverid string `json:"serverid"`
	Console string `json:"console"`
}

func (c *Client) Console(serverId string) (*CacConsole) {
	v := &CacConsole{}
	Url, err := url.Parse(c.BaseURL)
	if err != nil {
		panic("boom! Busted :F")
	}
	Url.Path += "console.php"
	parameters := url.Values{}
	parameters.Add("key", c.Token)
	parameters.Add("login", c.Login)
	parameters.Add("sid", serverId)

	request, err := http.NewRequest("POST", Url.String(), bytes.NewBufferString(parameters.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil
	}

	c.Do(request, &v)
	return v
}
