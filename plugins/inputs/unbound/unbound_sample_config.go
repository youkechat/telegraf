//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package unbound

func (s *Unbound) SampleConfig() string {
	return `# A plugin to collect stats from the Unbound DNS resolver
[[inputs.unbound]]
  ## Address of server to connect to, read from unbound conf default, optionally ':port'
  ## Will lookup IP if given a hostname
  server = "127.0.0.1:8953"

  ## If running as a restricted user you can prepend sudo for additional access:
  # use_sudo = false

  ## The default location of the unbound-control binary can be overridden with:
  # binary = "/usr/sbin/unbound-control"

  ## The default location of the unbound config file can be overridden with:
  # config_file = "/etc/unbound/unbound.conf"

  ## The default timeout of 1s can be overridden with:
  # timeout = "1s"

  ## When set to true, thread metrics are tagged with the thread id.
  ##
  ## The default is false for backwards compatibility, and will be changed to
  ## true in a future version.  It is recommended to set to true on new
  ## deployments.
  thread_as_tag = false
`
}
