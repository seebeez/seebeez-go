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
	Status   int    `json:"status,omitempty"`
	Format   string `json:"format,omitempty"`
	Progress int    `json:"progress,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Link     string `json:"link,omitempty"`
}

// ExportData ...
type ExportData struct {
	Status   int    `json:"status,omitempty"`
	URI      string `json:"uri,omitempty"`
	Progress int    `json:"progress,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

// ResponseData ...
type ResponseData struct {
	ID       string         `json:"id,omitempty"`
	Download []DownloadData `json:"download,omitempty"`
	Convert  []ConvertData  `json:"convert,omitempty"`
	Export   []ExportData   `json:"export,omitempty"`
	Duration int            `json:"duration,omitempty"`
}

// JobResponse is the combined response from CovertData, ExportData, ResponseData
type JobResponse struct {
	Data ResponseData `json:"data"`
}
