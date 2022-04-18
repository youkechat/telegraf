//go:build linux
// +build linux

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package mdstat

func (k *MdstatConf) SampleConfig() string {
	return `# Get kernel statistics from /proc/mdstat
[[inputs.mdstat]]
  ## Sets file path
  ## If not specified, then default is /proc/mdstat
  # file_name = "/proc/mdstat"
`
}
