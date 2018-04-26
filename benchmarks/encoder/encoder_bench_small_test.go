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
		if _, err := gojay.MarshalObject(benchmarks.NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func TestGoJayEncodeSmallStruct(t *testing.T) {
	if output, err := gojay.MarshalObject(benchmarks.NewSmallPayload()); err != nil {
		t.Fatal(err)
	} else {
		log.Print(output)
	}
}
