//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package github

func (g *GitHub) SampleConfig() string {
	return `# Gather repository information from GitHub hosted repositories.
[[inputs.github]]
  ## List of repositories to monitor
  repositories = [
    "influxdata/telegraf",
    "influxdata/influxdb"
  ]

  ## Github API access token.  Unauthenticated requests are limited to 60 per hour.
  # access_token = ""

  ## Github API enterprise url. Github Enterprise accounts must specify their base url.
  # enterprise_base_url = ""

  ## Timeout for HTTP requests.
  # http_timeout = "5s"

  ## List of additional fields to query.
  ## NOTE: Getting those fields might involve issuing additional API-calls, so please
  ##       make sure you do not exceed the rate-limit of GitHub.
  ##
  ## Available fields are:
  ##  - pull-requests -- number of open and closed pull requests (2 API-calls per repository)
  # additional_fields = []
`
}
