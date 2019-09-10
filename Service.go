package seebeez

import "encoding/json"

// ServiceAPI is used to get file information from a link using the Seebeez Service API provided the YouTube link
type ServiceAPI struct {
	URL    string
	Link   string
	Format string
}

// ServiceInstance API instance
var ServiceInstance = ServiceAPI{"https://service.seebeez.com/apps/youtube-dl/json", "", ""}

// SetLink sets the source link from where the information is to be collected
func (a *ServiceAPI) SetLink(link string) {
	a.Link = link
}

// SetFormat sets file format
func (a *ServiceAPI) SetFormat(format string) {
	a.Format = format
}

// SetURL sets the Service URL if you do not intend to use the default one
func (a *ServiceAPI) SetURL(url string) {
	a.URL = url
}

// GetJson gets the JSON from the response
func (a *ServiceAPI) GetJSON() ([]byte, error) {
	Handler := requestHandler{}
	return Handler.getServiceDetails(a)
}

// GetJsonString returns the STRING from the response
func (a *ServiceAPI) GetJsonString() (string, error) {
	body, err := a.GetJSON()
	return string(body), err
}

// GetMap gets the generated map[string]interface{} from the response
func (a *ServiceAPI) GetMap() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	body, err := a.GetJSON()
	err = json.Unmarshal(body, &m)
	return m, err
}
