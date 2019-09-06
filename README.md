# Seebeez Go
Seebeez is a super fast transcoding service made for developers for converting all types of file formats.

## Contents

- [Installation](#installation)
- [Quick-start](#quick-start)

## Installation

Install [Go](https://golang.org/) (**version 1.11+**) and get the following package

```sh
$ go get -u https://github.com/seebeez/seebeez-go
```

## Quick Start

Seebeez first fetches the file of the provided **format** from the **source** and exports the file to your storage based on the **export** link. You can also specify the **codec** you want to use on the provided file **format**. You can implicitly set the **server** type of your job instance.

Please note that, you need obtain an [API TOKEN](https://seebeez.com/api) to use this library.
```go
import "github.com/seebeez/seebeez-go"

source := "https://www.example.com/test-file.mp4"
server := "micro" // Defaults to "nano" if not set
format := "m4a"
codec := "copy" 
export := "s3://key:secret@endpoint/bucket/objectkey.m4a"
os.Setenv("SeebeezAuth", "YOUR_SEEBEEZ_KEY_GOES_HERE")
	
// Specify output format and export link
output := Output{format, codec, []string{export}}
instance := Seebeez{source, server, []Output{output}}
	
// Create a request
instance.MakeRequest()
resp, err := instance.MakeRequest()

// Get Information of the job dispatched
info, err := resp.Info()
fmt.Println("Job Information:", info, err)

// Check Information of the dispatched job
i, e := info.GetJobInfo()
fmt.Println("Job Status:", i.Data.Id, e)
```
