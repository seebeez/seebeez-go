package seebeez

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type requestHandler struct{}

func (r *requestHandler) handle(s Seebeez) (response, error) {
	// Stop application if no Auth Token is found
	if os.Getenv("SeebeezAuth") == "" {
		log.Fatal("No authorization token is set!")
		return response{}, errors.New("no seeebeez auth token")
	}

	// Prepare JSON
	obj, err := json.Marshal(s)
	req, err := http.NewRequest("POST", getURL("job"), bytes.NewBuffer(obj))
	if err != nil {
		log.Fatal(err.Error())
		return response{}, err
	}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SeebeezAuth"))

	// Prepare an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return response{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return response{body}, nil
}

func (r *requestHandler) checkStatus(res ResInfo) (SeebeezResponse, error) {
	// Stop application if no Auth Token is found
	if os.Getenv("SeebeezAuth") == "" {
		log.Fatal("No authorization token is set!")
		return SeebeezResponse{}, errors.New("No AUTH Token!")
	}

	req, err := http.NewRequest("GET", getURL("job/"+res.ID), nil)
	if err != nil {
		log.Fatal(err.Error())
		return SeebeezResponse{}, err
	}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SeebeezAuth"))

	// Prepare an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return SeebeezResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	data := SeebeezResponse{}
	err = json.Unmarshal(body, &data)
	return data, nil
}

func (r *requestHandler) getServiceDetails(a *ServiceAPI) ([]byte, error) {
	serviceJson := struct {
		Link   string `json:"link"`
		Format string `json:"format"`
	}{a.Link, a.Format}

	obj, err := json.Marshal(serviceJson)
	req, err := http.NewRequest("POST", a.URL, bytes.NewBuffer(obj))
	if err != nil {
		log.Fatal(err.Error())
		return []byte{}, err
	}

	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Prepare an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return []byte{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body, nil
}
