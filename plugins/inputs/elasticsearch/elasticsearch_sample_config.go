//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package elasticsearch

func (e *Elasticsearch) SampleConfig() string {
	return `# Read stats from one or more Elasticsearch servers or clusters
[[inputs.elasticsearch]]
  ## specify a list of one or more Elasticsearch servers
  ## you can add username and password to your url to use basic authentication:
  ## servers = ["http://user:pass@localhost:9200"]
  servers = ["http://localhost:9200"]

  ## Timeout for HTTP requests to the elastic search server(s)
  http_timeout = "5s"

  ## When local is true (the default), the node will read only its own stats.
  ## Set local to false when you want to read the node stats from all nodes
  ## of the cluster.
  local = true

  ## Set cluster_health to true when you want to obtain cluster health stats
  cluster_health = false

  ## Adjust cluster_health_level when you want to obtain detailed health stats
  ## The options are
  ##  - indices (default)
  ##  - cluster
  # cluster_health_level = "indices"

  ## Set cluster_stats to true when you want to obtain cluster stats.
  cluster_stats = false

  ## Only gather cluster_stats from the master node. To work this require local = true
  cluster_stats_only_from_master = true

  ## Indices to collect; can be one or more indices names or _all
  ## Use of wildcards is allowed. Use a wildcard at the end to retrieve index names that end with a changing value, like a date.
  indices_include = ["_all"]

  ## One of "shards", "cluster", "indices"
  ## Currently only "shards" is implemented
  indices_level = "shards"

  ## node_stats is a list of sub-stats that you want to have gathered. Valid options
  ## are "indices", "os", "process", "jvm", "thread_pool", "fs", "transport", "http",
  ## "breaker". Per default, all stats are gathered.
  # node_stats = ["jvm", "http"]

  ## HTTP Basic Authentication username and password.
  # username = ""
  # password = ""

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## Sets the number of most recent indices to return for indices that are configured with a date-stamped suffix.
  ## Each 'indices_include' entry ending with a wildcard (*) or glob matching pattern will group together all indices that match it, and 
  ## sort them by the date or number after the wildcard. Metrics then are gathered for only the 'num_most_recent_indices' amount of most 
  ## recent indices.
  # num_most_recent_indices = 0
`
}
