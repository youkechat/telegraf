//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package influxdb_v2_listener

func (h *InfluxDBV2Listener) SampleConfig() string {
	return `# Accept metrics over InfluxDB 2.x HTTP API
[[inputs.influxdb_v2_listener]]
  ## Address and port to host InfluxDB listener on
  ## (Double check the port. Could be 9999 if using OSS Beta)
  service_address = ":8086"

  ## Maximum allowed HTTP request body size in bytes.
  ## 0 means to use the default of 32MiB.
  # max_body_size = "32MiB"

  ## Optional tag to determine the bucket.
  ## If the write has a bucket in the query string then it will be kept in this tag name.
  ## This tag can be used in downstream outputs.
  ## The default value of nothing means it will be off and the database will not be recorded.
  # bucket_tag = ""

  ## Set one or more allowed client CA certificate file names to
  ## enable mutually authenticated TLS connections
  # tls_allowed_cacerts = ["/etc/telegraf/clientca.pem"]

  ## Add service certificate and key
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"

  ## Optional token to accept for HTTP authentication.
  ## You probably want to make sure you have TLS configured above for this.
  # token = "some-long-shared-secret-token"

  ## Influx line protocol parser
  ## 'internal' is the default. 'upstream' is a newer parser that is faster
  ## and more memory efficient.
  # parser_type = "internal"
`
}
