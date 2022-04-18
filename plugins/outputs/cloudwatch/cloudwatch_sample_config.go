//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package cloudwatch

func (c *CloudWatch) SampleConfig() string {
	return `# Configuration for AWS CloudWatch output.
[[outputs.cloudwatch]]
  ## Amazon REGION
  region = "us-east-1"

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

  ## Endpoint to make request against, the correct endpoint is automatically
  ## determined and this option should only be set if you wish to override the
  ## default.
  ##   ex: endpoint_url = "http://localhost:8000"
  # endpoint_url = ""

  ## Namespace for the CloudWatch MetricDatums
  namespace = "InfluxData/Telegraf"

  ## If you have a large amount of metrics, you should consider to send statistic
  ## values instead of raw metrics which could not only improve performance but
  ## also save AWS API cost. If enable this flag, this plugin would parse the required
  ## CloudWatch statistic fields (count, min, max, and sum) and send them to CloudWatch.
  ## You could use basicstats aggregator to calculate those fields. If not all statistic
  ## fields are available, all fields would still be sent as raw metrics.
  # write_statistics = false

  ## Enable high resolution metrics of 1 second (if not enabled, standard resolution are of 60 seconds precision)
  # high_resolution_metrics = false
`
}
