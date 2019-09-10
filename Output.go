package seebeez

// Output type for export
type Output struct {
	Format  string   `json:"format"`
	Codec   string   `json:"codec"`
	Exports []string `json:"exports"`
}

// AddExport(string) adds export link to the Output
func (o *Output) AddExport(export string) *Output {
	o.Exports = append(o.Exports, export)
	return o
}

// ClearExports() will remove all the exports
func (o *Output) ClearExports() *Output {
	o.Exports = []string{}
	return o
}

// SetFormat(string) sets the file output format
func (o *Output) SetFormat(format string) *Output {
	o.Format = format
	return o
}

// SetCodec(string) sets the codec for the output file
func (o *Output) SetCodec(format string) *Output {
	o.Codec = format
	return o
}

// SetExports(string) sets an array of links of the export links
func (o *Output) SetExports(exports []string) *Output {
	o.Exports = exports
	return o
}

// NewOutput(format string, codec string) adds a new Output type to the Export type
func NewOutput(format, codec string) *Output {
	output := Output{}
	output.Format = format
	output.Codec = codec
	return &output
}
