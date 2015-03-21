package cloudatgost

import(
	"bytes"
	"net/http"
	"net/url"
)

type PowerOp struct {
	Status string `json:"status"`
	Time int `json:"time"`
	API string `json:"api"`
	Serverid string `json:"serverid"`
	Action string `json:"action"`
	Taskid int64 `json:"taskid"`
	Result string `json:"result"`
}

func (c *Client) Action(serverID string, operation string) (*PowerOp) {
	v := &PowerOp{}
	URL, err := url.Parse(c.BaseURL)
	if err != nil {
		panic("boom! Busted :F")
	}
	URL.Path += "powerop.php"
	parameters := url.Values{}
	parameters.Add("key", c.Token)
	parameters.Add("login", c.Login)
	parameters.Add("sid", serverID)
	parameters.Add("action", operation)

	request, err := http.NewRequest("POST", URL.String(), bytes.NewBufferString(parameters.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil
	}

	c.Do(request, &v)
	return v
}

func (c *Client) PowerOn(serverID string) (*PowerOp) {
	return Action(serverID, "poweron")
}

func (c *Client) PowerOff(serverID string) (*PowerOp) {
	return Action(serverID, "poweroff")
}

func (c *Client) Reboot(serverID string) (*PowerOp) {
	return Action(serverID, "reset")
}
