//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package ping

func (*Ping) SampleConfig() string {
	return `# Ping given url(s) and return statistics
[[inputs.ping]]
  ## Hosts to send ping packets to.
  urls = ["example.org"]

  ## Method used for sending pings, can be either "exec" or "native".  When set
  ## to "exec" the systems ping command will be executed.  When set to "native"
  ## the plugin will send pings directly.
  ##
  ## While the default is "exec" for backwards compatibility, new deployments
  ## are encouraged to use the "native" method for improved compatibility and
  ## performance.
  # method = "exec"

  ## Number of ping packets to send per interval.  Corresponds to the "-c"
  ## option of the ping command.
  # count = 1

  ## Time to wait between sending ping packets in seconds.  Operates like the
  ## "-i" option of the ping command.
  # ping_interval = 1.0

  ## If set, the time to wait for a ping response in seconds.  Operates like
  ## the "-W" option of the ping command.
  # timeout = 1.0

  ## If set, the total ping deadline, in seconds.  Operates like the -w option
  ## of the ping command.
  # deadline = 10

  ## Interface or source address to send ping from.  Operates like the -I or -S
  ## option of the ping command.
  # interface = ""

  ## Percentiles to calculate. This only works with the native method.
  # percentiles = [50, 95, 99]

  ## Specify the ping executable binary.
  # binary = "ping"

  ## Arguments for ping command. When arguments is not empty, the command from
  ## the binary option will be used and other options (ping_interval, timeout,
  ## etc) will be ignored.
  # arguments = ["-c", "3"]

  ## Use only IPv6 addresses when resolving a hostname.
  # ipv6 = false

  ## Number of data bytes to be sent. Corresponds to the "-s"
  ## option of the ping command. This only works with the native method.
  # size = 56
`
}
