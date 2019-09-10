package seebeez

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

type requestHandler struct{}

func (r *requestHandler) handle(s Seebeez) (response, error) {
	// Stop application if no Auth Token is found
	if os.Getenv("SeebeezAuth") == "" {
		return response{}, errors.New("no auth token set")
	}

	// Prepare JSON
	obj, err := json.Marshal(s)
	req, err := http.NewRequest("POST", getURL("job"), bytes.NewBuffer(obj))
	//if err != nil {
	//	return response{}, err
	//}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SeebeezAuth"))

	// Prepare an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return response{body}, err
}

func (r *requestHandler) checkStatus(res ResInfo) (JobResponse, error) {
	// Stop application if no Auth Token is found
	if os.Getenv("SeebeezAuth") == "" {
		return JobResponse{}, errors.New("no auth token")
	}

	req, err := http.NewRequest("GET", getURL("job/"+res.ID), nil)
	//if err != nil {
	//	return JobResponse{}, err
	//}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SeebeezAuth"))

	// Prepare an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return JobResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	data := JobResponse{}
	err = json.Unmarshal(body, &data)
	return data, err
}

func (r *requestHandler) getServiceDetails(a *ServiceAPI) ([]byte, error) {
	serviceJSON := struct {
		Link   string `json:"link"`
		Format string `json:"format"`
	}{a.Link, a.Format}

	obj, err := json.Marshal(serviceJSON)
	req, err := http.NewRequest("POST", a.URL, bytes.NewBuffer(obj))

	//if err != nil {
	//	return []byte{}, err
	//}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Prepare an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body, err
}
