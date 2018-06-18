package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenStruct(t *testing.T) {
	testCases := map[string]struct {
		input          io.Reader
		expectedResult string
	}{
		"basicStruct": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Int int
	Int8 int8
	Int16 int16
	Int32 int32
	Int64 int64
	Uint64 uint64
	Float float64
	Str string
	Bool bool
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "int":
		return dec.Int(&v.Int)
	case "int8":
		return dec.Int8(&v.Int8)
	case "int16":
		return dec.Int16(&v.Int16)
	case "int32":
		return dec.Int32(&v.Int32)
	case "int64":
		return dec.Int64(&v.Int64)
	case "uint64":
		return dec.Uint64(&v.Uint64)
	case "float":
		return dec.Float(&v.Float)
	case "str":
		return dec.String(&v.Str)
	case "bool":
		return dec.Bool(&v.Bool)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 9 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("int", v.Int)
	enc.Int8Key("int8", v.Int8)
	enc.Int16Key("int16", v.Int16)
	enc.Int32Key("int32", v.Int32)
	enc.Int64Key("int64", v.Int64)
	enc.Uint64Key("uint64", v.Uint64)
	enc.FloatKey("float", v.Float)
	enc.StringKey("str", v.Str)
	enc.BoolKey("bool", v.Bool)
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
		"basicStructPtr": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Int *int
	Str *string
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "int":
		return dec.Int(v.Int)
	case "str":
		return dec.String(v.Str)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("int", *v.Int)
	enc.StringKey("str", *v.Str)
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
		"basicStructTags": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Int int ` + "`gojay:\"someInt\"`" + `
	Str string ` + "`gojay:\"someStr\"`" + `
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "someInt":
		return dec.Int(&v.Int)
	case "someStr":
		return dec.String(&v.Str)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("someInt", v.Int)
	enc.StringKey("someStr", v.Str)
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
		"basicStructTagsHideUnmarshal": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Int int ` + "`gojay:\"-u\"`" + `
	Str string ` + "`gojay:\"-u\"`" + `
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("int", v.Int)
	enc.StringKey("str", v.Str)
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
		"basicStructTagsHideUnmarshal2": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Int int ` + "`gojay:\"someInt,-u\"`" + `
	Str string ` + "`gojay:\"someStr,-u\"`" + `
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("someInt", v.Int)
	enc.StringKey("someStr", v.Str)
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
		"basicStructTagsHideUnmarshal3": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Int int ` + "`gojay:\"someInt,-m\"`" + `
	Str string ` + "`gojay:\"someStr,-m\"`" + `
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "someInt":
		return dec.Int(&v.Int)
	case "someStr":
		return dec.String(&v.Str)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 2 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
		"basicStructTagsHideUnmarshal4": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Int int ` + "`gojay:\"-\"`" + `
	Str string ` + "`gojay:\"-\"`" + `
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
		"complexStructStructTag": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Struct Struct ` + "`gojay:\"someStruct\"`" + `
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "someStruct":
		if v.Struct == nil {
			v.Struct = Struct{}
		}
		return dec.Object(v.Struct)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("someStruct", v.Struct)
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
		"complexStructStructPtrTag": {
			input: strings.NewReader(`package test

//gojay:json
type Struct struct{
	Struct *Struct ` + "`gojay:\"someStruct\"`" + `
}
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Struct) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "someStruct":
		if v.Struct == nil {
			v.Struct = &Struct{}
		}
		return dec.Object(v.Struct)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Struct) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Struct) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("someStruct", v.Struct)
}

// IsNil returns wether the structure is nil value or not
func (v *Struct) IsNil() bool { return v == nil }
`,
		},
	}
	for n, testCase := range testCases {
		t.Run(n, func(t *testing.T) {
			g, err := MakeGenFromReader(testCase.input)
			if err != nil {
				t.Fatal(err)
			}
			err = g.Gen()
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, testCase.expectedResult, g.b.String())
		})
	}
}
