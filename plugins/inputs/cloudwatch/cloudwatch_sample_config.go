//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package cloudwatch

func (c *CloudWatch) SampleConfig() string {
	return `# Pull Metric Statistics from Amazon CloudWatch
[[inputs.cloudwatch]]
  ## Amazon Region
  region = "us-east-1"

  ## Amazon Credentials
  ## Credentials are loaded in the following order
  ## 1) Web identity provider credentials via STS if role_arn and web_identity_token_file are specified
  ## 2) Assumed credentials via STS if role_arn is specified
  ## 3) explicit credentials from 'access_key' and 'secret_key'
  ## 4) shared profile from 'profile'
  ## 5) environment variables
  ## 6) shared credentials file
  ## 7) EC2 Instance Profile
  # access_key = ""
  # secret_key = ""
  # token = ""
  # role_arn = ""
  # web_identity_token_file = ""
  # role_session_name = ""
  # profile = ""
  # shared_credential_file = ""

  ## Endpoint to make request against, the correct endpoint is automatically
  ## determined and this option should only be set if you wish to override the
  ## default.
  ##   ex: endpoint_url = "http://localhost:8000"
  # endpoint_url = ""

  ## Set http_proxy (telegraf uses the system wide proxy settings if it's is not set)
  # http_proxy_url = "http://localhost:8888"

  # The minimum period for Cloudwatch metrics is 1 minute (60s). However not all
  # metrics are made available to the 1 minute period. Some are collected at
  # 3 minute, 5 minute, or larger intervals. See https://aws.amazon.com/cloudwatch/faqs/#monitoring.
  # Note that if a period is configured that is smaller than the minimum for a
  # particular metric, that metric will not be returned by the Cloudwatch API
  # and will not be collected by Telegraf.
  #
  ## Requested CloudWatch aggregation Period (required - must be a multiple of 60s)
  period = "5m"

  ## Collection Delay (required - must account for metrics availability via CloudWatch API)
  delay = "5m"

  ## Recommended: use metric 'interval' that is a multiple of 'period' to avoid
  ## gaps or overlap in pulled data
  interval = "5m"

  ## Recommended if "delay" and "period" are both within 3 hours of request time. Invalid values will be ignored.
  ## Recently Active feature will only poll for CloudWatch ListMetrics values that occurred within the last 3 Hours.
  ## If enabled, it will reduce total API usage of the CloudWatch ListMetrics API and require less memory to retain.
  ## Do not enable if "period" or "delay" is longer than 3 hours, as it will not return data more than 3 hours old.
  ## See https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html
  #recently_active = "PT3H"

  ## Configure the TTL for the internal cache of metrics.
  # cache_ttl = "1h"

  ## Metric Statistic Namespaces (required)
  namespaces = ["AWS/ELB"]
  # A single metric statistic namespace that will be appended to namespaces on startup
  # namespace = "AWS/ELB"

  ## Maximum requests per second. Note that the global default AWS rate limit is
  ## 50 reqs/sec, so if you define multiple namespaces, these should add up to a
  ## maximum of 50.
  ## See http://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_limits.html
  # ratelimit = 25

  ## Timeout for http requests made by the cloudwatch client.
  # timeout = "5s"

  ## Namespace-wide statistic filters. These allow fewer queries to be made to
  ## cloudwatch.
  # statistic_include = [ "average", "sum", "minimum", "maximum", sample_count" ]
  # statistic_exclude = []

  ## Metrics to Pull
  ## Defaults to all Metrics in Namespace if nothing is provided
  ## Refreshes Namespace available metrics every 1h
  #[[inputs.cloudwatch.metrics]]
  #  names = ["Latency", "RequestCount"]
  #
  #  ## Statistic filters for Metric.  These allow for retrieving specific
  #  ## statistics for an individual metric.
  #  # statistic_include = [ "average", "sum", "minimum", "maximum", sample_count" ]
  #  # statistic_exclude = []
  #
  #  ## Dimension filters for Metric.  All dimensions defined for the metric names
  #  ## must be specified in order to retrieve the metric statistics.
  #  ## 'value' has wildcard / 'glob' matching support such as 'p-*'.
  #  [[inputs.cloudwatch.metrics.dimensions]]
  #    name = "LoadBalancerName"
  #    value = "p-example"
`
}
