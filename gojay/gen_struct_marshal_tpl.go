package main

var structMarshalTpl = templateList{
	"def": &genTpl{
		strTpl: "\n// MarshalJSONObject implements gojay's MarshalerJSONObject" +
			"\nfunc (v *{{.StructName}}) MarshalJSONObject(enc *gojay.Encoder) {\n",
	},
	"isNil": &genTpl{
		strTpl: `
// IsNil returns wether the structure is nil value or not
func (v *{{.StructName}}) IsNil() bool { return v == nil }
`,
	},
	"string": &genTpl{
		strTpl: "\tenc.StringKey{{.OmitEmpty}}(\"{{.Key}}\", {{.Ptr}}v.{{.Field}})\n",
	},
	"int": &genTpl{
		strTpl: "\tenc.Int{{.IntLen}}Key{{.OmitEmpty}}(\"{{.Key}}\", {{.Ptr}}v.{{.Field}})\n",
	},
	"uint": &genTpl{
		strTpl: "\tenc.Uint{{.IntLen}}Key{{.OmitEmpty}}(\"{{.Key}}\", {{.Ptr}}v.{{.Field}})\n",
	},
	"float": &genTpl{
		strTpl: "\tenc.Float{{.IntLen}}Key{{.OmitEmpty}}(\"{{.Key}}\", {{.Ptr}}v.{{.Field}})\n",
	},
	"bool": &genTpl{
		strTpl: "\tenc.BoolKey{{.OmitEmpty}}(\"{{.Key}}\", {{.Ptr}}v.{{.Field}})\n",
	},
	"struct": &genTpl{
		strTpl: "\tenc.ObjectKey{{.OmitEmpty}}(\"{{.Key}}\", v.{{.Field}})\n",
	},
	"arr": &genTpl{
		strTpl: "\tenc.ArrayKey{{.OmitEmpty}}(\"{{.Key}}\", v.{{.Field}})\n",
	},
	"any": &genTpl{
		strTpl: "\tenc.AnyKey(\"{{.Key}}\", v.{{.Field}})\n",
	},
}

func init() {
	parseTemplates(structMarshalTpl, "structMarshal")
}
