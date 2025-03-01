//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package event_hubs

func (e *EventHubs) SampleConfig() string {
	return `# Configuration for Event Hubs output plugin
[[outputs.event_hubs]]
  ## The full connection string to the Event Hub (required)
  ## The shared access key must have "Send" permissions on the target Event Hub.
  connection_string = "Endpoint=sb://namespace.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=superSecret1234=;EntityPath=hubName"
  ## Client timeout (defaults to 30s)
  # timeout = "30s"
  ## Data format to output.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  data_format = "json"
`
}
