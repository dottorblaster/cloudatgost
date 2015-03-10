package cloudatgost

import(
	"bytes"
	"net/http"
	"net/url"
)

// A CacConsole represents an API response that contains an URL
// to access a VNC console for the given machines' server ID.
type CacConsole struct {
	Status string `json:"status"`
	Time int `json:"time"`
	API string `json:"api"`
	Serverid string `json:"serverid"`
	Console string `json:"console"`
}

// Console formulates an HTTP request to the console.php endpoint
// and maps the JSON response through Do to a CacConsole structure.
func (c *Client) Console(serverID string) (*CacConsole) {
	v := &CacConsole{}
	URL, err := url.Parse(c.BaseURL)
	if err != nil {
		panic("boom! Busted :F")
	}
	URL.Path += "console.php"
	parameters := url.Values{}
	parameters.Add("key", c.Token)
	parameters.Add("login", c.Login)
	parameters.Add("sid", serverID)

	request, err := http.NewRequest("POST", URL.String(), bytes.NewBufferString(parameters.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil
	}

	c.Do(request, &v)
	return v
}
