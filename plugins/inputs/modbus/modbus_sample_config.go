//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package modbus

func (m *Modbus) SampleConfig() string {
	return `# Retrieve data from MODBUS slave devices
[[inputs.modbus]]
  ## Connection Configuration
  ##
  ## The plugin supports connections to PLCs via MODBUS/TCP, RTU over TCP, ASCII over TCP or
  ## via serial line communication in binary (RTU) or readable (ASCII) encoding
  ##
  ## Device name
  name = "Device"

  ## Slave ID - addresses a MODBUS device on the bus
  ## Range: 0 - 255 [0 = broadcast; 248 - 255 = reserved]
  slave_id = 1

  ## Timeout for each request
  timeout = "1s"

  ## Maximum number of retries and the time to wait between retries
  ## when a slave-device is busy.
  # busy_retries = 0
  # busy_retries_wait = "100ms"

  # TCP - connect via Modbus/TCP
  controller = "tcp://localhost:502"

  ## Serial (RS485; RS232)
  # controller = "file:///dev/ttyUSB0"
  # baud_rate = 9600
  # data_bits = 8
  # parity = "N"
  # stop_bits = 1

  ## For Modbus over TCP you can choose between "TCP", "RTUoverTCP" and "ASCIIoverTCP"
  ## default behaviour is "TCP" if the controller is TCP
  ## For Serial you can choose between "RTU" and "ASCII"
  # transmission_mode = "RTU"

  ## Trace the connection to the modbus device as debug messages
  ## Note: You have to enable telegraf's debug mode to see those messages!
  # debug_connection = false

  ## Define the configuration schema
  ##  |---register -- define fields per register type in the original style (only supports one slave ID)
  ##  |---request  -- define fields on a requests base
  configuration_type = "register"

  ## --- "register" configuration style ---

  ## Measurements
  ##

  ## Digital Variables, Discrete Inputs and Coils
  ## measurement - the (optional) measurement name, defaults to "modbus"
  ## name        - the variable name
  ## address     - variable address

  discrete_inputs = [
    { name = "start",          address = [0]},
    { name = "stop",           address = [1]},
    { name = "reset",          address = [2]},
    { name = "emergency_stop", address = [3]},
  ]
  coils = [
    { name = "motor1_run",     address = [0]},
    { name = "motor1_jog",     address = [1]},
    { name = "motor1_stop",    address = [2]},
  ]

  ## Analog Variables, Input Registers and Holding Registers
  ## measurement - the (optional) measurement name, defaults to "modbus"
  ## name        - the variable name
  ## byte_order  - the ordering of bytes
  ##  |---AB, ABCD   - Big Endian
  ##  |---BA, DCBA   - Little Endian
  ##  |---BADC       - Mid-Big Endian
  ##  |---CDAB       - Mid-Little Endian
  ## data_type   - INT16, UINT16, INT32, UINT32, INT64, UINT64,
  ##               FLOAT32-IEEE, FLOAT64-IEEE (the IEEE 754 binary representation)
  ##               FLOAT32, FIXED, UFIXED (fixed-point representation on input)
  ## scale       - the final numeric variable representation
  ## address     - variable address

  holding_registers = [
    { name = "power_factor", byte_order = "AB",   data_type = "FIXED", scale=0.01,  address = [8]},
    { name = "voltage",      byte_order = "AB",   data_type = "FIXED", scale=0.1,   address = [0]},
    { name = "energy",       byte_order = "ABCD", data_type = "FIXED", scale=0.001, address = [5,6]},
    { name = "current",      byte_order = "ABCD", data_type = "FIXED", scale=0.001, address = [1,2]},
    { name = "frequency",    byte_order = "AB",   data_type = "UFIXED", scale=0.1,  address = [7]},
    { name = "power",        byte_order = "ABCD", data_type = "UFIXED", scale=0.1,  address = [3,4]},
  ]
  input_registers = [
    { name = "tank_level",   byte_order = "AB",   data_type = "INT16",   scale=1.0,     address = [0]},
    { name = "tank_ph",      byte_order = "AB",   data_type = "INT16",   scale=1.0,     address = [1]},
    { name = "pump1_speed",  byte_order = "ABCD", data_type = "INT32",   scale=1.0,     address = [3,4]},
  ]

  ## --- "request" configuration style ---

  ## Per request definition
  ##

  ## Define a request sent to the device
  ## Multiple of those requests can be defined. Data will be collated into metrics at the end of data collection.
  [[inputs.modbus.request]]
    ## ID of the modbus slave device to query.
    ## If you need to query multiple slave-devices, create several "request" definitions.
    slave_id = 1

    ## Byte order of the data.
    ##  |---ABCD -- Big Endian (Motorola)
    ##  |---DCBA -- Little Endian (Intel)
    ##  |---BADC -- Big Endian with byte swap
    ##  |---CDAB -- Little Endian with byte swap
    byte_order = "ABCD"

    ## Type of the register for the request
    ## Can be "coil", "discrete", "holding" or "input"
    register = "coil"

    ## Name of the measurement.
    ## Can be overriden by the individual field definitions. Defaults to "modbus"
    # measurement = "modbus"

    ## Field definitions
    ## Analog Variables, Input Registers and Holding Registers
    ## address        - address of the register to query. For coil and discrete inputs this is the bit address.
    ## name *1        - field name
    ## type *1,2      - type of the modbus field, can be INT16, UINT16, INT32, UINT32, INT64, UINT64 and
    ##                  FLOAT32, FLOAT64 (IEEE 754 binary representation)
    ## scale *1,2     - (optional) factor to scale the variable with
    ## output *1,2    - (optional) type of resulting field, can be INT64, UINT64 or FLOAT64. Defaults to FLOAT64 if
    ##                  "scale" is provided and to the input "type" class otherwise (i.e. INT* -> INT64, etc).
    ## measurement *1 - (optional) measurement name, defaults to the setting of the request
    ## omit           - (optional) omit this field. Useful to leave out single values when querying many registers
    ##                  with a single request. Defaults to "false".
    ##
    ## *1: Those fields are ignored if field is omitted ("omit"=true)
    ##
    ## *2: Thise fields are ignored for both "coil" and "discrete"-input type of registers. For those register types
    ##     the fields are output as zero or one in UINT64 format by default.

    ## Coil / discrete input example
    fields = [
      { address=0, name="motor1_run"},
      { address=1, name="jog", measurement="motor"},
      { address=2, name="motor1_stop", omit=true},
      { address=3, name="motor1_overheating"},
    ]

    [[inputs.modbus.request.tags]]
      machine = "impresser"
      location = "main building"

  [[inputs.modbus.request]]
    ## Holding example
    ## All of those examples will result in FLOAT64 field outputs
    slave_id = 1
    byte_order = "DCBA"
    register = "holding"
    fields = [
      { address=0, name="voltage",      type="INT16",   scale=0.1   },
      { address=1, name="current",      type="INT32",   scale=0.001 },
      { address=3, name="power",        type="UINT32",  omit=true   },
      { address=5, name="energy",       type="FLOAT32", scale=0.001, measurement="W" },
      { address=7, name="frequency",    type="UINT32",  scale=0.1   },
      { address=8, name="power_factor", type="INT64",   scale=0.01  },
    ]

    [[inputs.modbus.request.tags]]
      machine = "impresser"
      location = "main building"

  [[inputs.modbus.request]]
    ## Input example with type conversions
    slave_id = 1
    byte_order = "ABCD"
    register = "input"
    fields = [
      { address=0, name="rpm",         type="INT16"                   },  # will result in INT64 field
      { address=1, name="temperature", type="INT16", scale=0.1        },  # will result in FLOAT64 field
      { address=2, name="force",       type="INT32", output="FLOAT64" },  # will result in FLOAT64 field
      { address=4, name="hours",       type="UINT32"                  },  # will result in UIN64 field
    ]

    [[inputs.modbus.request.tags]]
      machine = "impresser"
      location = "main building"

  ## Enable workarounds required by some devices to work correctly
  # [inputs.modbus.workarounds]
    ## Pause between read requests sent to the device. This might be necessary for (slow) serial devices.
    # pause_between_requests = "0ms"
    ## Close the connection after every gather cycle. Usually the plugin closes the connection after a certain
    ## idle-timeout, however, if you query a device with limited simultaneous connectivity (e.g. serial devices)
    ## from multiple instances you might want to only stay connected during gather and disconnect afterwards.
    # close_connection_after_gather = false
`
}
