//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package ethtool

func (e *Ethtool) SampleConfig() string {
	return `# Returns ethtool statistics for given interfaces
[[inputs.ethtool]]
  ## List of interfaces to pull metrics for
  # interface_include = ["eth0"]

  ## List of interfaces to ignore when pulling metrics.
  # interface_exclude = ["eth1"]

  ## Some drivers declare statistics with extra whitespace, different spacing,
  ## and mix cases. This list, when enabled, can be used to clean the keys.
  ## Here are the current possible normalizations:
  ##  * snakecase: converts fooBarBaz to foo_bar_baz
  ##  * trim: removes leading and trailing whitespace
  ##  * lower: changes all capitalized letters to lowercase
  ##  * underscore: replaces spaces with underscores
  # normalize_keys = ["snakecase", "trim", "lower", "underscore"]
`
}
