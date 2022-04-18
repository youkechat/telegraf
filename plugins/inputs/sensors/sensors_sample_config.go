//go:build linux
// +build linux

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package sensors

func (*Sensors) SampleConfig() string {
	return `# Monitor sensors, requires lm-sensors package
[[inputs.sensors]]
  ## Remove numbers from field names.
  ## If true, a field name like 'temp1_input' will be changed to 'temp_input'.
  # remove_numbers = true

  ## Timeout is the maximum amount of time that the sensors command can run.
  # timeout = "5s"
`
}
