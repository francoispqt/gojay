package main

var arrUnmarshalTpl = templateList{
	"def": &genTpl{
		strTpl: "\n// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray" +
			"\nfunc (v *{{.TypeName}}) UnmarshalJSONArray(dec *gojay.Decoder) error {\n",
	},
	"string": &genTpl{
		strTpl: "\tvar str string" +
			"\n\tif err := dec.String(&str); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, str)\n",
	},
	"stringPtr": &genTpl{
		strTpl: "\n// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray" +
			"\nfunc (v *{{.TypeName}}) UnmarshalJSONArray(dec *gojay.Decoder) error {\n",
	},
	"int": &genTpl{
		strTpl: "\tvar i int{{.IntLen}}" +
			"\n\tif err := dec.Int{{.IntLen}}(&i); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, i)\n",
	},
	"uint": &genTpl{
		strTpl: "\tvar i uint{{.IntLen}}" +
			"\n\tif err := dec.Uint{{.IntLen}}(&i); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, i)\n",
	},
	"float": &genTpl{
		strTpl: "\tvar i float{{.IntLen}}" +
			"\n\tif err := dec.Float{{.IntLen}}(&i); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, i)\n",
	},
	"bool": &genTpl{
		strTpl: "\tvar b bool" +
			"\n\tif err := dec.Bool(&b); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, b)\n",
	},
	"struct": &genTpl{
		strTpl: "\tvar s = {{.StructName}}{}" +
			"\n\tif err := dec.Object(&s); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, s)\n",
	},
	"structPtr": &genTpl{
		strTpl: "\tvar s = &{{.StructName}}{}" +
			"\n\tif err := dec.Object(s); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, s)\n",
	},
	"arr": &genTpl{
		strTpl: "\tvar s = make({{.StructName}}, 0)" +
			"\n\tif err := dec.Array(&s); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, s)\n",
	},
	"arrPtr": &genTpl{
		strTpl: "\tvar s = make({{.StructName}}, 0)" +
			"\n\tif err := dec.Array(&s); err != nil {\n" +
			"\t\treturn err\n\t}\n" +
			"\t*v = append(*v, &s)\n",
	},
}

func init() {
	parseTemplates(arrUnmarshalTpl, "arrUnmarshal")
}
