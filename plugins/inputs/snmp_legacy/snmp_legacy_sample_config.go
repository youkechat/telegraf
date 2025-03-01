//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package snmp_legacy

func (s *Snmp) SampleConfig() string {
	return `# DEPRECATED! PLEASE USE inputs.snmp INSTEAD.
[[inputs.snmp_legacy]]
  ## Use 'oids.txt' file to translate oids to names
  ## To generate 'oids.txt' you need to run:
  ##   snmptranslate -m all -Tz -On | sed -e 's/"//g' > /tmp/oids.txt
  ## Or if you have an other MIB folder with custom MIBs
  ##   snmptranslate -M /mycustommibfolder -Tz -On -m all | sed -e 's/"//g' > oids.txt
  snmptranslate_file = "/tmp/oids.txt"
  [[inputs.snmp.host]]
    address = "192.168.2.2:161"
    # SNMP community
    community = "public" # default public
    # SNMP version (1, 2 or 3)
    # Version 3 not supported yet
    version = 2 # default 2
    # SNMP response timeout
    timeout = 2.0 # default 2.0
    # SNMP request retries
    retries = 2 # default 2
    # Which get/bulk do you want to collect for this host
    collect = ["mybulk", "sysservices", "sysdescr"]
    # Simple list of OIDs to get, in addition to "collect"
    get_oids = []
  [[inputs.snmp.host]]
    address = "192.168.2.3:161"
    community = "public"
    version = 2
    timeout = 2.0
    retries = 2
    collect = ["mybulk"]
    get_oids = [
        "ifNumber",
        ".1.3.6.1.2.1.1.3.0",
    ]
  [[inputs.snmp.get]]
    name = "ifnumber"
    oid = "ifNumber"
  [[inputs.snmp.get]]
    name = "interface_speed"
    oid = "ifSpeed"
    instance = "0"
  [[inputs.snmp.get]]
    name = "sysuptime"
    oid = ".1.3.6.1.2.1.1.3.0"
    unit = "second"
  [[inputs.snmp.bulk]]
    name = "mybulk"
    max_repetition = 127
    oid = ".1.3.6.1.2.1.1"
  [[inputs.snmp.bulk]]
    name = "ifoutoctets"
    max_repetition = 127
    oid = "ifOutOctets"
  [[inputs.snmp.host]]
    address = "192.168.2.13:161"
    #address = "127.0.0.1:161"
    community = "public"
    version = 2
    timeout = 2.0
    retries = 2
    #collect = ["mybulk", "sysservices", "sysdescr", "systype"]
    collect = ["sysuptime" ]
    [[inputs.snmp.host.table]]
      name = "iftable3"
      include_instances = ["enp5s0", "eth1"]
  # SNMP TABLEs
  # table without mapping neither subtables
  [[inputs.snmp.table]]
    name = "iftable1"
    oid = ".1.3.6.1.2.1.31.1.1.1"
  # table without mapping but with subtables
  [[inputs.snmp.table]]
    name = "iftable2"
    oid = ".1.3.6.1.2.1.31.1.1.1"
    sub_tables = [".1.3.6.1.2.1.2.2.1.13"]
  # table with mapping but without subtables
  [[inputs.snmp.table]]
    name = "iftable3"
    oid = ".1.3.6.1.2.1.31.1.1.1"
    # if empty. get all instances
    mapping_table = ".1.3.6.1.2.1.31.1.1.1.1"
    # if empty, get all subtables
  # table with both mapping and subtables
  [[inputs.snmp.table]]
    name = "iftable4"
    oid = ".1.3.6.1.2.1.31.1.1.1"
    # if empty get all instances
    mapping_table = ".1.3.6.1.2.1.31.1.1.1.1"
    # if empty get all subtables
    # sub_tables could be not "real subtables"
    sub_tables=[".1.3.6.1.2.1.2.2.1.13", "bytes_recv", "bytes_send"]
`
}
