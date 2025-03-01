//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package mongodb

func (s *MongoDB) SampleConfig() string {
	return `# A plugin that can transmit logs to mongodb
[[outputs.mongodb]]
  # connection string examples for mongodb
  dsn = "mongodb://localhost:27017"
  # dsn = "mongodb://mongod1:27017,mongod2:27017,mongod3:27017/admin&replicaSet=myReplSet&w=1"

  # overrides serverSelectionTimeoutMS in dsn if set
  # timeout = "30s"

  # default authentication, optional
  # authentication = "NONE"

  # for SCRAM-SHA-256 authentication
  # authentication = "SCRAM"
  # username = "root"
  # password = "***"

  # for x509 certificate authentication
  # authentication = "X509"
  # tls_ca = "ca.pem"
  # tls_key = "client.pem"
  # # tls_key_pwd = "changeme" # required for encrypted tls_key
  # insecure_skip_verify = false

  # database to store measurements and time series collections
  # database = "telegraf"

  # granularity can be seconds, minutes, or hours.
  # configuring this value will be based on your input collection frequency.
  # see https://docs.mongodb.com/manual/core/timeseries-collections/#create-a-time-series-collection
  # granularity = "seconds"

  # optionally set a TTL to automatically expire documents from the measurement collections.
  # ttl = "360h"
`
}
