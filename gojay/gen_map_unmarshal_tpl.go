package main

var mapUnmarshalTpl = templateList{
	"def": &genTpl{
		strTpl: "\n// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject" +
			"\nfunc (v {{.TypeName}}) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {\n",
	},
	"nKeys": &genTpl{
		strTpl: `
// NKeys returns the number of keys to unmarshal
func (v {{.StructName}}) NKeys() int { return {{.NKeys}} }
`,
	},
	"string": &genTpl{
		strTpl: `	var str string
	if err := dec.String(&str); err != nil {
		return err
	}
	v[k] = {{.Ptr}}str
`,
	},
	"int": &genTpl{
		strTpl: `	var i int{{.IntLen}}
	if err := dec.Int{{.IntLen}}(&i); err != nil {
		return err
	}
	v[k] = {{.Ptr}}i
`,
	},
	"uint": &genTpl{
		strTpl: `	var i uint{{.IntLen}}
	if err := dec.Uint{{.IntLen}}(&i); err != nil {
		return err
	}
	v[k] = {{.Ptr}}i
`,
	},
	"float": &genTpl{
		strTpl: `	var i float{{.IntLen}}
	if err := dec.Float{{.IntLen}}(&i); err != nil {
		return err
	}
	v[k] = {{.Ptr}}i
`,
	},
	"bool": &genTpl{
		strTpl: `	var b bool
	if err := dec.Bool(&b); err != nil {
		return err
	}
	v[k] = {{.Ptr}}b
`,
	},
	"struct": &genTpl{
		strTpl: `	var s = {{.StructName}}{}
	if err := dec.Object(&s); err != nil {
		return err
	}
	v[k] = s
`,
	},
	"structPtr": &genTpl{
		strTpl: `	var s = &{{.StructName}}{}
	if err := dec.Object(s); err != nil {
		return err
	}
	v[k] = s
`,
	},
	"arr": &genTpl{
		strTpl: `	var s = &{{.StructName}}{}
		if err := dec.Array(s); err != nil {
			return err
		}
		v[k] = s
`,
	},
}

func init() {
	parseTemplates(mapUnmarshalTpl, "mapUnmarshal")
}
