package cloudatgost

import(
	"bytes"
	"net/http"
	"net/url"
)

// A PowerOp represents a successful power operation job response.
// It contains miscellaneous informations about the API and the job,
// like the task id and the result.
type PowerOp struct {
	Status string `json:"status"`
	Time int `json:"time"`
	API string `json:"api"`
	Serverid string `json:"serverid"`
	Action string `json:"action"`
	Taskid int64 `json:"taskid"`
	Result string `json:"result"`
}

// Action is a function that behaves as the actual component of the power management.
// It accepts a serverID and an operation string as its parameters, then forms an
// HTTP POST request to the endpoint. It can be used as a standalone support for
// power operations, but it serves well as a base for shorthands.
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

// PowerOn is a shorthand function that performs a "poweron" request through
// the Action function. It accepts a serverID as its unique parameter.
func (c *Client) PowerOn(serverID string) (*PowerOp) {
	return c.Action(serverID, "poweron")
}

// PowerOff is a shorthand function that performs a "poweroff" request through
// the Action function. It accepts a serverID as its unique parameter.
func (c *Client) PowerOff(serverID string) (*PowerOp) {
	return c.Action(serverID, "poweroff")
}

// Reboot is a shorthand function that performs a "reset" request through
// the Action function. It accepts a serverID as its unique parameter.
func (c *Client) Reboot(serverID string) (*PowerOp) {
	return c.Action(serverID, "reset")
}
