//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package proxmox

func (px *Proxmox) SampleConfig() string {
	return `# Provides metrics from Proxmox nodes (Proxmox Virtual Environment > 6.2).
[[inputs.proxmox]]
  ## API connection configuration. The API token was introduced in Proxmox v6.2. Required permissions for user and token: PVEAuditor role on /.
  base_url = "https://localhost:8006/api2/json"
  api_token = "USER@REALM!TOKENID=UUID"
  ## Node name, defaults to OS hostname
  # node_name = ""

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  insecure_skip_verify = false

  # HTTP response timeout (default: 5s)
  response_timeout = "5s"
`
}
