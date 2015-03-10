package cloudatgost

import(
	"net/http"
	"net/url"
)

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

func (c *Client) ListTasks() (*TaskList) {
	v := &TaskList{}
	Url, err := url.Parse(c.BaseURL)
	if err != nil {
		panic("boom! Busted :F")
	}
	Url.Path += "listtasks.php"
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
