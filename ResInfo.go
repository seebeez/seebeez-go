package seebeez

// Response returned when a job is dispatched
type ResInfo struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Code   int    `json:"code"`
}
// Gets the id of the job
func (r *ResInfo) GetId() string {
	return r.Id
}
// Gets the status of the job
func (r *ResInfo) GetStatus() string {
	return r.Status
}
// Gets the status code of the job
func (r *ResInfo) GetCode() int {
	return r.Code
}
// Return response on the state of the requested job
func (r *ResInfo) GetJobInfo() (SeebeezResponse, error) {
	handler := requestHandler{}
	resp, err := handler.checkStatus(*r)
	if err != nil {
		return SeebeezResponse{}, err
	}
	return resp, nil
}
