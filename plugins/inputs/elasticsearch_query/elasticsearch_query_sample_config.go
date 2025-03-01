//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package elasticsearch_query

func (e *ElasticsearchQuery) SampleConfig() string {
	return `# Derive metrics from aggregating Elasticsearch query results
[[inputs.elasticsearch_query]]
  ## The full HTTP endpoint URL for your Elasticsearch instance
  ## Multiple urls can be specified as part of the same cluster,
  ## this means that only ONE of the urls will be written to each interval.
  urls = [ "http://node1.es.example.com:9200" ] # required.

  ## Elasticsearch client timeout, defaults to "5s".
  # timeout = "5s"

  ## Set to true to ask Elasticsearch a list of all cluster nodes,
  ## thus it is not necessary to list all nodes in the urls config option
  # enable_sniffer = false

  ## Set the interval to check if the Elasticsearch nodes are available
  ## This option is only used if enable_sniffer is also set (0s to disable it)
  # health_check_interval = "10s"

  ## HTTP basic authentication details (eg. when using x-pack)
  # username = "telegraf"
  # password = "mypassword"

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  [[inputs.elasticsearch_query.aggregation]]
    ## measurement name for the results of the aggregation query
    measurement_name = "measurement"

    ## Elasticsearch indexes to query (accept wildcards).
    index = "index-*"

    ## The date/time field in the Elasticsearch index (mandatory).
    date_field = "@timestamp"

    ## If the field used for the date/time field in Elasticsearch is also using
    ## a custom date/time format it may be required to provide the format to
    ## correctly parse the field.
    ##
    ## If using one of the built in elasticsearch formats this is not required.
    # date_field_custom_format = ""

    ## Time window to query (eg. "1m" to query documents from last minute).
    ## Normally should be set to same as collection interval
    query_period = "1m"

    ## Lucene query to filter results
    # filter_query = "*"

    ## Fields to aggregate values (must be numeric fields)
    # metric_fields = ["metric"]

    ## Aggregation function to use on the metric fields
    ## Must be set if 'metric_fields' is set
    ## Valid values are: avg, sum, min, max, sum
    # metric_function = "avg"

    ## Fields to be used as tags
    ## Must be text, non-analyzed fields. Metric aggregations are performed per tag
    # tags = ["field.keyword", "field2.keyword"]

    ## Set to true to not ignore documents when the tag(s) above are missing
    # include_missing_tag = false

    ## String value of the tag when the tag does not exist
    ## Used when include_missing_tag is true
    # missing_tag_value = "null"
`
}
