//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package quantile

func (q *Quantile) SampleConfig() string {
	return `# Keep the aggregate quantiles of each metric passing through.
[[aggregators.quantile]]
  ## General Aggregator Arguments:
  ## The period on which to flush & clear the aggregator.
  period = "30s"

  ## If true, the original metric will be dropped by the
  ## aggregator and will not get sent to the output plugins.
  drop_original = false

  ## Quantiles to output in the range [0,1]
  # quantiles = [0.25, 0.5, 0.75]

  ## Type of aggregation algorithm
  ## Supported are:
  ##  "t-digest" -- approximation using centroids, can cope with large number of samples
  ##  "exact R7" -- exact computation also used by Excel or NumPy (Hyndman & Fan 1996 R7)
  ##  "exact R8" -- exact computation (Hyndman & Fan 1996 R8)
  ## NOTE: Do not use "exact" algorithms with large number of samples
  ##       to not impair performance or memory consumption!
  # algorithm = "t-digest"

  ## Compression for approximation (t-digest). The value needs to be
  ## greater or equal to 1.0. Smaller values will result in more
  ## performance but less accuracy.
  # compression = 100.0
`
}
