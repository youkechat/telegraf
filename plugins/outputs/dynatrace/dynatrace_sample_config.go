//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package dynatrace

func (d *Dynatrace) SampleConfig() string {
	return `# Send telegraf metrics to a Dynatrace environment
[[outputs.dynatrace]]
  ## For usage with the Dynatrace OneAgent you can omit any configuration,
  ## the only requirement is that the OneAgent is running on the same host.
  ## Only setup environment url and token if you want to monitor a Host without the OneAgent present.
  ##
  ## Your Dynatrace environment URL.
  ## For Dynatrace OneAgent you can leave this empty or set it to "http://127.0.0.1:14499/metrics/ingest" (default)
  ## For Dynatrace SaaS environments the URL scheme is "https://{your-environment-id}.live.dynatrace.com/api/v2/metrics/ingest"
  ## For Dynatrace Managed environments the URL scheme is "https://{your-domain}/e/{your-environment-id}/api/v2/metrics/ingest"
  url = ""

  ## Your Dynatrace API token.
  ## Create an API token within your Dynatrace environment, by navigating to Settings > Integration > Dynatrace API
  ## The API token needs data ingest scope permission. When using OneAgent, no API token is required.
  api_token = ""

  ## Optional prefix for metric names (e.g.: "telegraf")
  prefix = "telegraf"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Optional flag for ignoring tls certificate check
  # insecure_skip_verify = false

  ## Connection timeout, defaults to "5s" if not set.
  timeout = "5s"

  ## If you want metrics to be treated and reported as delta counters, add the metric names here
  additional_counters = [ ]

  ## Optional dimensions to be added to every metric
  # [outputs.dynatrace.default_dimensions]
  # default_key = "default value"
`
}
