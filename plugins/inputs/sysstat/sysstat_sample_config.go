//go:build linux
// +build linux

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package sysstat

func (*Sysstat) SampleConfig() string {
	return `# Sysstat metrics collector
[[inputs.sysstat]]
  ## Path to the sadc command.
  #
  ## Common Defaults:
  ##   Debian/Ubuntu: /usr/lib/sysstat/sadc
  ##   Arch:          /usr/lib/sa/sadc
  ##   RHEL/CentOS:   /usr/lib64/sa/sadc
  sadc_path = "/usr/lib/sa/sadc" # required

  ## Path to the sadf command, if it is not in PATH
  # sadf_path = "/usr/bin/sadf"

  ## Activities is a list of activities, that are passed as argument to the
  ## sadc collector utility (e.g: DISK, SNMP etc...)
  ## The more activities that are added, the more data is collected.
  # activities = ["DISK"]

  ## Group metrics to measurements.
  ##
  ## If group is false each metric will be prefixed with a description
  ## and represents itself a measurement.
  ##
  ## If Group is true, corresponding metrics are grouped to a single measurement.
  # group = true

  ## Options for the sadf command. The values on the left represent the sadf options and
  ## the values on the right their description (wich are used for grouping and prefixing metrics).
  ##
  ## Run 'sar -h' or 'man sar' to find out the supported options for your sysstat version.
  [inputs.sysstat.options]
    -C = "cpu"
    -B = "paging"
    -b = "io"
    -d = "disk"             # requires DISK activity
    "-n ALL" = "network"
    "-P ALL" = "per_cpu"
    -q = "queue"
    -R = "mem"
    -r = "mem_util"
    -S = "swap_util"
    -u = "cpu_util"
    -v = "inode"
    -W = "swap"
    -w = "task"
  # -H = "hugepages"        # only available for newer linux distributions
  # "-I ALL" = "interrupts" # requires INT activity

  ## Device tags can be used to add additional tags for devices. For example the configuration below
  ## adds a tag vg with value rootvg for all metrics with sda devices.
  # [[inputs.sysstat.device_tags.sda]]
  #  vg = "rootvg"
`
}
