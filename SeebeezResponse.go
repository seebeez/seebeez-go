package seebeez

// DownloadData ...
type DownloadData struct {
	Status   int    `json:"status"`
	Source   string `json:"source"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
	Link     string `json:"link"`
}

// ConvertData ...
type ConvertData struct {
	Status   int    `json:"status"`
	Format   string `json:"format"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
	Link     string `json:"link"`
}

// ExportData ...
type ExportData struct {
	Status   int    `json:"status"`
	URI      string `json:"uri"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
}

// ResponseData ...
type ResponseData struct {
	ID       string         `json:"id"`
	Download []DownloadData `json:"download"`
	Convert  []ConvertData  `json:"conver"`
	Export   []ExportData   `json:"export"`
	Duration int            `json:"duration"`
}

// JobResponse is the combined response from CovertData, ExportData, ResponseData
type JobResponse struct {
	Data ResponseData `json:"data"`
}
