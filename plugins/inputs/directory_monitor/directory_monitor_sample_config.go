//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package directory_monitor

func (monitor *DirectoryMonitor) SampleConfig() string {
	return `# Ingests files in a directory and then moves them to a target directory.
[[inputs.directory_monitor]]
  ## The directory to monitor and read files from.
  directory = ""
  #
  ## The directory to move finished files to.
  finished_directory = ""
  #
  ## The directory to move files to upon file error.
  ## If not provided, erroring files will stay in the monitored directory.
  # error_directory = ""
  #
  ## The amount of time a file is allowed to sit in the directory before it is picked up.
  ## This time can generally be low but if you choose to have a very large file written to the directory and it's potentially slow,
  ## set this higher so that the plugin will wait until the file is fully copied to the directory.
  # directory_duration_threshold = "50ms"
  #
  ## A list of the only file names to monitor, if necessary. Supports regex. If left blank, all files are ingested.
  # files_to_monitor = ["^.*\.csv"]
  #
  ## A list of files to ignore, if necessary. Supports regex.
  # files_to_ignore = [".DS_Store"]
  #
  ## Maximum lines of the file to process that have not yet be written by the
  ## output. For best throughput set to the size of the output's metric_buffer_limit.
  ## Warning: setting this number higher than the output's metric_buffer_limit can cause dropped metrics.
  # max_buffered_metrics = 10000
  #
  ## The maximum amount of file paths to queue up for processing at once, before waiting until files are processed to find more files.
  ## Lowering this value will result in *slightly* less memory use, with a potential sacrifice in speed efficiency, if absolutely necessary.
  # file_queue_size = 100000
  #
  ## Name a tag containing the name of the file the data was parsed from.  Leave empty
  ## to disable. Cautious when file name variation is high, this can increase the cardinality
  ## significantly. Read more about cardinality here:
  ## https://docs.influxdata.com/influxdb/cloud/reference/glossary/#series-cardinality
  # file_tag = ""
  #
  ## The dataformat to be read from the files.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
  ## NOTE: We currently only support parsing newline-delimited JSON. See the format here: https://github.com/ndjson/ndjson-spec
  data_format = "influx"
`
}
