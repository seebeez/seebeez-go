package seebeez

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type response struct {
	Bytes []byte
}

// Byte returns the []byte format of the response
func (r *response) Byte() []byte {
	body, _ := ioutil.ReadAll(bytes.NewReader(r.Bytes))
	//if err != nil {
	//	return []byte("")
	//}
	return body
}

// String returns the string format of the response
func (r *response) String() string {
	return string(r.Byte())
}

// ResInfo returns response containing dispatched job information
func (r *response) Info() (ResInfo, error) {
	var info ResInfo
	err := json.Unmarshal(r.Byte(), &info)
	return info, err
}
