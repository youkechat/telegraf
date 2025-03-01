//go:build !windows
// +build !windows

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package varnish

func (s *Varnish) SampleConfig() string {
	return `# A plugin to collect stats from Varnish HTTP Cache
[[inputs.varnish]]
  ## If running as a restricted user you can prepend sudo for additional access:
  #use_sudo = false

  ## The default location of the varnishstat binary can be overridden with:
  binary = "/usr/bin/varnishstat"

  ## Additional custom arguments for the varnishstat command
  # binary_args = ["-f", "MAIN.*"]

  ## The default location of the varnishadm binary can be overridden with:
  adm_binary = "/usr/bin/varnishadm"

  ## Custom arguments for the varnishadm command
  # adm_binary_args = [""]

  ## Metric version defaults to metric_version=1, use metric_version=2 for removal of nonactive vcls
  ## Varnish 6.0.2 and newer is required for metric_version=2.
  metric_version = 1

  ## Additional regexps to override builtin conversion of varnish metrics into telegraf metrics.
  ## Regexp group "_vcl" is used for extracting the VCL name. Metrics that contain nonactive VCL's are skipped.
  ## Regexp group "_field" overrides the field name. Other named regexp groups are used as tags.
  # regexps = ['^XCNT\.(?P<_vcl>[\w\-]*)(\.)*(?P<group>[\w\-.+]*)\.(?P<_field>[\w\-.+]*)\.val']

  ## By default, telegraf gather stats for 3 metric points.
  ## Setting stats will override the defaults shown below.
  ## Glob matching can be used, ie, stats = ["MAIN.*"]
  ## stats may also be set to ["*"], which will collect all stats
  stats = ["MAIN.cache_hit", "MAIN.cache_miss", "MAIN.uptime"]

  ## Optional name for the varnish instance (or working directory) to query
  ## Usually append after -n in varnish cli
  # instance_name = instanceName

  ## Timeout for varnishstat command
  # timeout = "1s"
`
}
