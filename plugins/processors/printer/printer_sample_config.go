//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package printer

func (p *Printer) SampleConfig() string {
	return `# Print all metrics that pass through this filter.
[[processors.printer]]
`
}
