package benchmarks

import (
	"testing"

	"github.com/buger/jsonparser"
	"github.com/francoispqt/gojay"
	"github.com/francoispqt/gojay/benchmarks"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

func BenchmarkJsonParserDecodeObjLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.ArrayEach(benchmarks.LargeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "username")
			nothing()
		}, "users")

		jsonparser.ArrayEach(benchmarks.LargeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.GetInt(value, "id")
			jsonparser.Get(value, "slug")
			nothing()
		}, "topics", "topics")
	}
}

func BenchmarkJsonIterDecodeObjLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := benchmarks.LargePayload{}
		jsoniter.Unmarshal(benchmarks.LargeFixture, &result)
	}
}

func BenchmarkEasyJsonDecodeObjLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := benchmarks.LargePayload{}
		easyjson.Unmarshal(benchmarks.LargeFixture, &result)
	}
}

func BenchmarkGoJayDecodeObjLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := benchmarks.LargePayload{}
		gojay.UnmarshalJSONObject(benchmarks.LargeFixture, &result)
	}
}

func BenchmarkGoJayUnsafeDecodeObjLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := benchmarks.LargePayload{}
		gojay.Unsafe.UnmarshalJSONObject(benchmarks.LargeFixture, &result)
	}
}
