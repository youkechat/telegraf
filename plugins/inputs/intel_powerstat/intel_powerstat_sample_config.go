//go:build linux
// +build linux

//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package intel_powerstat

func (p *PowerStat) SampleConfig() string {
	return `# Intel PowerStat plugin enables monitoring of platform metrics (power, TDP) and per-CPU metrics like temperature, power and utilization.
[[inputs.intel_powerstat]]
  ## All global metrics are always collected by Intel PowerStat plugin.
  ## User can choose which per-CPU metrics are monitored by the plugin in cpu_metrics array.
  ## Empty array means no per-CPU specific metrics will be collected by the plugin - in this case only platform level
  ## telemetry will be exposed by Intel PowerStat plugin.
  ## Supported options:
  ## "cpu_frequency", "cpu_busy_frequency", "cpu_temperature", "cpu_c1_state_residency", "cpu_c6_state_residency", "cpu_busy_cycles"
  # cpu_metrics = []
`
}
