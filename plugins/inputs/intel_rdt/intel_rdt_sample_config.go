//go:build !windows
// +build !windows

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package intel_rdt

func (r *IntelRDT) SampleConfig() string {
	return `# Read Intel RDT metrics
[[inputs.intel_rdt]]
  ## Optionally set sampling interval to Nx100ms. 
  ## This value is propagated to pqos tool. Interval format is defined by pqos itself.
  ## If not provided or provided 0, will be set to 10 = 10x100ms = 1s.
  # sampling_interval = "10"
 
  ## Optionally specify the path to pqos executable. 
  ## If not provided, auto discovery will be performed.
  # pqos_path = "/usr/local/bin/pqos"

  ## Optionally specify if IPC and LLC_Misses metrics shouldn't be propagated.
  ## If not provided, default value is false.
  # shortened_metrics = false

  ## Specify the list of groups of CPU core(s) to be provided as pqos input. 
  ## Mandatory if processes aren't set and forbidden if processes are specified.
  ## e.g. ["0-3", "4,5,6"] or ["1-3,4"]
  # cores = ["0-3"]

  ## Specify the list of processes for which Metrics will be collected.
  ## Mandatory if cores aren't set and forbidden if cores are specified.
  ## e.g. ["qemu", "pmd"]
  # processes = ["process"]

  ## Specify if the pqos process should be called with sudo.
  ## Mandatory if the telegraf process does not run as root.
  # use_sudo = false
`
}
