//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package bond

func (bond *Bond) SampleConfig() string {
	return `# Collect bond interface status, slaves statuses and failures count
[[inputs.bond]]
  ## Sets 'proc' directory path
  ## If not specified, then default is /proc
  # host_proc = "/proc"

  ## Sets 'sys' directory path
  ## If not specified, then default is /sys
  # host_sys = "/sys"

  ## By default, telegraf gather stats for all bond interfaces
  ## Setting interfaces will restrict the stats to the specified
  ## bond interfaces.
  # bond_interfaces = ["bond0"]

  ## Tries to collect additional bond details from /sys/class/net/{bond}
  ## currently only useful for LACP (mode 4) bonds
  # collect_sys_details = false
`
}
