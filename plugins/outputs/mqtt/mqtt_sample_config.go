//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package mqtt

func (m *MQTT) SampleConfig() string {
	return `# Configuration for MQTT server to send metrics to
[[outputs.mqtt]]
  ## MQTT Brokers
  ## The list of brokers should only include the hostname or IP address and the
  ## port to the broker. This should follow the format '{host}:{port}'. For
  ## example, "localhost:1883" or "127.0.0.1:8883".
  servers = ["localhost:1883"]

  ## MQTT Topic for Producer Messages
  ## MQTT outputs send metrics to this topic format:
  ## <topic_prefix>/<hostname>/<pluginname>/ (e.g. prefix/web01.example.com/mem)
  topic_prefix = "telegraf"

  ## QoS policy for messages
  ## The mqtt QoS policy for sending messages.
  ## See https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_9.0.0/com.ibm.mq.dev.doc/q029090_.htm
  ##   0 = at most once
  ##   1 = at least once
  ##   2 = exactly once
  # qos = 2

  ## Keep Alive
  ## Defines the maximum length of time that the broker and client may not
  ## communicate. Defaults to 0 which turns the feature off.
  ##
  ## For version v2.0.12 and later mosquitto there is a bug
  ## (see https://github.com/eclipse/mosquitto/issues/2117), which requires
  ## this to be non-zero. As a reference eclipse/paho.mqtt.golang defaults to 30.
  # keep_alive = 0

  ## username and password to connect MQTT server.
  # username = "telegraf"
  # password = "metricsmetricsmetricsmetrics"

  ## client ID
  ## The unique client id to connect MQTT server. If this parameter is not set
  ## then a random ID is generated.
  # client_id = ""

  ## Timeout for write operations. default: 5s
  # timeout = "5s"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"

  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## When true, metrics will be sent in one MQTT message per flush. Otherwise,
  ## metrics are written one metric per MQTT message.
  # batch = false

  ## When true, metric will have RETAIN flag set, making broker cache entries until someone
  ## actually reads it
  # retain = false

  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  data_format = "influx"
`
}
