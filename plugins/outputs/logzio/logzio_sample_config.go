//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package logzio

func (l *Logzio) SampleConfig() string {
	return `# A plugin that can send metrics over HTTPs to Logz.io
[[outputs.logzio]]
  ## Set to true if Logz.io sender checks the disk space before adding metrics to the disk queue.
  # check_disk_space = true

  ## The percent of used file system space at which the sender will stop queueing.
  ## When we will reach that percentage, the file system in which the queue is stored will drop
  ## all new logs until the percentage of used space drops below that threshold.
  # disk_threshold = 98

  ## How often Logz.io sender should drain the queue.
  ## Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
  # drain_duration = "3s"

  ## Where Logz.io sender should store the queue
  ## queue_dir = Sprintf("%s%s%s%s%d", os.TempDir(), string(os.PathSeparator),
  ##                     "logzio-buffer", string(os.PathSeparator), time.Now().UnixNano())

  ## Logz.io account token
  token = "your Logz.io token" # required

  ## Use your listener URL for your Logz.io account region.
  # url = "https://listener.logz.io:8071"
`
}
