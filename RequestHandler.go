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

type RequestHandler struct {}
func (r *RequestHandler) Handle(s Seebeez) (Response, error) {
	// Stop application if no Auth Token is found
	if os.Getenv("SeebeezAuth")==""{
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

func (r *RequestHandler) CheckStatus(res ResInfo) (ResInfo, error) {
	// Stop application if no Auth Token is found
	if os.Getenv("SeebeezAuth")==""{
		log.Fatal("No authorization token is set!")
		return ResInfo{}, errors.New("No AUTH Token!")
	}

	req, err := http.NewRequest("GET", GetURL("job/"+res.Id), nil)
	if err != nil {
		log.Fatal(err.Error())
		return ResInfo{}, err
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
		return ResInfo{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
	return ResInfo{}, nil
}
