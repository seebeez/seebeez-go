package seebeez

import (
	"os"
)

var SeebeezInstance *Seebeez

// Seebeez is used for creating a Seebeez request with a given source, server type and outputs
type Seebeez struct {
	Source  string   `json:"source"`
	Server  string   `json:"server"`
	Outputs []Output `json:"outputs"`
}

// Init initialize instance
func Init(source, server string) *Seebeez {
	if SeebeezInstance == nil {
		SeebeezInstance = new(Seebeez)
		SeebeezInstance.Source = source
		SeebeezInstance.Server = server
	}
	return SeebeezInstance
}

// AddOutput adds an output to the export type
func (s *Seebeez) AddOutput(output *Output) *Seebeez {
	s.Outputs = append(s.Outputs, *output)
	return s
}

// ClearOutput clears the output array
func (s *Seebeez) ClearOutput(output Output) *Seebeez {
	s.Outputs = []Output{}
	return s
}

// SetSource sets source url of the file
func (s *Seebeez) SetSource(source string) *Seebeez {
	s.Source = source
	return s
}

// SetServer sets server type
func (s *Seebeez) SetServer(server string) *Seebeez {
	s.Server = server
	return s
}

// Set sets source url of the file and server type
func (s *Seebeez) Set(source, server string) *Seebeez {
	s.Source = source
	s.Server = server
	return s
}

// SetToken sets the SeebeezAuth environment variable
func (s *Seebeez) SetToken(token string) *Seebeez {
	os.Setenv("SeebeezAuth", token)
	return s
}

// MakeRequest makes a request based on the export information and return response with job information
func (s *Seebeez) MakeRequest() (response, error) {
	handler := requestHandler{}
	resp, err := handler.handle(*s)
	return resp, err
}
