package seebeez

import "os"

var defaultURL = "https://beta.seebeez.com/api/v1"

func getURL(endpoint string) string {
	url := os.Getenv("SeebeezURL")
	if (url == "") {
		url = defaultURL
	}
	return url + "/" + endpoint
}
