//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package azure_data_explorer

func (adx *AzureDataExplorer) SampleConfig() string {
	return `# Sends metrics to Azure Data Explorer
[[outputs.azure_data_explorer]]
  ## The URI property of the Azure Data Explorer resource on Azure
  ## ex: endpoint_url = https://myadxresource.australiasoutheast.kusto.windows.net
  endpoint_url = ""

  ## The Azure Data Explorer database that the metrics will be ingested into.
  ## The plugin will NOT generate this database automatically, it's expected that this database already exists before ingestion.
  ## ex: "exampledatabase"
  database = ""

  ## Timeout for Azure Data Explorer operations
  # timeout = "20s"

  ## Type of metrics grouping used when pushing to Azure Data Explorer.
  ## Default is "TablePerMetric" for one table per different metric.
  ## For more information, please check the plugin README.
  # metrics_grouping_type = "TablePerMetric"

  ## Name of the single table to store all the metrics (Only needed if metrics_grouping_type is "SingleTable").
  # table_name = ""

  ## Creates tables and relevant mapping if set to true(default).
  ## Skips table and mapping creation if set to false, this is useful for running Telegraf with the lowest possible permissions i.e. table ingestor role.
  # create_tables = true
`
}
