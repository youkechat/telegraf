//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package converter

func (p *Converter) SampleConfig() string {
	return `# Convert values to another metric value type
[[processors.converter]]
  ## Tags to convert
  ##
  ## The table key determines the target type, and the array of key-values
  ## select the keys to convert.  The array may contain globs.
  ##   <target-type> = [<tag-key>...]
  [processors.converter.tags]
    measurement = []
    string = []
    integer = []
    unsigned = []
    boolean = []
    float = []

  ## Fields to convert
  ##
  ## The table key determines the target type, and the array of key-values
  ## select the keys to convert.  The array may contain globs.
  ##   <target-type> = [<field-key>...]
  [processors.converter.fields]
    measurement = []
    tag = []
    string = []
    integer = []
    unsigned = []
    boolean = []
    float = []
`
}
