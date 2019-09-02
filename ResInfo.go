package seebeez

type ResInfo struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Code int `json:"code"`
}

func (r *ResInfo) GetId() string {
	return r.Id
}

func (r *ResInfo) GetStatus() string {
	return r.Status
}

func (r *ResInfo) GetCode() int {
	return r.Code
}

func (r *ResInfo) GetJobInfo() (SeebeezResponse, error) {
	handler := RequestHandler{}
	resp, err := handler.CheckStatus(*r)
	if err != nil {
		return SeebeezResponse{}, err
	}
	return resp, nil
}
