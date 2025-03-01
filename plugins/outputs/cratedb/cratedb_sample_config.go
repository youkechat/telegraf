//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package cratedb

func (c *CrateDB) SampleConfig() string {
	return `# Configuration for CrateDB to send metrics to.
[[outputs.cratedb]]
  # A github.com/jackc/pgx/v4 connection string.
  # See https://pkg.go.dev/github.com/jackc/pgx/v4#ParseConfig
  url = "postgres://user:password@localhost/schema?sslmode=disable"
  # Timeout for all CrateDB queries.
  timeout = "5s"
  # Name of the table to store metrics in.
  table = "metrics"
  # If true, and the metrics table does not exist, create it automatically.
  table_create = true
  # The character(s) to replace any '.' in an object key with
  key_separator = "_"
`
}
