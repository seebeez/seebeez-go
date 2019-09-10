package seebeez

// Download data format
type DownloadData struct {
	Status   int    `json:"status"`
	Source   string `json:"source"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
	Link     string `json:"link"`
}
// Convert data format
type ConvertData struct {
	Status   int    `json:"status"`
	Format   string `json:"format"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
	Link     string `json:"link"`
}
// Export data format
type ExportData struct {
	Status   int    `json:"status"`
	Uri      string `json:"uri"`
	Progress int    `json:"progress"`
	Duration int    `json:"duration"`
}
// Response data format
type ResponseData struct {
	Id       string         `json:"id"`
	Download []DownloadData `json:"download"`
	Convert  []ConvertData  `json:"conver"`
	Export   []ExportData   `json:"export"`
	Duration int            `json:"duration"`
}
// Combined response
type SeebeezResponse struct {
	Data ResponseData `json:"data"`
}