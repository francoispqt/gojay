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

func BenchmarkEncodingJsonEncodeLargeStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(benchmarks.NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkJsonIterEncodeLargeStruct(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(benchmarks.NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyJsonEncodeObjLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := easyjson.Marshal(benchmarks.NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoJayEncodeLargeStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(benchmarks.NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func TestGoJayEncodeLargeStruct(t *testing.T) {
	if output, err := gojay.MarshalJSONObject(benchmarks.NewLargePayload()); err != nil {
		t.Fatal(err)
	} else {
		log.Print(string(output))
	}
}
