//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package statsd

func (s *Statsd) SampleConfig() string {
	return `# Statsd Server
[[inputs.statsd]]
  ## Protocol, must be "tcp", "udp4", "udp6" or "udp" (default=udp)
  protocol = "udp"

  ## MaxTCPConnection - applicable when protocol is set to tcp (default=250)
  max_tcp_connections = 250

  ## Enable TCP keep alive probes (default=false)
  tcp_keep_alive = false

  ## Specifies the keep-alive period for an active network connection.
  ## Only applies to TCP sockets and will be ignored if tcp_keep_alive is false.
  ## Defaults to the OS configuration.
  # tcp_keep_alive_period = "2h"

  ## Address and port to host UDP listener on
  service_address = ":8125"

  ## The following configuration options control when telegraf clears it's cache
  ## of previous values. If set to false, then telegraf will only clear it's
  ## cache when the daemon is restarted.
  ## Reset gauges every interval (default=true)
  delete_gauges = true
  ## Reset counters every interval (default=true)
  delete_counters = true
  ## Reset sets every interval (default=true)
  delete_sets = true
  ## Reset timings & histograms every interval (default=true)
  delete_timings = true

  ## Percentiles to calculate for timing & histogram stats.
  percentiles = [50.0, 90.0, 99.0, 99.9, 99.95, 100.0]

  ## separator to use between elements of a statsd metric
  metric_separator = "_"

  ## Parses tags in the datadog statsd format
  ## http://docs.datadoghq.com/guides/dogstatsd/
  ## deprecated in 1.10; use datadog_extensions option instead
  parse_data_dog_tags = false

  ## Parses extensions to statsd in the datadog statsd format
  ## currently supports metrics and datadog tags.
  ## http://docs.datadoghq.com/guides/dogstatsd/
  datadog_extensions = false

  ## Parses distributions metric as specified in the datadog statsd format
  ## https://docs.datadoghq.com/developers/metrics/types/?tab=distribution#definition
  datadog_distributions = false

  ## Statsd data translation templates, more info can be read here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/TEMPLATE_PATTERN.md
  # templates = [
  #     "cpu.* measurement*"
  # ]

  ## Number of UDP messages allowed to queue up, once filled,
  ## the statsd server will start dropping packets
  allowed_pending_messages = 10000

  ## Number of timing/histogram values to track per-measurement in the
  ## calculation of percentiles. Raising this limit increases the accuracy
  ## of percentiles but also increases the memory usage and cpu time.
  percentile_limit = 1000

  ## Maximum socket buffer size in bytes, once the buffer fills up, metrics
  ## will start dropping.  Defaults to the OS default.
  # read_buffer_size = 65535

  ## Max duration (TTL) for each metric to stay cached/reported without being updated.
  # max_ttl = "10h"

  ## Sanitize name method
  ## By default, telegraf will pass names directly as they are received.
  ## However, upstream statsd now does sanitization of names which can be
  ## enabled by using the "upstream" method option. This option will a) replace
  ## white space with '_', replace '/' with '-', and remove charachters not
  ## matching 'a-zA-Z_\-0-9\.;='.
  #sanitize_name_method = ""
`
}
