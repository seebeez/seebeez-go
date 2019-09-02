package seebeez

type Output struct {
	Format string `json:"format"`
	Codec string `json:"codec"`
	Exports []string `json:"exports"`
}

func (o *Output) AddExport(export string) *Output {
	o.Exports = append(o.Exports, export)
	return o
}

func (o *Output) ClearExports() *Output{
	o.Exports = []string{}
	return o
}

func (o *Output) SetFormat(format string) *Output {
	o.Format = format
	return o
}

func (o *Output) SetCodec(format string) *Output {
	o.Codec = format
	return o
}

func (o *Output) SetExports(exports []string) *Output {
	o.Exports = exports
	return o
}

func NewOutput(format, codec string) *Output{
	output := Output{}
	output.Format = format
	output.Codec = codec
	return &output
}