//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package disque

func (d *Disque) SampleConfig() string {
	return `# Read metrics from one or many disque servers
[[inputs.disque]]
  ## An array of URI to gather stats about. Specify an ip or hostname
  ## with optional port and password.
  ## ie disque://localhost, disque://10.10.3.33:18832, 10.0.0.1:10000, etc.
  ## If no servers are specified, then localhost is used as the host.
  servers = ["localhost"]
`
}
