//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package influxdb_listener

func (h *InfluxDBListener) SampleConfig() string {
	return `# Accept metrics over InfluxDB 1.x HTTP API
[[inputs.influxdb_listener]]
  ## Address and port to host HTTP listener on
  service_address = ":8186"

  ## maximum duration before timing out read of the request
  read_timeout = "10s"
  ## maximum duration before timing out write of the response
  write_timeout = "10s"

  ## Maximum allowed HTTP request body size in bytes.
  ## 0 means to use the default of 32MiB.
  max_body_size = 0

  ## Maximum line size allowed to be sent in bytes.
  ##   deprecated in 1.14; parser now handles lines of unlimited length and option is ignored
  # max_line_size = 0

  ## Set one or more allowed client CA certificate file names to
  ## enable mutually authenticated TLS connections
  tls_allowed_cacerts = ["/etc/telegraf/clientca.pem"]

  ## Add service certificate and key
  tls_cert = "/etc/telegraf/cert.pem"
  tls_key = "/etc/telegraf/key.pem"

  ## Optional tag name used to store the database name.
  ## If the write has a database in the query string then it will be kept in this tag name.
  ## This tag can be used in downstream outputs.
  ## The default value of nothing means it will be off and the database will not be recorded.
  ## If you have a tag that is the same as the one specified below, and supply a database,
  ## the tag will be overwritten with the database supplied.
  # database_tag = ""

  ## If set the retention policy specified in the write query will be added as
  ## the value of this tag name.
  # retention_policy_tag = ""

  ## Optional username and password to accept for HTTP basic authentication.
  ## You probably want to make sure you have TLS configured above for this.
  # basic_username = "foobar"
  # basic_password = "barfoo"

  ## Influx line protocol parser
  ## 'internal' is the default. 'upstream' is a newer parser that is faster
  ## and more memory efficient.
  # parser_type = "internal"
`
}
