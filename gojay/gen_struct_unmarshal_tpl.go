package main

var structUnmarshalTpl = templateList{
	"def": &genTpl{
		strTpl: "\n// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject" +
			"\nfunc (v *{{.StructName}}) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {\n",
	},
	"nKeys": &genTpl{
		strTpl: `
// NKeys returns the number of keys to unmarshal
func (v *{{.StructName}}) NKeys() int { return {{.NKeys}} }
`,
	},
	"case": &genTpl{
		strTpl: "\tcase \"{{.Key}}\":\n",
	},
	"string": &genTpl{
		strTpl: "\t\treturn dec.String({{.Ptr}}v.{{.Field}})\n",
	},
	"int": &genTpl{
		strTpl: "\t\treturn dec.Int{{.IntLen}}({{.Ptr}}v.{{.Field}})\n",
	},
	"uint": &genTpl{
		strTpl: "\t\treturn dec.Uint{{.IntLen}}({{.Ptr}}v.{{.Field}})\n",
	},
	"bool": &genTpl{
		strTpl: "\t\treturn dec.Bool({{.Ptr}}v.{{.Field}})\n",
	},
	"struct": &genTpl{
		strTpl: `		if v.{{.Field}} == nil {
			v.{{.Field}} = &{{.StructName}}{}
		}
		dec.Object(v.{{.Field}})
`,
	},
	"arr": &genTpl{
		strTpl: `		if v.{{.Field}} == nil {
			arr := make({{.TypeName}}, 0)
			v.{{.Field}} = &arr
		}
		dec.Array(v.{{.Field}})
`,
	},
}

func init() {
	parseTemplates(structUnmarshalTpl, "structUnmarshal")
}
