//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package websocket

func (w *WebSocket) SampleConfig() string {
	return `# A plugin that can transmit metrics over WebSocket.
[[outputs.websocket]]
  ## URL is the address to send metrics to. Make sure ws or wss scheme is used.
  url = "ws://127.0.0.1:3000/telegraf"

  ## Timeouts (make sure read_timeout is larger than server ping interval or set to zero).
  # connect_timeout = "30s"
  # write_timeout = "30s"
  # read_timeout = "30s"

  ## Optionally turn on using text data frames (binary by default).
  # use_text_frames = false

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## Optional SOCKS5 proxy to use
  # socks5_enabled = true
  # socks5_address = "127.0.0.1:1080"
  # socks5_username = "alice"
  # socks5_password = "pass123"

  ## Data format to output.
  ## Each data format has it's own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  # data_format = "influx"

  ## Additional HTTP Upgrade headers
  # [outputs.websocket.headers]
  #   Authorization = "Bearer <TOKEN>"
`
}
