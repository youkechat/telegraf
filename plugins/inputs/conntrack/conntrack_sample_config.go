//go:build linux
// +build linux

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package conntrack

func (c *Conntrack) SampleConfig() string {
	return `# Collects conntrack stats from the configured directories and files.
[[inputs.conntrack]]
  ## The following defaults would work with multiple versions of conntrack.
  ## Note the nf_ and ip_ filename prefixes are mutually exclusive across
  ## kernel versions, as are the directory locations.

  ## Superset of filenames to look for within the conntrack dirs.
  ## Missing files will be ignored.
  files = ["ip_conntrack_count","ip_conntrack_max",
          "nf_conntrack_count","nf_conntrack_max"]

  ## Directories to search within for the conntrack files above.
  ## Missing directories will be ignored.
  dirs = ["/proc/sys/net/ipv4/netfilter","/proc/sys/net/netfilter"]
`
}
