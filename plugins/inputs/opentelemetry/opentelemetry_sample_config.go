//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package opentelemetry

func (o *OpenTelemetry) SampleConfig() string {
	return `# Receive OpenTelemetry traces, metrics, and logs over gRPC
[[inputs.opentelemetry]]
  ## Override the default (0.0.0.0:4317) destination OpenTelemetry gRPC service
  ## address:port
  # service_address = "0.0.0.0:4317"

  ## Override the default (5s) new connection timeout
  # timeout = "5s"

  ## Override the default (prometheus-v1) metrics schema.
  ## Supports: "prometheus-v1", "prometheus-v2"
  ## For more information about the alternatives, read the Prometheus input
  ## plugin notes.
  # metrics_schema = "prometheus-v1"

  ## Optional TLS Config.
  ## For advanced options: https://github.com/influxdata/telegraf/blob/v1.18.3/docs/TLS.md
  ##
  ## Set one or more allowed client CA certificate file names to
  ## enable mutually authenticated TLS connections.
  # tls_allowed_cacerts = ["/etc/telegraf/clientca.pem"]
  ## Add service certificate and key.
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
`
}
