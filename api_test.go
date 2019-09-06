package seebeez

import (
	"os"
	"testing"
)

var RESP Response
var RESINFO ResInfo

func TestMakeRequest(t *testing.T) {
	source := os.Getenv("TEST_SOURCE")
	server := os.Getenv("TEST_SERVER")
	format := os.Getenv("TEST_FORMAT")
	codec := os.Getenv("TEST_CODEC")
	export := os.Getenv("TEST_EXPORT")
	output := Output{format, codec, []string{export}}
	instance := Seebeez{source, server, []Output{output}}
	res, err := instance.MakeRequest()
	if err != nil {
		t.Fail()
	}
	RESP = res
	t.Log("Request made successfully!")
}

func TestCreateJob(t *testing.T) {
	res, err := RESP.Info()
	if err != nil {
		t.Fatal("Job disaptch error!")
		t.Fail()
	}
	RESINFO = res
	t.Log("Job dispatched successfully!")
}
func TestGetJobInfo(t *testing.T) {
	_, err := RESINFO.GetJobInfo()
	if err != nil {
		t.Fatal("Job info retrieval error!")
		t.Fail()
	}
	t.Log("Job info fetched successfully!")
}
func TestServiceAPI(t *testing.T) {
	ServiceInstance.Format = "bestaudio"
	ServiceInstance.Link = os.Getenv("TEST_SOURCE")
	_, err := ServiceInstance.GetMap()
	if err != nil {
		t.Fatal("Service API not working!")
		t.Fail()
	}
	t.Log("ServiceAPI fetched information successfully!")
}
