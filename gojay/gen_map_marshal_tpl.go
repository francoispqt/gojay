package main

var mapMarshalTpl = templateList{
	"def": &genTpl{
		strTpl: "\n// MarshalJSONObject implements gojay's MarshalerJSONObject" +
			"\nfunc (v {{.StructName}}) MarshalJSONObject(enc *gojay.Encoder) {\n",
	},
	"isNil": &genTpl{
		strTpl: `
// IsNil returns wether the structure is nil value or not
func (v {{.StructName}}) IsNil() bool { return v == nil || len(v) == 0 }
`,
	},
	"string": &genTpl{
		strTpl: `	for k, s := range v {
		enc.StringKey(k, s)
	}
`,
	},
	"int": &genTpl{
		strTpl: `	for k, s := range v {
		enc.Int{{.IntLen}}Key(k, s)
	}
`,
	},
	"uint": &genTpl{
		strTpl: `	for k, s := range v {
		enc.Uint{{.IntLen}}Key(k, s)
	}
`,
	},
	"float": &genTpl{
		strTpl: `	for k, s := range v {
		enc.Float{{.IntLen}}Key(k, s)
	}
`,
	},
	"bool": &genTpl{
		strTpl: `	for k, s := range v {
		enc.BoolKey(k, s)
	}
`,
	},
	"struct": &genTpl{
		strTpl: `	for k, s := range v {
		enc.ObjectKey(k, s)
	}
`,
	},
	"arr": &genTpl{
		strTpl: `	for k, s := range v {
		enc.ArrayKey(k, s)
	}
`,
	},
}

func init() {
	parseTemplates(mapMarshalTpl, "mapMarshal")
}
