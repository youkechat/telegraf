//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package nfsclient

func (n *NFSClient) SampleConfig() string {
	return `# Read per-mount NFS client metrics from /proc/self/mountstats
[[inputs.nfsclient]]
  ## Read more low-level metrics (optional, defaults to false)
  # fullstat = false

  ## List of mounts to explictly include or exclude (optional)
  ## The pattern (Go regexp) is matched against the mount point (not the
  ## device being mounted).  If include_mounts is set, all mounts are ignored
  ## unless present in the list. If a mount is listed in both include_mounts
  ## and exclude_mounts, it is excluded.  Go regexp patterns can be used.
  # include_mounts = []
  # exclude_mounts = []

  ## List of operations to include or exclude from collecting.  This applies
  ## only when fullstat=true.  Symantics are similar to {include,exclude}_mounts:
  ## the default is to collect everything; when include_operations is set, only
  ## those OPs are collected; when exclude_operations is set, all are collected
  ## except those listed.  If include and exclude are set, the OP is excluded.
  ## See /proc/self/mountstats for a list of valid operations; note that
  ## NFSv3 and NFSv4 have different lists.  While it is not possible to
  ## have different include/exclude lists for NFSv3/4, unused elements
  ## in the list should be okay.  It is possible to have different lists
  ## for different mountpoints:  use mulitple [[input.nfsclient]] stanzas,
  ## with their own lists.  See "include_mounts" above, and be careful of
  ## duplicate metrics.
  # include_operations = []
  # exclude_operations = []
`
}
