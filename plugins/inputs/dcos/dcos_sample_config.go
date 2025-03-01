//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package dcos

func (d *DCOS) SampleConfig() string {
	return `# Input plugin for DC/OS metrics
[[inputs.dcos]]
  ## The DC/OS cluster URL.
  cluster_url = "https://dcos-master-1"

  ## The ID of the service account.
  service_account_id = "telegraf"
  ## The private key file for the service account.
  service_account_private_key = "/etc/telegraf/telegraf-sa-key.pem"

  ## Path containing login token.  If set, will read on every gather.
  # token_file = "/home/dcos/.dcos/token"

  ## In all filter options if both include and exclude are empty all items
  ## will be collected.  Arrays may contain glob patterns.
  ##
  ## Node IDs to collect metrics from.  If a node is excluded, no metrics will
  ## be collected for its containers or apps.
  # node_include = []
  # node_exclude = []
  ## Container IDs to collect container metrics from.
  # container_include = []
  # container_exclude = []
  ## Container IDs to collect app metrics from.
  # app_include = []
  # app_exclude = []

  ## Maximum concurrent connections to the cluster.
  # max_connections = 10
  ## Maximum time to receive a response from cluster.
  # response_timeout = "20s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## If false, skip chain & host verification
  # insecure_skip_verify = true

  ## Recommended filtering to reduce series cardinality.
  # [inputs.dcos.tagdrop]
  #   path = ["/var/lib/mesos/slave/slaves/*"]
`
}
