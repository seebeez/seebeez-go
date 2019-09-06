package seebeez

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Under development

type DownloadData struct {
	Status   int    `json:"status"'`
	Source   string `json:"source"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
	Link     string `json:"link"`
}

type ConvertData struct {
	Status   int    `json:"status"`
	Format   string `json:"format"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
	Link     string `json:"link"`
}

type ExportData struct {
	Status   int    `json:"status"`
	Uri      string `json:"uri"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
}

type ResponseData struct {
	Id       string         `json:"id"`
	Download []DownloadData `json:"download"`
	Convert  []ConvertData  `json:"conver"`
	Export   []ExportData   `json:"export"`
	Duration int            `json:"duration"`
}

type SeebeezResponse struct {
	Data ResponseData `json:"data"`
}

func GetJsonAndBind(url string, obj interface{}) error {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	jsonBody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonBody, &obj)
	if err != nil {
		return err
	}
	return nil
}
