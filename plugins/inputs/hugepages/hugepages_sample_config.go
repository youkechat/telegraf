//go:build linux
// +build linux

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package hugepages

func (h *Hugepages) SampleConfig() string {
	return `# Gathers huge pages measurements.
[[inputs.hugepages]]
  ## Supported huge page types:
  ##   - "root" - based on root huge page control directory: /sys/kernel/mm/hugepages
  ##   - "per_node" - based on per NUMA node directories: /sys/devices/system/node/node[0-9]*/hugepages
  ##   - "meminfo" - based on /proc/meminfo file
  # types = ["root", "per_node"]
`
}
