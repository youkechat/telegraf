//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package cisco_telemetry_mdt

func (c *CiscoTelemetryMDT) SampleConfig() string {
	return `# Cisco model-driven telemetry (MDT) input plugin for IOS XR, IOS XE and NX-OS platforms
[[inputs.cisco_telemetry_mdt]]
 ## Telemetry transport can be "tcp" or "grpc".  TLS is only supported when
 ## using the grpc transport.
 transport = "grpc"

 ## Address and port to host telemetry listener
 service_address = ":57000"

 ## Grpc Maximum Message Size, default is 4MB, increase the size.
 max_msg_size = 4000000

 ## Enable TLS; grpc transport only.
 # tls_cert = "/etc/telegraf/cert.pem"
 # tls_key = "/etc/telegraf/key.pem"

 ## Enable TLS client authentication and define allowed CA certificates; grpc
 ##  transport only.
 # tls_allowed_cacerts = ["/etc/telegraf/clientca.pem"]

 ## Define (for certain nested telemetry measurements with embedded tags) which fields are tags
 # embedded_tags = ["Cisco-IOS-XR-qos-ma-oper:qos/interface-table/interface/input/service-policy-names/service-policy-instance/statistics/class-stats/class-name"]

 ## Define aliases to map telemetry encoding paths to simple measurement names
 [inputs.cisco_telemetry_mdt.aliases]
   ifstats = "ietf-interfaces:interfaces-state/interface/statistics"
 ## Define Property Xformation, please refer README and https://pubhub.devnetcloud.com/media/dme-docs-9-3-3/docs/appendix/ for Model details.
 [inputs.cisco_telemetry_mdt.dmes]
#    Global Property Xformation.
#    prop1 = "uint64 to int"
#    prop2 = "uint64 to string"
#    prop3 = "string to uint64"
#    prop4 = "string to int64"
#    prop5 = "string to float64"
#    auto-prop-xfrom = "auto-float-xfrom" #Xform any property which is string, and has float number to type float64
#    Per Path property xformation, Name is telemetry configuration under sensor-group, path configuration "WORD         Distinguished Name"
#    Per Path configuration is better as it avoid property collision issue of types.
#    dnpath = '{"Name": "show ip route summary","prop": [{"Key": "routes","Value": "string"}, {"Key": "best-paths","Value": "string"}]}'
#    dnpath2 = '{"Name": "show processes cpu","prop": [{"Key": "kernel_percent","Value": "float"}, {"Key": "idle_percent","Value": "float"}, {"Key": "process","Value": "string"}, {"Key": "user_percent","Value": "float"}, {"Key": "onesec","Value": "float"}]}'
#    dnpath3 = '{"Name": "show processes memory physical","prop": [{"Key": "processname","Value": "string"}]}'
`
}
