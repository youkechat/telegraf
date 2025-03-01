//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package yandex_cloud_monitoring

func (a *YandexCloudMonitoring) SampleConfig() string {
	return `# Send aggregated metrics to Yandex.Cloud Monitoring
[[outputs.yandex_cloud_monitoring]]
  ## Timeout for HTTP writes.
  # timeout = "20s"

  ## Yandex.Cloud monitoring API endpoint. Normally should not be changed
  # endpoint_url = "https://monitoring.api.cloud.yandex.net/monitoring/v2/data/write"

  ## All user metrics should be sent with "custom" service specified. Normally should not be changed
  # service = "custom"
`
}
