//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package postgresql_extensible

func (p *Postgresql) SampleConfig() string {
	return `# Read metrics from one or many postgresql servers
[[inputs.postgresql_extensible]]
  # specify address via a url matching:
  # postgres://[pqgotest[:password]]@host:port[/dbname]?sslmode=...
  # or a simple string:
  #   host=localhost port=5432 user=pqgotest password=... sslmode=... dbname=app_production
  #
  # All connection parameters are optional.
  # Without the dbname parameter, the driver will default to a database
  # with the same name as the user. This dbname is just for instantiating a
  # connection with the server and doesn't restrict the databases we are trying
  # to grab metrics for.
  #
  address = "host=localhost user=postgres sslmode=disable"
  # A list of databases to pull metrics about. If not specified, metrics for all
  # databases are gathered.
  # databases = ["app_production", "testing"]

  ## Whether to use prepared statements when connecting to the database.
  ## This should be set to false when connecting through a PgBouncer instance
  ## with pool_mode set to transaction.
  prepared_statements = true

  # Define the toml config where the sql queries are stored
  # New queries can be added, if the withdbname is set to true and there is no
  # databases defined in the 'databases field', the sql query is ended by a 'is
  # not null' in order to make the query succeed.
  # Be careful that the sqlquery must contain the where clause with a part of
  # the filtering, the plugin will add a 'IN (dbname list)' clause if the
  # withdbname is set to true
  # Example :
  # The sqlquery : "SELECT * FROM pg_stat_database where datname" become
  # "SELECT * FROM pg_stat_database where datname IN ('postgres', 'pgbench')"
  # because the databases variable was set to ['postgres', 'pgbench' ] and the
  # withdbname was true.
  # Be careful that if the withdbname is set to false you don't have to define
  # the where clause (aka with the dbname)
  #
  # The script option can be used to specify the .sql file path.
  # If script and sqlquery options specified at same time, sqlquery will be used
  #
  # the tagvalue field is used to define custom tags (separated by comas).
  # the query is expected to return columns which match the names of the
  # defined tags. The values in these columns must be of a string-type,
  # a number-type or a blob-type.
  #
  # The timestamp field is used to override the data points timestamp value. By
  # default, all rows inserted with current time. By setting a timestamp column,
  # the row will be inserted with that column's value.
  #
  # Structure :
  # [[inputs.postgresql_extensible.query]]
  #   sqlquery string
  #   version string
  #   withdbname boolean
  #   tagvalue string (coma separated)
  #   timestamp string
  [[inputs.postgresql_extensible.query]]
    sqlquery="SELECT * FROM pg_stat_database where datname"
    version=901
    withdbname=false
    tagvalue=""
  [[inputs.postgresql_extensible.query]]
    script="your_sql-filepath.sql"
    version=901
    withdbname=false
    tagvalue=""
`
}
