//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package icinga2

func (i *Icinga2) SampleConfig() string {
	return `# Gather Icinga2 status
[[inputs.icinga2]]
  ## Required Icinga2 server address
  # server = "https://localhost:5665"
  
  ## Required Icinga2 object type ("services" or "hosts")
  # object_type = "services"

  ## Credentials for basic HTTP authentication
  # username = "admin"
  # password = "admin"

  ## Maximum time to receive response.
  # response_timeout = "5s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = true
`
}
