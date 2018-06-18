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
			assert.Equal(t, testCase.expectedResult, g.b.String())
		})
	}
}
