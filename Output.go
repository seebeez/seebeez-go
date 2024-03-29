package seebeez

// Output is the specified output information on what to do with the retrieved information form Export
type Output struct {
	Format  string   `json:"format"`
	Codec   string   `json:"codec"`
	Exports []string `json:"exports"`
}

// AddExport adds export link to the Output
func (o *Output) AddExport(export string) *Output {
	o.Exports = append(o.Exports, export)
	return o
}

// ClearExports will remove all the exports
func (o *Output) ClearExports() *Output {
	o.Exports = []string{}
	return o
}

// SetFormat sets the file output format
func (o *Output) SetFormat(format string) *Output {
	o.Format = format
	return o
}

// SetCodec sets the codec for the output file
func (o *Output) SetCodec(format string) *Output {
	o.Codec = format
	return o
}

// SetExports sets an array of links of the export links
func (o *Output) SetExports(exports []string) *Output {
	o.Exports = exports
	return o
}

// NewOutput adds a new Output type to the Export type
func NewOutput(format, codec string) *Output {
	output := Output{}
	output.Format = format
	output.Codec = codec
	return &output
}
