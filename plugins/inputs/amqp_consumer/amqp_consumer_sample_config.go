//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package amqp_consumer

func (a *AMQPConsumer) SampleConfig() string {
	return `# AMQP consumer plugin
[[inputs.amqp_consumer]]
  ## Brokers to consume from.  If multiple brokers are specified a random broker
  ## will be selected anytime a connection is established.  This can be
  ## helpful for load balancing when not using a dedicated load balancer.
  brokers = ["amqp://localhost:5672/influxdb"]

  ## Authentication credentials for the PLAIN auth_method.
  # username = ""
  # password = ""

  ## Name of the exchange to declare.  If unset, no exchange will be declared.
  exchange = "telegraf"

  ## Exchange type; common types are "direct", "fanout", "topic", "header", "x-consistent-hash".
  # exchange_type = "topic"

  ## If true, exchange will be passively declared.
  # exchange_passive = false

  ## Exchange durability can be either "transient" or "durable".
  # exchange_durability = "durable"

  ## Additional exchange arguments.
  # exchange_arguments = { }
  # exchange_arguments = {"hash_property" = "timestamp"}

  ## AMQP queue name.
  queue = "telegraf"

  ## AMQP queue durability can be "transient" or "durable".
  queue_durability = "durable"

  ## If true, queue will be passively declared.
  # queue_passive = false

  ## A binding between the exchange and queue using this binding key is
  ## created.  If unset, no binding is created.
  binding_key = "#"

  ## Maximum number of messages server should give to the worker.
  # prefetch_count = 50

  ## Maximum messages to read from the broker that have not been written by an
  ## output.  For best throughput set based on the number of metrics within
  ## each message and the size of the output's metric_batch_size.
  ##
  ## For example, if each message from the queue contains 10 metrics and the
  ## output metric_batch_size is 1000, setting this to 100 will ensure that a
  ## full batch is collected and the write is triggered immediately without
  ## waiting until the next flush_interval.
  # max_undelivered_messages = 1000

  ## Auth method. PLAIN and EXTERNAL are supported
  ## Using EXTERNAL requires enabling the rabbitmq_auth_mechanism_ssl plugin as
  ## described here: https://www.rabbitmq.com/plugins.html
  # auth_method = "PLAIN"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## Content encoding for message payloads, can be set to "gzip" to or
  ## "identity" to apply no encoding.
  # content_encoding = "identity"

  ## Data format to consume.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
  data_format = "influx"
`
}
