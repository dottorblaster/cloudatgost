package cloudatgost

import(
	"net/http"
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
	url := c.BaseURL + "listtemplates.php?key=" + c.Token + "&login=" + c.Login
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	c.Do(request, &v)
	return v
}
