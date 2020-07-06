package gojay

import (
	"encoding/json"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestDecoder_NextToken(t *testing.T) {
	type args struct {
		json string
	}
	testCases := []struct {
		name string
		args args
		want Token
	}{
		{
			name: "should determine that a slice is next in the buffer",
			args: args{
				json: `[`,
			},
			want: TokenArray,
		},
		{
			name: "should determine that a string is next in the buffer the implementation should convert to a slice",
			args: args{
				json: `"`,
			},
			want: TokenString,
		},
		{
			name: "should determine that a numeric is next in the buffer",
			args: args{
				json: `{`,
			},
			want: TokenObject,
		},
		{
			name: "should determine that a numeric is next in the buffer",
			args: args{
				json: `0`,
			},
			want: TokenNumber,
		},
		{
			name: "should determine that a numeric is next in the buffer",
			args: args{
				json: `1`,
			},
			want: TokenNumber,
		},
		{
			name: "should determine that a numeric is next in the buffer",
			args: args{
				json: `9`,
			},
			want: TokenNumber,
		},
		{
			name: "should determine that a numeric is next in the buffer",
			args: args{
				json: `39`,
			},
			want: TokenNumber,
		},
		{
			name: "should determine that a boolean is next in the buffer",
			args: args{
				json: `true`,
			},
			want: TokenBoolean,
		},
		{
			name: "should determine that a boolean is next in the buffer",
			args: args{
				json: `false`,
			},
			want: TokenBoolean,
		},
		{
			name: "should determine that a null is next in the buffer",
			args: args{
				json: `null`,
			},
			want: TokenNull,
		},
		{
			name: "should not be able to determine the token",
			args: args{
				json: `z`,
			},
			want: TokenUnknown,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			dec := BorrowDecoder(strings.NewReader(tt.args.json))
			if got := dec.NextToken(); got != tt.want {
				t.Fatalf("NextToken() got %v wanted %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_NextToken_UnmarshalJSONObject(t *testing.T) {
	type args struct {
		json string
	}
	testCases := []struct {
		name string
		args args
		want sliceTestObjectWithToken
	}{
		{
			name: "should determine that a slice is next in the buffer",
			args: args{
				json: `{"sliceString": ["zip", "zap"]}`,
			},
			want: sliceTestObjectWithToken{
				SliceString: []string{"zip", "zap"},
			},
		},
		{
			name: "should determine that a string is next in the buffer the implementation should convert to a slice",
			args: args{
				json: `{"sliceString": "zip,zap"}`,
			},
			want: sliceTestObjectWithToken{
				SliceString: []string{"zip", "zap"},
			},
		},
		{
			name: "should determine that a numeric is next in the buffer the implementation should convert to a slice",
			args: args{
				json: `{"sliceString": 10}`,
			},
			want: sliceTestObjectWithToken{
				SliceString: []string{"10"},
			},
		},
		{
			name: "should determine that an object is next in the buffer the implementation should convert to a slice",
			args: args{
				json: `{"sliceString": {"zip": "zap"}}`,
			},
			want: sliceTestObjectWithToken{
				SliceString: []string{"zip", "zap"},
			},
		},
		{
			name: "should determine that a boolean is next in the buffer the implementation should convert to a slice",
			args: args{
				json: `{"sliceString": true}`,
			},
			want: sliceTestObjectWithToken{
				SliceString: []string{"true"},
			},
		},
		{
			name: "should determine that null is next in the buffer the implementation should convert to a slice",
			args: args{
				json: `{"sliceString": null}`,
			},
			want: sliceTestObjectWithToken{
				SliceString: []string{},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			dec := BorrowDecoder(strings.NewReader(tt.args.json))

			var got sliceTestObjectWithToken
			err := dec.Decode(&got)
			if err != nil {
				t.Fatalf("Decode() err %v", err)
			}

			sort.Strings(got.SliceString)
			sort.Strings(tt.want.SliceString)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("NextToken_UnmarshalJSONObject() got %+v wanted %+v", got, tt.want)
			}
		})
	}
}

type sliceTestObjectWithToken struct {
	SliceString []string
}

type testSlice []string

// UnmarshalJSONArray decodes JSON array elements into slice
func (ts *testSlice) UnmarshalJSONArray(dec *Decoder) error {
	var value string
	if err := dec.String(&value); err != nil {
		return err
	}
	*ts = append(*ts, value)
	return nil
}

func (s *sliceTestObjectWithToken) UnmarshalJSONObject(dec *Decoder, k string) error {
	switch k {
	case "sliceString":
		switch dec.NextToken() {
		case TokenArray:
			var aSlice testSlice
			err := dec.Array(&aSlice)
			if err != nil {
				return err
			}

			s.SliceString = aSlice
		case TokenString:
			var ss string
			err := dec.String(&ss)
			if err != nil {
				return err
			}
			arrSlice := strings.Split(ss, ",")
			s.SliceString = arrSlice
		case TokenNumber:
			var n int
			err := dec.Int(&n)
			if err != nil {
				return err
			}
			s.SliceString = []string{strconv.Itoa(n)}
		case TokenObject:
			eb := make(EmbeddedJSON, 0, 128)
			if err := dec.EmbeddedJSON(&eb); err != nil {
				return err
			}

			obj := make(map[string]string)
			err := json.Unmarshal(eb, &obj)
			if err != nil {
				return err
			}

			var arrSlice []string
			for key, value := range obj {
				arrSlice = append(arrSlice, key, value)
			}
			s.SliceString = arrSlice
		case TokenBoolean:
			var n bool
			err := dec.Bool(&n)
			if err != nil {
				return err
			}
			s.SliceString = []string{strconv.FormatBool(n)}
		case TokenNull:
			s.SliceString = []string{}
		}
	}
	return nil
}

func (s *sliceTestObjectWithToken) NKeys() int {
	return 0
}
