package seebeez

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Response struct {
	Bytes []byte
}

func (r *Response) Byte() []byte {
	body, err := ioutil.ReadAll(bytes.NewReader(r.Bytes))
	if err != nil {
		return []byte("")
	}
	return body
}
func (r *Response) String() string {
	return string(r.Byte())
}

//func (r *Response) Resend() (Response, error) {
//	r.Req.Header.Set("Content-Type", "application/json")
//	r.Req.Header.Set("Accept", "application/json")
//	r.Req.Header.Set("Authorization", "Bearer "+os.Getenv("SeebeezAuth"))
//	resp, err := r.Client.Do(r.Req)
//	if err != nil {
//		log.Fatal(err)
//		return Response{}, nil
//	}
//	r.Resp = resp
//	return *r, nil
//}

func (r *Response) Info() (ResInfo, error) {
	var info ResInfo
	fmt.Println(string(r.Byte()))
	err := json.Unmarshal(r.Byte(), &info)
	//err := json.NewDecoder(r.Resp.Body).Decode(info)
	if err != nil {
		return ResInfo{}, err
	}
	return info, nil
}
