//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package histogram

func (h *HistogramAggregator) SampleConfig() string {
	return `# Configuration for aggregate histogram metrics
[[aggregators.histogram]]
  ## The period in which to flush the aggregator.
  period = "30s"

  ## If true, the original metric will be dropped by the
  ## aggregator and will not get sent to the output plugins.
  drop_original = false

  ## If true, the histogram will be reset on flush instead
  ## of accumulating the results.
  reset = false

  ## Whether bucket values should be accumulated. If set to false, "gt" tag will be added.
  ## Defaults to true.
  cumulative = true

  ## Expiration interval for each histogram. The histogram will be expired if
  ## there are no changes in any buckets for this time interval. 0 == no expiration.
  # expiration_interval = "0m"

  ## If true, aggregated histogram are pushed to output only if it was updated since
  ## previous push. Defaults to false.
  # push_only_on_update = false

  ## Example config that aggregates all fields of the metric.
  # [[aggregators.histogram.config]]
  #   ## Right borders of buckets (with +Inf implicitly added).
  #   buckets = [0.0, 15.6, 34.5, 49.1, 71.5, 80.5, 94.5, 100.0]
  #   ## The name of metric.
  #   measurement_name = "cpu"

  ## Example config that aggregates only specific fields of the metric.
  # [[aggregators.histogram.config]]
  #   ## Right borders of buckets (with +Inf implicitly added).
  #   buckets = [0.0, 10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 100.0]
  #   ## The name of metric.
  #   measurement_name = "diskio"
  #   ## The concrete fields of metric
  #   fields = ["io_time", "read_time", "write_time"]
`
}
