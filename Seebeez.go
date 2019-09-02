package seebeez

import (
	"os"
)

var SeebeezInstance *Seebeez

type Seebeez struct {
	Source  string `json:"source"`
	Server  string `json:"server"`
	Outputs []Output `json:"outputs"`
}

func Init(source, server string) *Seebeez {
	if SeebeezInstance == nil {
		SeebeezInstance = new(Seebeez)
		SeebeezInstance.Source = source
		SeebeezInstance.Server = server
	}
	return SeebeezInstance
}

func (s *Seebeez) AddOutput(output *Output) *Seebeez {
	s.Outputs = append(s.Outputs, *output)
	return s
}

func (s *Seebeez) ClearOutput(output Output) *Seebeez {
	s.Outputs = []Output{}
	return s
}

func (s *Seebeez) SetSource(source string) *Seebeez {
	s.Source = source
	return s
}

func (s *Seebeez) SetServer(server string) *Seebeez {
	s.Server = server
	return s
}

func (s *Seebeez) Set(source, server string) *Seebeez {
	s.Source = source
	s.Server = server
	return s
}

func (s *Seebeez) SetToken(token string) *Seebeez {
	os.Setenv("SeebeezAuth", token)
	return s
}

func (s *Seebeez) MakeRequest() (Response, error) {
	handler := RequestHandler{}
	resp, err := handler.Handle(*s)
	return resp, err
}

