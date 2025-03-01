//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package ipset

func (i *Ipset) SampleConfig() string {
	return `# Gather packets and bytes counters from Linux ipsets
  [[inputs.ipset]]
    ## By default, we only show sets which have already matched at least 1 packet.
    ## set include_unmatched_sets = true to gather them all.
    include_unmatched_sets = false
    ## Adjust your sudo settings appropriately if using this option ("sudo ipset save")
    ## You can avoid using sudo or root, by setting appropriate privileges for
    ## the telegraf.service systemd service.
    use_sudo = false
    ## The default timeout of 1s for ipset execution can be overridden here:
    # timeout = "1s"

`
}
