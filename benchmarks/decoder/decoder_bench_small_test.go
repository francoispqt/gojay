package benchmarks

import (
	"encoding/json"
	_ "fmt"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/francoispqt/gojay"
	"github.com/francoispqt/gojay/benchmarks"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

func BenchmarkJSONDecodeObjSmall(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		result := benchmarks.SmallPayload{}
		json.Unmarshal(benchmarks.SmallFixture, &result)
	}
}

func nothing(_ ...interface{}) {}
func BenchmarkJSONParserSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.GetInt(benchmarks.SmallFixture, "tz")
		jsonparser.GetInt(benchmarks.SmallFixture, "v")
		jsonparser.GetInt(benchmarks.SmallFixture, "sid")
		jsonparser.GetInt(benchmarks.SmallFixture, "st")
		jsonparser.GetInt(benchmarks.SmallFixture, "gr")
		jsonparser.Get(benchmarks.SmallFixture, "uuid")
		jsonparser.Get(benchmarks.SmallFixture, "ua")

		nothing()
	}
}

func BenchmarkJsonIterDecodeObjSmall(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		result := benchmarks.SmallPayload{}
		jsoniter.Unmarshal(benchmarks.SmallFixture, &result)
	}
}

func BenchmarkEasyJsonDecodeObjSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := benchmarks.SmallPayload{}
		easyjson.Unmarshal(benchmarks.SmallFixture, &result)
	}
}

func BenchmarkGoJayDecodeObjSmall(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		result := benchmarks.SmallPayload{}
		gojay.UnmarshalJSONObject(benchmarks.SmallFixture, &result)
	}
}

func BenchmarkGoJayUnsafeDecodeObjSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := benchmarks.SmallPayload{}
		gojay.Unsafe.UnmarshalJSONObject(benchmarks.SmallFixture, &result)
	}
}
