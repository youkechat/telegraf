//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package netstat

func (n *NetStats) SampleConfig() string {
	return `# Read TCP metrics such as established, time wait and sockets counts.
[[inputs.netstat]]
  # no configuration
`
}
