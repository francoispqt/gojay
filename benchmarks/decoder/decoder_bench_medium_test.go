package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/francoispqt/gojay"
	"github.com/francoispqt/gojay/benchmarks"
	jsoniter "github.com/json-iterator/go"
)

func BenchmarkJsonIterDecodeObjMedium(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		result := benchmarks.MediumPayload{}
		jsoniter.Unmarshal(benchmarks.MediumFixture, &result)
	}
}

/*
   github.com/buger/jsonparser
*/
func BenchmarkJSONParserDecodeObjMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsonparser.Get(benchmarks.MediumFixture, "person", "name", "fullName")
		jsonparser.GetInt(benchmarks.MediumFixture, "person", "github", "followers")
		jsonparser.Get(benchmarks.MediumFixture, "company")

		jsonparser.ArrayEach(benchmarks.MediumFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "url")
			nothing()
		}, "person", "gravatar", "avatars")
	}
}

func BenchmarkEncodingJsonStructMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var data = benchmarks.MediumPayload{}
		json.Unmarshal(benchmarks.MediumFixture, &data)

		nothing(data.Person.Name.FullName, data.Person.Github.Followers, data.Company)

		for _, el := range data.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkGoJayDecodeObjMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := benchmarks.MediumPayload{}
		err := gojay.UnmarshalObject(benchmarks.MediumFixture, &result)
		if err != nil {
			b.Error(err)
		}
	}
}
