package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/francoispqt/gojay"
)

var bigf = []byte(`0.00058273999999999999`)

// BenchmarkBigFloatEncodingJSON decodes a big float with the standard package
func BenchmarkBigFloatEncodingJSON(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var f float64
		var _ = json.Unmarshal(bigf, &f)
	}
}

// BenchmarkBigFloatGojay decodes a big float with gojay
func BenchmarkBigFloatGojay(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		var f float64
		var _ = gojay.Unmarshal(bigf, &f)
	}
}
