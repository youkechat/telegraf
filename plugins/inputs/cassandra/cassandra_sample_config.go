//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package cassandra

func (c *Cassandra) SampleConfig() string {
	return `# Read Cassandra metrics through Jolokia
[[inputs.cassandra]]
  ## DEPRECATED: The cassandra plugin has been deprecated.  Please use the
  ## jolokia2 plugin instead.
  ##
  ## see https://github.com/influxdata/telegraf/tree/master/plugins/inputs/jolokia2

  context = "/jolokia/read"
  ## List of cassandra servers exposing jolokia read service
  servers = ["myuser:mypassword@10.10.10.1:8778","10.10.10.2:8778",":8778"]
  ## List of metrics collected on above servers
  ## Each metric consists of a jmx path.
  ## This will collect all heap memory usage metrics from the jvm and
  ## ReadLatency metrics for all keyspaces and tables.
  ## "type=Table" in the query works with Cassandra3.0. Older versions might
  ## need to use "type=ColumnFamily"
  metrics  = [
    "/java.lang:type=Memory/HeapMemoryUsage",
    "/org.apache.cassandra.metrics:type=Table,keyspace=*,scope=*,name=ReadLatency"
  ]
`
}
