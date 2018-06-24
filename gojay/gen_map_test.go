package main

import (
	"io"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenMap(t *testing.T) {
	testCases := map[string]struct {
		input          io.Reader
		expectedResult string
	}{
		"basicMapStringString": {
			input: strings.NewReader(`package test

//gojay:json
type StrMap map[string]string
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v StrMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var str string
	if err := dec.String(&str); err != nil {
		return err
	}
	v[k] = str
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v StrMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v StrMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.StringKey(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v StrMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringStringPtr": {
			input: strings.NewReader(`package test

//gojay:json
type StrMap map[string]*string
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v StrMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var str string
	if err := dec.String(&str); err != nil {
		return err
	}
	v[k] = &str
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v StrMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v StrMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.StringKey(k, *s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v StrMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringInt": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]int
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i int
	if err := dec.Int(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.IntKey(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringIntPtr": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]*int
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i int
	if err := dec.Int(&i); err != nil {
		return err
	}
	v[k] = &i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.IntKey(k, *s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringInt64": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]int64
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i int64
	if err := dec.Int64(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Int64Key(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringInt64Ptr": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]*int64
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i int64
	if err := dec.Int64(&i); err != nil {
		return err
	}
	v[k] = &i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Int64Key(k, *s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringInt32": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]int32
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i int32
	if err := dec.Int32(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Int32Key(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringInt16": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]int16
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i int16
	if err := dec.Int16(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Int16Key(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringInt8": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]int8
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i int8
	if err := dec.Int8(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Int8Key(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringUint64": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]uint64
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i uint64
	if err := dec.Uint64(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Uint64Key(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringUint64Ptr": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]*uint64
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i uint64
	if err := dec.Uint64(&i); err != nil {
		return err
	}
	v[k] = &i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Uint64Key(k, *s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringUint32": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]uint32
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i uint32
	if err := dec.Uint32(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Uint32Key(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringUint16": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]uint16
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i uint16
	if err := dec.Uint16(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Uint16Key(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringUint8": {
			input: strings.NewReader(`package test

//gojay:json
type IntMap map[string]uint8
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v IntMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var i uint8
	if err := dec.Uint8(&i); err != nil {
		return err
	}
	v[k] = i
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v IntMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v IntMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.Uint8Key(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v IntMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringBool": {
			input: strings.NewReader(`package test

//gojay:json
type BoolMap map[string]bool
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v BoolMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var b bool
	if err := dec.Bool(&b); err != nil {
		return err
	}
	v[k] = b
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v BoolMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v BoolMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.BoolKey(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v BoolMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringBoolPtr": {
			input: strings.NewReader(`package test

//gojay:json
type BoolMap map[string]*bool
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v BoolMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var b bool
	if err := dec.Bool(&b); err != nil {
		return err
	}
	v[k] = &b
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v BoolMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v BoolMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.BoolKey(k, *s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v BoolMap) IsNil() bool { return v == nil || len(v) == 0 }
`,
		},
		"basicMapStringStruct": {
			input: strings.NewReader(`package test

//gojay:json
type BoolMap map[string]*Test

type Test struct{}
`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v BoolMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	var s = &Test{}
	if err := dec.Object(s); err != nil {
		return err
	}
	v[k] = s
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v BoolMap) NKeys() int { return 0 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v BoolMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, s := range v {
		enc.ObjectKey(k, s)
	}
}

// IsNil returns wether the structure is nil value or not
func (v BoolMap) IsNil() bool { return v == nil || len(v) == 0 }
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
			log.Print(g.b.String())
			assert.Equal(
				t,
				string(genHeader)+testCase.expectedResult,
				g.b.String(),
			)
		})
	}
}
