package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func (s *slicesTestObject) MarshalJSONObject(enc *Encoder) {
	enc.AddSliceStringKey("sliceString", s.sliceString)
	enc.AddSliceIntKey("sliceInt", s.sliceInt)
	enc.AddSliceFloat64Key("sliceFloat64", s.sliceFloat64)
	enc.AddSliceBoolKey("sliceBool", s.sliceBool)
}

func (s *slicesTestObject) IsNil() bool {
	return s == nil
}

func TestEncodeSlices(t *testing.T) {
	testCases := []struct {
		name string
		json string
		obj  slicesTestObject
	}{
		{
			name: "basic slice string",
			json: `{
				"sliceString": ["foo","bar"],
				"sliceInt": [],
				"sliceFloat64": [],
				"sliceBool": []
			}`,
			obj: slicesTestObject{
				sliceString: []string{"foo", "bar"},
			},
		},
		{
			name: "basic slice bool",
			json: `{
				"sliceString": [],
				"sliceInt": [],
				"sliceFloat64": [],
				"sliceBool": [true,false]
			}`,
			obj: slicesTestObject{
				sliceBool: []bool{true, false},
			},
		},
		{
			name: "basic slice int",
			json: `{
				"sliceString": [],
				"sliceFloat64": [],
				"sliceInt": [1,2,3],
				"sliceBool": []
			}`,
			obj: slicesTestObject{
				sliceInt: []int{1, 2, 3},
			},
		},
		{
			name: "basic slice float64",
			json: `{
				"sliceString": [],
				"sliceFloat64": [1.3,2.4,3.1],
				"sliceInt": [],
				"sliceBool": []
			}`,
			obj: slicesTestObject{
				sliceFloat64: []float64{1.3, 2.4, 3.1},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				b := strings.Builder{}
				enc := BorrowEncoder(&b)
				err := enc.Encode(&testCase.obj)
				require.Nil(t, err, "err should be nil")
				require.JSONEq(t, testCase.json, b.String())
			},
		)
	}
}
