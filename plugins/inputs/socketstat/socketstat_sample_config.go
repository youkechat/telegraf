//go:build !windows
// +build !windows

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package socketstat

func (ss *Socketstat) SampleConfig() string {
	return `# Gather indicators from established connections, using iproute2's ss command.
[[inputs.socketstat]]
  ## ss can display information about tcp, udp, raw, unix, packet, dccp and sctp sockets
  ## Specify here the types you want to gather
  socket_types = [ "tcp", "udp" ]
  ## The default timeout of 1s for ss execution can be overridden here:
  # timeout = "1s"
`
}
