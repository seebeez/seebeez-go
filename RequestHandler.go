package seebeez

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RequestHandler struct{}

func (r *RequestHandler) Handle(s Seebeez) (Response, error) {
	// Stop application if no Auth Token is found
	if os.Getenv("SeebeezAuth") == "" {
		log.Fatal("No authorization token is set!")
		return Response{}, errors.New("No AUTH Token!")
	}

	// Prepare JSON
	obj, err := json.Marshal(s)
	req, err := http.NewRequest("POST", GetURL("job"), bytes.NewBuffer(obj))
	if err != nil {
		log.Fatal(err.Error())
		return Response{}, err
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
		return Response{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return Response{body}, nil
}

func (r *RequestHandler) CheckStatus(res ResInfo) (SeebeezResponse, error) {
	// Stop application if no Auth Token is found
	if os.Getenv("SeebeezAuth") == "" {
		log.Fatal("No authorization token is set!")
		return SeebeezResponse{}, errors.New("No AUTH Token!")
	}

	req, err := http.NewRequest("GET", GetURL("job/"+res.Id), nil)
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
	fmt.Println(string(body))
	err = json.Unmarshal(body, &data)
	return data, nil
}

func (r *RequestHandler) GetServiceDetails(a *ServiceAPI) ([]byte, error) {
	serviceJson := struct {
		Link   string `json:"link"`
		Format string `json:"format"`
	}{a.Link, a.Format}

	obj, err := json.Marshal(serviceJson)
	req, err := http.NewRequest("POST", a.Url, bytes.NewBuffer(obj))
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
