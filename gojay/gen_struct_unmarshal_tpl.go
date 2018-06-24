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
	"float": &genTpl{
		strTpl: "\t\treturn dec.Float{{.IntLen}}({{.Ptr}}v.{{.Field}})\n",
	},
	"bool": &genTpl{
		strTpl: "\t\treturn dec.Bool({{.Ptr}}v.{{.Field}})\n",
	},
	"struct": &genTpl{
		strTpl: `		if v.{{.Field}} == nil {
			v.{{.Field}} = {{.StructName}}{}
		}
		return dec.Object(v.{{.Field}})
`,
	},
	"structPtr": &genTpl{
		strTpl: `		if v.{{.Field}} == nil {
			v.{{.Field}} = &{{.StructName}}{}
		}
		return dec.Object(v.{{.Field}})
`,
	},
	"arr": &genTpl{
		strTpl: `		if v.{{.Field}} == nil {
			arr := make({{.TypeName}}, 0)
			v.{{.Field}} = arr
		}
		return dec.Array(&v.{{.Field}})
`,
	},
	"arrPtr": &genTpl{
		strTpl: `		if v.{{.Field}} == nil {
			arr := make({{.TypeName}}, 0)
			v.{{.Field}} = &arr
		}
		return dec.Array(v.{{.Field}})
`,
	},
	"any": &genTpl{
		strTpl: "\t\treturn dec.Any(&v.{{.Field}})\n",
	},
	"anyPtr": &genTpl{
		strTpl: "\t\treturn dec.Any(v.{{.Field}})\n",
	},
}

func init() {
	parseTemplates(structUnmarshalTpl, "structUnmarshal")
}
