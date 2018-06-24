package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenArray(t *testing.T) {
	testCases := map[string]struct {
		input          io.Reader
		expectedResult string
	}{
		"basicStringSlice": {
			input: strings.NewReader(`package test

//gojay:json
type StrSlice []string
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *StrSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var str string
	if err := dec.String(&str); err != nil {
		return err
	}
	*v = append(*v, str)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *StrSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.String(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *StrSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicStringBool": {
			input: strings.NewReader(`package test

//gojay:json
type BoolSlice []bool
			`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *BoolSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var b bool
	if err := dec.Bool(&b); err != nil {
		return err
	}
	*v = append(*v, b)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *BoolSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Bool(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *BoolSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicIntSlice": {
			input: strings.NewReader(`package test

//gojay:json
type IntSlice []int
							`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i int
	if err := dec.Int(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Int(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicInt8Slice": {
			input: strings.NewReader(`package test
	
	//gojay:json
	type IntSlice []int8
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i int8
	if err := dec.Int8(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Int8(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicInt16Slice": {
			input: strings.NewReader(`package test

//gojay:json
type IntSlice []int16
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i int16
	if err := dec.Int16(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Int16(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicInt32Slice": {
			input: strings.NewReader(`package test
	
	//gojay:json
	type IntSlice []int32
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i int32
	if err := dec.Int32(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Int32(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicInt64Slice": {
			input: strings.NewReader(`package test
	
	//gojay:json
	type IntSlice []int64
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i int64
	if err := dec.Int64(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Int64(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicUint64Slice": {
			input: strings.NewReader(`package test
	
	//gojay:json
	type IntSlice []uint64
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i uint64
	if err := dec.Uint64(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Uint64(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicFloatSlice": {
			input: strings.NewReader(`package test
	
//gojay:json
type IntSlice []float64
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i float64
	if err := dec.Float64(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Float64(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicFloat32Slice": {
			input: strings.NewReader(`package test
	
//gojay:json
type IntSlice []float32
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i float32
	if err := dec.Float32(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Float32(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicStructSlice": {
			input: strings.NewReader(`package test
	
//gojay:json
type StructSlice []*Struct

type Struct struct{
	Str string
}
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *StructSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var s = &Struct{}
	if err := dec.Object(s); err != nil {
		return err
	}
	*v = append(*v, s)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *StructSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Object(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *StructSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
`,
		},
		"basicSliceSlice": {
			input: strings.NewReader(`package test
	
//gojay:json
type SliceStrSlice []StrSlice

type StrSlice []string
								`),
			expectedResult: `package  

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *SliceStrSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var s = make(StrSlice, 0)
	if err := dec.Array(&s); err != nil {
		return err
	}
	*v = append(*v, s)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *SliceStrSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Array(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *SliceStrSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
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
			assert.Equal(
				t,
				string(genHeader)+testCase.expectedResult,
				g.b.String(),
			)
		})
	}
}
