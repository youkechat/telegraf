//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package http

func (h *HTTP) SampleConfig() string {
	return `# A plugin that can transmit metrics over HTTP
[[outputs.http]]
  ## URL is the address to send metrics to
  url = "http://127.0.0.1:8080/telegraf"

  ## Timeout for HTTP message
  # timeout = "5s"

  ## HTTP method, one of: "POST" or "PUT"
  # method = "POST"

  ## HTTP Basic Auth credentials
  # username = "username"
  # password = "pa$$word"

  ## OAuth2 Client Credentials Grant
  # client_id = "clientid"
  # client_secret = "secret"
  # token_url = "https://indentityprovider/oauth2/v1/token"
  # scopes = ["urn:opc:idm:__myscopes__"]

  ## Optional TLS Config
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## Optional Cookie authentication
  # cookie_auth_url = "https://localhost/authMe"
  # cookie_auth_method = "POST"
  # cookie_auth_username = "username"
  # cookie_auth_password = "pa$$word"
  # cookie_auth_headers = '{"Content-Type": "application/json", "X-MY-HEADER":"hello"}'
  # cookie_auth_body = '{"username": "user", "password": "pa$$word", "authenticate": "me"}'
  ## cookie_auth_renewal not set or set to "0" will auth once and never renew the cookie
  # cookie_auth_renewal = "5m"

  ## Data format to output.
  ## Each data format has it's own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  # data_format = "influx"

  ## Use batch serialization format (default) instead of line based format.
  ## Batch format is more efficient and should be used unless line based
  ## format is really needed.
  # use_batch_format = true

  ## HTTP Content-Encoding for write request body, can be set to "gzip" to
  ## compress body or "identity" to apply no encoding.
  # content_encoding = "identity"

  ## Additional HTTP headers
  # [outputs.http.headers]
  #   # Should be set manually to "application/json" for json data_format
  #   Content-Type = "text/plain; charset=utf-8"

  ## Idle (keep-alive) connection timeout.
  ## Maximum amount of time before idle connection is closed.
  ## Zero means no limit.
  # idle_conn_timeout = 0

  ## Amazon Region
  #region = "us-east-1"

  ## Amazon Credentials
  ## Credentials are loaded in the following order
  ## 1) Web identity provider credentials via STS if role_arn and web_identity_token_file are specified
  ## 2) Assumed credentials via STS if role_arn is specified
  ## 3) explicit credentials from 'access_key' and 'secret_key'
  ## 4) shared profile from 'profile'
  ## 5) environment variables
  ## 6) shared credentials file
  ## 7) EC2 Instance Profile
  #access_key = ""
  #secret_key = ""
  #token = ""
  #role_arn = ""
  #web_identity_token_file = ""
  #role_session_name = ""
  #profile = ""
  #shared_credential_file = ""

  ## Optional list of statuscodes (<200 or >300) upon which requests should not be retried
  # non_retryable_statuscodes = [409, 413]
`
}
