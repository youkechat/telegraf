//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package azure_monitor

func (a *AzureMonitor) SampleConfig() string {
	return `# Send aggregate metrics to Azure Monitor
[[outputs.azure_monitor]]
  ## Timeout for HTTP writes.
  # timeout = "20s"

  ## Set the namespace prefix, defaults to "Telegraf/<input-name>".
  # namespace_prefix = "Telegraf/"

  ## Azure Monitor doesn't have a string value type, so convert string
  ## fields to dimensions (a.k.a. tags) if enabled. Azure Monitor allows
  ## a maximum of 10 dimensions so Telegraf will only send the first 10
  ## alphanumeric dimensions.
  # strings_as_dimensions = false

  ## Both region and resource_id must be set or be available via the
  ## Instance Metadata service on Azure Virtual Machines.
  #
  ## Azure Region to publish metrics against.
  ##   ex: region = "southcentralus"
  # region = ""
  #
  ## The Azure Resource ID against which metric will be logged, e.g.
  ##   ex: resource_id = "/subscriptions/<subscription_id>/resourceGroups/<resource_group>/providers/Microsoft.Compute/virtualMachines/<vm_name>"
  # resource_id = ""

  ## Optionally, if in Azure US Government, China, or other sovereign
  ## cloud environment, set the appropriate REST endpoint for receiving
  ## metrics. (Note: region may be unused in this context)
  # endpoint_url = "https://monitoring.core.usgovcloudapi.net"
`
}
