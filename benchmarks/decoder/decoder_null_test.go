package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/francoispqt/gojay"
)

type ObjNullReflect struct {
	O *ObjNullReflect
}

func (o *ObjNullReflect) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "o":
		return dec.ObjectNullReflect(&o.O)
	}
	return nil
}

func (o *ObjNullReflect) NKeys() int {
	return 0
}

type ObjNullFactory struct {
	O *ObjNullFactory
}

func (o *ObjNullFactory) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "o":
		return dec.ObjectNullFactory(func() gojay.UnmarshalerJSONObject {
			o.O = &ObjNullFactory{}
			return o.O
		})
	}
	return nil
}

func (o *ObjNullFactory) NKeys() int {
	return 0
}

var objNullJSON = []byte(`{"o":{}}`)

func BenchmarkJSONDecodeObjNullReflection(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		result := &ObjNullReflect{}
		json.Unmarshal(objNullJSON, &result)
	}
}

func BenchmarkJSONDecodeObjNullFactory(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		result := &ObjNullFactory{}
		json.Unmarshal(objNullJSON, &result)
	}
}
