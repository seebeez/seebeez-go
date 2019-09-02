package seebeez

const API = "https://beta.seebeez.com/api/"
const VERSION = "v1"

func GetURL(endpoint string) string {
	return API + VERSION + "/" + endpoint
}
