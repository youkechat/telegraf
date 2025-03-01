//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package prometheus

func (p *Prometheus) SampleConfig() string {
	return `# Read metrics from one or many prometheus clients
[[inputs.prometheus]]
  ## An array of urls to scrape metrics from.
  urls = ["http://localhost:9100/metrics"]
  
  ## Metric version controls the mapping from Prometheus metrics into
  ## Telegraf metrics.  When using the prometheus_client output, use the same
  ## value in both plugins to ensure metrics are round-tripped without
  ## modification.
  ##
  ##   example: metric_version = 1; 
  ##            metric_version = 2; recommended version
  # metric_version = 1
  
  ## Url tag name (tag containing scrapped url. optional, default is "url")
  # url_tag = "url"
  
  ## Whether the timestamp of the scraped metrics will be ignored.
  ## If set to true, the gather time will be used.
  # ignore_timestamp = false
  
  ## An array of Kubernetes services to scrape metrics from.
  # kubernetes_services = ["http://my-service-dns.my-namespace:9100/metrics"]
  
  ## Kubernetes config file to create client from.
  # kube_config = "/path/to/kubernetes.config"
  
  ## Scrape Kubernetes pods for the following prometheus annotations:
  ## - prometheus.io/scrape: Enable scraping for this pod
  ## - prometheus.io/scheme: If the metrics endpoint is secured then you will need to
  ##     set this to 'https' & most likely set the tls config.
  ## - prometheus.io/path: If the metrics path is not /metrics, define it with this annotation.
  ## - prometheus.io/port: If port is not 9102 use this annotation
  # monitor_kubernetes_pods = true
  
  ## Get the list of pods to scrape with either the scope of
  ## - cluster: the kubernetes watch api (default, no need to specify)
  ## - node: the local cadvisor api; for scalability. Note that the config node_ip or the environment variable NODE_IP must be set to the host IP.
  # pod_scrape_scope = "cluster"
  
  ## Only for node scrape scope: node IP of the node that telegraf is running on.
  ## Either this config or the environment variable NODE_IP must be set.
  # node_ip = "10.180.1.1"
 
  ## Only for node scrape scope: interval in seconds for how often to get updated pod list for scraping.
  ## Default is 60 seconds.
  # pod_scrape_interval = 60
  
  ## Restricts Kubernetes monitoring to a single namespace
  ##   ex: monitor_kubernetes_pods_namespace = "default"
  # monitor_kubernetes_pods_namespace = ""
  # label selector to target pods which have the label
  # kubernetes_label_selector = "env=dev,app=nginx"
  # field selector to target pods
  # eg. To scrape pods on a specific node
  # kubernetes_field_selector = "spec.nodeName=$HOSTNAME"

  # cache refresh interval to set the interval for re-sync of pods list. 
  # Default is 60 minutes.
  # cache_refresh_interval = 60

  ## Scrape Services available in Consul Catalog
  # [inputs.prometheus.consul]
  #   enabled = true
  #   agent = "http://localhost:8500"
  #   query_interval = "5m"

  #   [[inputs.prometheus.consul.query]]
  #     name = "a service name"
  #     tag = "a service tag"
  #     url = 'http://