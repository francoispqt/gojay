package benchmarks

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/francoispqt/gojay/benchmarks"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

func BenchmarkEncodingJsonEncodeSmallStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(benchmarks.NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyJsonEncodeObjSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := easyjson.Marshal(benchmarks.NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonIterEncodeSmallStruct(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(benchmarks.NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoJayEncodeSmallStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(benchmarks.NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoJayEncodeSmallFunc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(gojay.EncodeObjectFunc(func(enc *gojay.Encoder) {
			enc.AddIntKey("st", 1)
			enc.AddIntKey("sid", 1)
			enc.AddStringKey("tt", "test")
			enc.AddIntKey("gr", 1)
			enc.AddStringKey("uuid", "test")
			enc.AddStringKey("ip", "test")
			enc.AddStringKey("ua", "test")
			enc.AddIntKey("tz", 1)
			enc.AddIntKey("v", 1)
		})); err != nil {
			b.Fatal(err)
		}
	}
}

func TestGoJayEncodeSmallStruct(t *testing.T) {
	if output, err := gojay.MarshalJSONObject(benchmarks.NewSmallPayload()); err != nil {
		t.Fatal(err)
	} else {
		log.Print(output)
	}
}
