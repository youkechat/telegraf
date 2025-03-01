//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package beat

func (beat *Beat) SampleConfig() string {
	return `# Read metrics exposed by Beat
[[inputs.beat]]
  ## An URL from which to read Beat-formatted JSON
  ## Default is "http://127.0.0.1:5066".
  url = "http://127.0.0.1:5066"

  ## Enable collection of the listed stats
  ## An empty list means collect all. Available options are currently
  ## "beat", "libbeat", "system" and "filebeat".
  # include = ["beat", "libbeat", "filebeat"]

  ## HTTP method
  # method = "GET"

  ## Optional HTTP headers
  # headers = {"X-Special-Header" = "Special-Value"}

  ## Override HTTP "Host" header
  # host_header = "logstash.example.com"

  ## Timeout for HTTP requests
  # timeout = "5s"

  ## Optional HTTP Basic Auth credentials
  # username = "username"
  # password = "pa$$word"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}
