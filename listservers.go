package cloudatgost

import(
	"net/http"
	"net/url"
)

// A ServerList represents an API response that contains a list
// of servers owned by the user, with data about them.
type ServerList struct {
	Status string `json:"status"`
	Time int `json:"time"`
	API string `json:"api"`
	Action string `json:"action"`
	Data []struct {
		Sid string `json:"sid"`
		ID string `json:"id"`
		Packageid string `json:"packageid"`
		Servername string `json:"servername"`
		Lable interface{} `json:"lable"`
		Vmname string `json:"vmname"`
		IP string `json:"ip"`
		Netmask string `json:"netmask"`
		Gateway string `json:"gateway"`
		Portgroup string `json:"portgroup"`
		Hostname string `json:"hostname"`
		Rootpass string `json:"rootpass"`
		Vncport string `json:"vncport"`
		Vncpass string `json:"vncpass"`
		Servertype string `json:"servertype"`
		Template string `json:"template"`
		CPU string `json:"cpu"`
		Cpuusage string `json:"cpuusage"`
		RAM string `json:"ram"`
		Ramusage string `json:"ramusage"`
		Storage string `json:"storage"`
		Hdusage string `json:"hdusage"`
		Sdate string `json:"sdate"`
		Status string `json:"status"`
		PanelNote string `json:"panel_note"`
		Mode string `json:"mode"`
		UID string `json:"uid"`
	} `json:"data"`
}

// ListServers formulates an HTTP request to the listservers.php
// endpoint and maps the JSON response through Do to a ServerList
// structure.
func (c *Client) ListServers() (*ServerList) {
	v := &ServerList{}
	URL, err := url.Parse(c.BaseURL)
	if err != nil {
		panic("boom! Busted :F")
	}
	URL.Path += "listservers.php"
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
