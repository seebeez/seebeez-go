package seebeez

import "encoding/json"

type ServiceAPI struct {
	Url    string
	Link   string
	Format string
}

var ServiceInstance = ServiceAPI{"https://service.seebeez.com/apps/youtube-dl/json", "", ""}

func (a *ServiceAPI) SetLink(link string) {
	a.Link = link
}

func (a *ServiceAPI) SetFormat(format string) {
	a.Format = format
}

func (a *ServiceAPI) SetUrl(url string) {
	a.Url = url
}

func (a *ServiceAPI) GetJson() ([]byte, error) {
	Handler := RequestHandler{}
	return Handler.GetServiceDetails(a)
}

func (a *ServiceAPI) GetJsonString() (string, error) {
	body, err := a.GetJson()
	return string(body), err
}

func (a *ServiceAPI) GetMap() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	body, err := a.GetJson()
	err = json.Unmarshal(body, &m)
	return m, err
}
