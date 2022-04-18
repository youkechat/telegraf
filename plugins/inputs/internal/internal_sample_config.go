//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package internal

func (s *Self) SampleConfig() string {
	return `# Collect statistics about itself
[[inputs.internal]]
  ## If true, collect telegraf memory stats.
  # collect_memstats = true
`
}
