//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package ifname

func (d *IfName) SampleConfig() string {
	return `# Add a tag of the network interface name looked up over SNMP by interface number
[[processors.ifname]]
  ## Name of tag holding the interface number
  # tag = "ifIndex"

  ## Name of output tag where service name will be added
  # dest = "ifName"

  ## Name of tag of the SNMP agent to request the interface name from
  # agent = "agent"

  ## Timeout for each request.
  # timeout = "5s"

  ## SNMP version; can be 1, 2, or 3.
  # version = 2

  ## SNMP community string.
  # community = "public"

  ## Number of retries to attempt.
  # retries = 3

  ## The GETBULK max-repetitions parameter.
  # max_repetitions = 10

  ## SNMPv3 authentication and encryption options.
  ##
  ## Security Name.
  # sec_name = "myuser"
  ## Authentication protocol; one of "MD5", "SHA", or "".
  # auth_protocol = "MD5"
  ## Authentication password.
  # auth_password = "pass"
  ## Security Level; one of "noAuthNoPriv", "authNoPriv", or "authPriv".
  # sec_level = "authNoPriv"
  ## Context Name.
  # context_name = ""
  ## Privacy protocol used for encrypted messages; one of "DES", "AES" or "".
  # priv_protocol = ""
  ## Privacy password used for encrypted messages.
  # priv_password = ""

  ## max_parallel_lookups is the maximum number of SNMP requests to
  ## make at the same time.
  # max_parallel_lookups = 100

  ## ordered controls whether or not the metrics need to stay in the
  ## same order this plugin received them in. If false, this plugin
  ## may change the order when data is cached.  If you need metrics to
  ## stay in order set this to true.  keeping the metrics ordered may
  ## be slightly slower
  # ordered = false

  ## cache_ttl is the amount of time interface names are cached for a
  ## given agent.  After this period elapses if names are needed they
  ## will be retrieved again.
  # cache_ttl = "8h"
`
}
