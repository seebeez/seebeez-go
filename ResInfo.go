package seebeez

// ResInfo is returned when a job is dispatched
type ResInfo struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Code   int    `json:"code"`
}
// GetID retrieves the ID of the dispatched job
func (r *ResInfo) GetID() string {
	return r.ID
}
// GetStatus retrieves the status of the job
func (r *ResInfo) GetStatus() string {
	return r.Status
}
// GetCode retrieves the status code of the job
func (r *ResInfo) GetCode() int {
	return r.Code
}
// GetJobInfo returns a response on the state of the requested job
func (r *ResInfo) GetJobInfo() (SeebeezResponse, error) {
	handler := requestHandler{}
	resp, err := handler.checkStatus(*r)
	if err != nil {
		return SeebeezResponse{}, err
	}
	return resp, nil
}
