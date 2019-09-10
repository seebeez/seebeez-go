package seebeez

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type response struct {
	Bytes []byte
}

// Returns the []byte format of the response
func (r *response) Byte() []byte {
	body, err := ioutil.ReadAll(bytes.NewReader(r.Bytes))
	if err != nil {
		return []byte("")
	}
	return body
}

// Returns the string format of the response
func (r *response) String() string {
	return string(r.Byte())
}

// Returns response containing dispatched job information
func (r *response) Info() (ResInfo, error) {
	var info ResInfo
	fmt.Println(string(r.Byte()))
	err := json.Unmarshal(r.Byte(), &info)
	//err := json.NewDecoder(r.Resp.Body).Decode(info)
	if err != nil {
		return ResInfo{}, err
	}
	return info, nil
}
