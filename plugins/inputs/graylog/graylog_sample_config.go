//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package graylog

func (h *GrayLog) SampleConfig() string {
	return `# Read flattened metrics from one or more GrayLog HTTP endpoints
[[inputs.graylog]]
  ## API endpoint, currently supported API:
  ##
  ##   - multiple  (e.g. http://<host>:9000/api/system/metrics/multiple)
  ##   - namespace (e.g. http://<host>:9000/api/system/metrics/namespace/{namespace})
  ##
  ## For namespace endpoint, the metrics array will be ignored for that call.
  ## Endpoint can contain namespace and multiple type calls.
  ##
  ## Please check http://[graylog-server-ip]:9000/api/api-browser for full list
  ## of endpoints
  servers = [
    "http://[graylog-server-ip]:9000/api/system/metrics/multiple",
  ]

  ## Set timeout (default 5 seconds)
  # timeout = "5s"

  ## Metrics list
  ## List of metrics can be found on Graylog webservice documentation.
  ## Or by hitting the web service api at:
  ##   http://[graylog-host]:9000/api/system/metrics
  metrics = [
    "jvm.cl.loaded",
    "jvm.memory.pools.Metaspace.committed"
  ]

  ## Username and password
  username = ""
  password = ""

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false
`
}
