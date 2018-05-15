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

func BenchmarkEncodingJsonEncodeMediumStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(benchmarks.NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonIterEncodeMediumStruct(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(benchmarks.NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyJsonEncodeObjMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := easyjson.Marshal(benchmarks.NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoJayEncodeMediumStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(benchmarks.NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func TestGoJayEncodeMediumStruct(t *testing.T) {
	if output, err := gojay.MarshalJSONObject(benchmarks.NewMediumPayload()); err != nil {
		t.Fatal(err)
	} else {
		log.Print(string(output))
	}
}
