package seebeez

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var RESP response
var SB Seebeez
var RESINFO ResInfo

var source = "https://www.youtube.com/watch?v=h6fcK_fRYaI&list=LL9kFC3S_lUn5WJFmhk9nPZg&index=2&t=0s"
var server = "micro"
var format = "m4a"
var codec = "copy"
var export = "export"

//func TestMain(t *testing.M) {
//
//	//os.Exit(t.Run())
//	//gin.SetMode(gin.ReleaseMode)
//	//
//	//router := gin.New()
//	//router.POST("/job", func(c *gin.Context) {
//	//	c.JSON(200, gin.H{
//	//		"id":     "test_job_id",
//	//		"status": "test_status",
//	//		"code":   200,
//	//	})
//	//})
//	//router.POST("/job/test_job_id", func(c *gin.Context) {
//	//	c.JSON(200, gin.H{
//	//		"data": JobResponse{Data: ResponseData{
//	//			ID: "test_id",
//	//			Download: []DownloadData{DownloadData{
//	//				Status:   1,
//	//				Source:   "downloadurl",
//	//				Progress: 100,
//	//				Duration: 10,
//	//				Link:     "link",
//	//			}},
//	//			Convert: []ConvertData{ConvertData{
//	//				Status:   1,
//	//				Format:   "format",
//	//				Progress: 100,
//	//				Duration: 10,
//	//				Link:     "link",
//	//			}},
//	//			Export: []ExportData{ExportData{
//	//				Status:   200,
//	//				URI:      "uri",
//	//				Progress: 100,
//	//				Duration: 10,
//	//			}},
//	//			Duration: 10,
//	//		}},
//	//	})
//	//})
//	//t.Run()
//	//go router.Run(":9000")
//}
//
////func TestIsServerWorking(t *testing.T) {
////	resp, err := http.Post("http://localhost:9000/job", "application/json", bytes.NewBuffer([]byte{}))
////	body, err := ioutil.ReadAll(resp.Body)
////	fmt.Println(body, err)
////}

func TestMakeRequests(t *testing.T) {
	_ = os.Setenv("SeebeezURL", "https://service.seebeez.com/mock")
	_ = os.Setenv("SeebeezAuth", "authtest")
	fmt.Println("Test: MakeRequests")
	output := Output{format, codec, []string{export}}
	output.SetCodec(codec)
	output.SetFormat(format)
	output.ClearExports()
	output.SetExports([]string{"test_export_url1", "text_export_url2"})
	output.AddExport(export)

	instance := Seebeez{source, server, []Output{output}}
	instance.ClearOutput()
	instance.AddOutput(NewOutput(format, codec))
	instance.Set(source, server)
	instance.SetToken("authtest2")
	instance.SetServer(server)
	instance.SetSource(source)

	res, err := instance.MakeRequest()
	if err != nil {
		log.Fatal("MakeRequest()", err.Error())
	}
	fmt.Println("expected: request made successfully")
	RESP = res

	sb := Init(source, server)
	SB = *sb
}

func TestCreateJob(t *testing.T) {
	fmt.Println("Test: CreateJob")
	res, err := RESP.Info()
	resString := RESP.String()
	fmt.Println("expected: status#", res.GetID(), "code#", res.GetCode(), "status#", res.GetStatus())
	fmt.Println(resString)
	if err != nil {
		t.Fatal("Info()", err.Error())
	}
	RESINFO = res
	fmt.Println("expected: job request dispatched successfully")
}
func TestGetJobInfo(t *testing.T) {
	fmt.Println("Test: GetJobInfo")
	_, err := RESINFO.GetJobInfo()
	if err != nil {
		t.Fatal("job info retrieval error")
	}
	fmt.Println("expected: job info fetched successfully")
	_, err = RESP.Info()
	if err != nil {
		t.Fatal("Info()", err.Error())
	}
	fmt.Println("expected: file information fetched successfully")
	_, err = RESINFO.GetJobInfo()
	if err != nil {
		t.Fatal("GetJobInfo()", err.Error())
	}
	fmt.Println("expected: job info re-fetched successfully")
}

func TestServiceAPI(t *testing.T) {
	fmt.Println("Test: ServiceAPI")
	ServiceInstance.SetURL("https://service.seebeez.com/mock/youtube-dl/json")
	ServiceInstance.SetFormat("bestaudio")
	ServiceInstance.SetLink("copy")
	_, err := ServiceInstance.GetMap()
	_, err = ServiceInstance.GetJSONString()
	if err != nil {
		t.Fatal("Service API not working!")
	}
	fmt.Println("expected: could fetch job information")
}

func TestRequestHandlerNoToken(t *testing.T) {
	fmt.Println("Test: RequestHandlerNoToken")
	rh := requestHandler{}
	// No auth
	os.Setenv("SeebeezAuth", "")
	_, err := rh.checkStatus(RESINFO)
	if err != nil {
		fmt.Println("expected:", "no token detected for resinfo")
	} else {
		log.Fatal("token is parsed but, token is not set")
	}
	_, err = rh.handle(SB)
	if err != nil {
		fmt.Println("expected", "no token detected for sb")
	} else {
		log.Fatal("token is parsed but, token is not set")
	}
}

func TestRequestHandlerInvalidRequests(t *testing.T) {
	os.Setenv("SeebeezURL", "https://localhost/")
	os.Setenv("SeebeezAuth", "authtest")
	fmt.Println("Test: RequestHandlerInvalidRequests")
	rh := requestHandler{}
	RESINFO.ID = ""
	_, err := rh.checkStatus(RESINFO)
	if err != nil {
		fmt.Println("expected:", err.Error())
	} else {
		log.Fatal("error passed")
	}
	SB.Server = ""
	SB.Source = ""
	SB.Outputs = []Output{}
	_, err = rh.handle(SB)
	if err != nil {
		fmt.Println("expected:", err.Error())
	} else {
		log.Fatal("error passed")
	}
	_, err = rh.getServiceDetails(&ServiceAPI{
		URL:    "http://localhost/",
		Link:   "",
		Format: "",
	})
	if err != nil {
		fmt.Println("expected: service api invalid")
	} else {
		log.Fatal("valid service api even though it is incorrect")
	}
}

func TestDefaultURLSet(t *testing.T) {
	os.Setenv("SeebeezURL", "")
	url := getURL("")
	newURL := defaultURL + "/"
	fmt.Println("expects:", newURL)
	fmt.Println("got:", url)
	if newURL != url {
		log.Fatal("getURL() invalid url generated")
	}
	fmt.Println("expected: default url is set to", url)
}
