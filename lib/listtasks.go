package cloudatgost

import(
	"net/http"
	"net/url"
)

// A TaskList represents an API response that contains a list
// of tasks in operation.
type TaskList struct {
	Status string `json:"status"`
	Time int `json:"time"`
	API string `json:"api"`
	Cid string `json:"cid"`
	Action string `json:"action"`
	Data []struct {
		Cid string `json:"cid"`
		Idf string `json:"idf"`
		Serverid string `json:"serverid"`
		Action string `json:"action"`
		Status string `json:"status"`
		Starttime string `json:"starttime"`
		Finishtime string `json:"finishtime"`
	} `json:"data"`
}

// ListTasks formulates an HTTP request to the listtasks.php
// endpoint and maps the JSON response through Do to a TaskList
// structure.
func (c *Client) ListTasks() (*TaskList) {
	v := &TaskList{}
	URL, err := url.Parse(c.BaseURL)
	if err != nil {
		panic("boom! Busted :F")
	}
	URL.Path += "listtasks.php"
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
