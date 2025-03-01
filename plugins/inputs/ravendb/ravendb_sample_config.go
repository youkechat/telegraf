//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package ravendb

func (r *RavenDB) SampleConfig() string {
	return `# Reads metrics from RavenDB servers via the Monitoring Endpoints
[[inputs.ravendb]]
  ## Node URL and port that RavenDB is listening on. By default,
  ## attempts to connect securely over HTTPS, however, if the user
  ## is running a local unsecure development cluster users can use
  ## HTTP via a URL like "http://localhost:8080"
  url = "https://localhost:4433"

  ## RavenDB X509 client certificate setup
  # tls_cert = "/etc/telegraf/raven.crt"
  # tls_key = "/etc/telegraf/raven.key"

  ## Optional request timeout
  ##
  ## Timeout, specifies the amount of time to wait
  ## for a server's response headers after fully writing the request and
  ## time limit for requests made by this client
  # timeout = "5s"

  ## List of statistics which are collected
  # At least one is required
  # Allowed values: server, databases, indexes, collections
  #
  # stats_include = ["server", "databases", "indexes", "collections"]

  ## List of db where database stats are collected
  ## If empty, all db are concerned
  # db_stats_dbs = []

  ## List of db where index status are collected
  ## If empty, all indexes from all db are concerned
  # index_stats_dbs = []

  ## List of db where collection status are collected
  ## If empty, all collections from all db are concerned
  # collection_stats_dbs = []
`
}
