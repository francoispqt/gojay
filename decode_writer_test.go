package gojay

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var randStringLetters = []rune("abcdefABCDEF'\"\n\t\r\b\fテュールストマ\\ーティンヤコブ 😁𠜎𠜱𠝹𠱓𠱸𠲖𠳏𠳕𠴕𠵼𠵿𝄞ۼ")

type PayloadForDecoding struct {
	Bool         bool
	Base64Stream bytes.Buffer
	Int          int
	StrStream    bytes.Buffer
	Float        float64
}

func (p *PayloadForDecoding) UnmarshalJSONObject(dec *Decoder, key string) error {
	switch key {
	case "bool":
		return dec.Bool(&p.Bool)
	case "base64_stream":
		return dec.WriterFromBase64(&p.Base64Stream, base64.StdEncoding)
	case "int":
		return dec.Int(&p.Int)
	case "str_stream":
		return dec.WriterFromEscaped(&p.StrStream)
	case "float":
		return dec.Float64(&p.Float)
	default:
		return nil
	}
}

func (p *PayloadForDecoding) NKeys() int { return 5 }

func makeRandString(rng *rand.Rand, size int) string {
	var sb strings.Builder
	for sb.Len() < size {
		idx := rng.Int63() % int64(len(randStringLetters))
		sb.WriteRune(randStringLetters[idx])
	}
	return sb.String()
}

func TestDecodeWriterEscapedString(t *testing.T) {
	type testCase struct {
		name            string
		json            string
		expectedPayload PayloadForDecoding
		expectedError   error
	}

	f := func(tc testCase) {
		t.Helper()

		var payload PayloadForDecoding
		decoder := BorrowDecoder(strings.NewReader(tc.json))
		defer decoder.Release()

		err := decoder.Decode(&payload)

		if tc.expectedError != nil {
			assert.IsType(t, tc.expectedError, err)
			return
		}
		if assert.NoError(t, err) {
			assert.Equal(t, tc.expectedPayload, payload)
		}
	}

	f(testCase{
		name: "basic string",
		json: `{"int":2,"str_stream":"hello world","bool":true}`,
		expectedPayload: PayloadForDecoding{
			Int:       2,
			StrStream: *bytes.NewBufferString("hello world"),
			Bool:      true,
		},
	})
	f(testCase{
		name: "empty string",
		json: `{"int":1,"float":2,"str_stream":""}`,
		expectedPayload: PayloadForDecoding{
			Int:       1,
			Float:     2,
			StrStream: *bytes.NewBufferString(""),
		},
	})
	f(testCase{
		name: "escaped solidus",
		json: `{"str_stream":"\/","int":1,"float":2}`,
		expectedPayload: PayloadForDecoding{
			Int:       1,
			Float:     2,
			StrStream: *bytes.NewBufferString("/"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\\/","int":1,"float":2,}`,
		expectedPayload: PayloadForDecoding{
			Int:       1,
			Float:     2,
			StrStream: *bytes.NewBufferString("\\/"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\\t"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString("\\t"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\t"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString("\t"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\\b"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString("\\b"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\b"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString("\b"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\\f"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString("\\f"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\f"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString("\f"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\\r"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString("\\r"),
		},
	})
	f(testCase{
		json: `{"str_stream":"\r"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString("\r"),
		},
	})
	f(testCase{
		json: `{"str_stream":"𠜎 𠜱 𠝹 𠱓 𠱸 𠲖 𠳏 𠳕 𠴕 𠵼 𠵿"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString(`𠜎 𠜱 𠝹 𠱓 𠱸 𠲖 𠳏 𠳕 𠴕 𠵼 𠵿`),
		},
	})
	f(testCase{
		json: `{"str_stream":"\u06fc","bool":false,}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString(`ۼ`),
		},
	})
	f(testCase{
		json: `{"bool":false,"str_stream":"\\u2070"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBufferString(`\u2070`),
		},
	})
	f(testCase{
		json: `{"str_stream":"\uD834\uDD1E","int":10}`,
		expectedPayload: PayloadForDecoding{
			Int:       10,
			StrStream: *bytes.NewBufferString(`𝄞`),
		},
	})
	f(testCase{
		json: `{"bool":true,"str_stream":"\uD834\\","int":11}`,
		expectedPayload: PayloadForDecoding{
			Int:       11,
			StrStream: *bytes.NewBufferString(`�\`),
			Bool:      true,
		},
	})
	f(testCase{
		json: `{"str_stream":"\uD834\uD834","int":11}`,
		expectedPayload: PayloadForDecoding{
			Int:       11,
			StrStream: *bytes.NewBufferString("�\x00\x00\x00"),
		},
	})
	f(testCase{
		json: `{"float":11,"str_stream":"\uD834"}`,
		expectedPayload: PayloadForDecoding{
			Float:     11,
			StrStream: *bytes.NewBufferString("�"),
		},
	})
	f(testCase{
		json:          `{"str_stream":"\u2Z80"}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json:          `{"float":11,"str_stream":"\uD834\"}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json:          `{"str_stream":"\uD834\uDZ1E"}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json:          `{"str_stream":"\uD834}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json: `{"int": 1, "str_stream":"\uD834\t", "key":2}`,
		expectedPayload: PayloadForDecoding{
			Int:       1,
			StrStream: *bytes.NewBufferString("�\t"),
		},
	})
	f(testCase{
		json: `{"int": 1, "str_stream":"\uD834\n", "key":2}`,
		expectedPayload: PayloadForDecoding{
			Int:       1,
			StrStream: *bytes.NewBufferString("�\n"),
		},
	})
	f(testCase{
		json: `{"int": 1, "str_stream":"\uD834\f", "key":2}`,
		expectedPayload: PayloadForDecoding{
			Int:       1,
			StrStream: *bytes.NewBufferString("�\f"),
		},
	})
	f(testCase{
		json: `{"int": 1, "str_stream":"\uD834\b", "key":2}`,
		expectedPayload: PayloadForDecoding{
			Int:       1,
			StrStream: *bytes.NewBufferString("�\b"),
		},
	})
	f(testCase{
		json: `{"int": 1, "str_stream":"\uD834\r", "key":2}`,
		expectedPayload: PayloadForDecoding{
			Int:       1,
			StrStream: *bytes.NewBufferString("�\r"),
		},
	})
	f(testCase{
		json:          `{"int": 1, "str_stream":"\uD834\h", "key":2}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json: `{"str_stream":null, "aaa":"bb"}`,
		expectedPayload: PayloadForDecoding{
			StrStream: *bytes.NewBuffer(nil),
		},
	})
	f(testCase{
		json:          `{"str_stream":nall, "aaa":"bb"}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json: `{"key":null,"str_stream":"test string \" escaped", "float":1010.123}`,
		expectedPayload: PayloadForDecoding{
			Float:     1010.123,
			StrStream: *bytes.NewBufferString("test string \" escaped"),
		},
	})
	f(testCase{
		json: `{"key":null,"str_stream":"test string \t escaped", "float":1010.123}`,
		expectedPayload: PayloadForDecoding{
			Float:     1010.123,
			StrStream: *bytes.NewBufferString("test string \t escaped"),
		},
	})
	f(testCase{
		json: `{"key":null,"str_stream":"test string \r escaped", "float":1010.123}`,
		expectedPayload: PayloadForDecoding{
			Float:     1010.123,
			StrStream: *bytes.NewBufferString("test string \r escaped"),
		},
	})
	f(testCase{
		json: `{"key":null,"str_stream":"test string \b escaped", "float":1010.123}`,
		expectedPayload: PayloadForDecoding{
			Float:     1010.123,
			StrStream: *bytes.NewBufferString("test string \b escaped"),
		},
	})
	f(testCase{
		json: `{"key":null,"str_stream":"test string \n escaped", "float":1010.123}`,
		expectedPayload: PayloadForDecoding{
			Float:     1010.123,
			StrStream: *bytes.NewBufferString("test string \n escaped"),
		},
	})
	f(testCase{
		json:          `{"str_stream":"test string \\\" escaped, "int":1010}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json:          `{"str_stream":"test string \\\l escaped", "int":1010}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json:          `{"str_stream":invalid, "int":1010}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		json: `{"float":1.2,"str_stream":"string with spaces and \"escape\"d \"quotes\" and escaped line returns \n and escaped \\\\ escaped char","int":2}`,
		expectedPayload: PayloadForDecoding{
			Float:     1.2,
			StrStream: *bytes.NewBufferString("string with spaces and \"escape\"d \"quotes\" and escaped line returns \n and escaped \\\\ escaped char"),
			Int:       2,
		},
	})
}

func TestDecodeWriterEscapedString_multiple_reads(t *testing.T) {
	randStr := makeRandString(rand.New(rand.NewSource(time.Now().Unix())), 100*1024) // 100KiB
	jsonStr, err := json.Marshal(randStr)
	assert.NoError(t, err)

	decoder := BorrowDecoder(strings.NewReader(`{"int":1,"bool":true,"str_stream":` + string(jsonStr) + `,"float":2.1}`))
	defer decoder.Release()

	var payload PayloadForDecoding
	if assert.NoError(t, decoder.DecodeObject(&payload)) {
		assert.Equal(t, PayloadForDecoding{
			Bool:      true,
			Int:       1,
			StrStream: *bytes.NewBufferString(randStr),
			Float:     2.1,
		}, payload)
	}
}

func TestDecodeWriterBase64(t *testing.T) {
	type testCase struct {
		name            string
		json            string
		expectedPayload PayloadForDecoding
		expectedError   error
	}

	f := func(tc testCase) {
		t.Helper()

		var payload PayloadForDecoding
		decoder := BorrowDecoder(strings.NewReader(tc.json))
		defer decoder.Release()

		err := decoder.Decode(&payload)

		if tc.expectedError != nil {
			assert.IsType(t, tc.expectedError, err)
			return
		}
		if assert.NoError(t, err) {
			assert.Equal(t, tc.expectedPayload, payload)
		}
	}

	f(testCase{
		name: "basic (at the end)",
		json: `{"int":10,"base64_stream": "YWttZGxzYXNsbWR5amtsLGttam5oYmdoamtsLGtqbg=="}`,
		expectedPayload: PayloadForDecoding{
			Int:          10,
			Base64Stream: *bytes.NewBufferString("akmdlsaslmdyjkl,kmjnhbghjkl,kjn"),
		},
	})
	f(testCase{
		name: "basic (at the begin)",
		json: `{"base64_stream": "YWttZGxzYXNsbWR5amtsLGttam5oYmdoamtsLGtqbg==","float":2.1}`,
		expectedPayload: PayloadForDecoding{
			Float:        2.1,
			Base64Stream: *bytes.NewBufferString("akmdlsaslmdyjkl,kmjnhbghjkl,kjn"),
		},
	})
	f(testCase{
		name: "basic (at the middle)",
		json: `{"bool":true,"base64_stream": "YWttZGxzYXNsbWR5amtsLGttam5oYmdoamtsLGtqbg==","float":2.1}`,
		expectedPayload: PayloadForDecoding{
			Float:        2.1,
			Base64Stream: *bytes.NewBufferString("akmdlsaslmdyjkl,kmjnhbghjkl,kjn"),
			Bool:         true,
		},
	})
	f(testCase{
		name: "empty",
		json: `{"bool":true,"base64_stream": "","float":2.1}`,
		expectedPayload: PayloadForDecoding{
			Float:        2.1,
			Base64Stream: *bytes.NewBufferString(""),
			Bool:         true,
		},
	})
	f(testCase{
		name: "null",
		json: `{"bool":true,"base64_stream": null}`,
		expectedPayload: PayloadForDecoding{
			Base64Stream: *bytes.NewBuffer(nil),
			Bool:         true,
		},
	})
	f(testCase{
		name:          "invalid json",
		json:          `{"bool":true,"base64_stream": nall}`,
		expectedError: InvalidJSONError(""),
	})
	f(testCase{
		name:          "invalid json",
		json:          `{"base64_stream": "A ,"aaa":"bcd"}`,
		expectedError: io.ErrUnexpectedEOF, // base64 decoder converts InvalidJSONError to this
	})
	f(testCase{
		name:          "invalid character",
		json:          `{"base64_stream": "Ax" ,"aaa":"bcd"}`,
		expectedError: io.ErrUnexpectedEOF,
	})
	f(testCase{
		name:          "invalid padding",
		json:          `{"base64_stream": "YWttZGxzYXNsbWR5amtsLGttam5oYmdoamtsLGtqbg=" ,"aaa":"bcd"}`,
		expectedError: io.ErrUnexpectedEOF,
	})
}

func TestDecodeWriterBase64_multiple_reads(t *testing.T) {
	randBytes := make([]byte, 100*1024) // 100KiB
	rand.New(rand.NewSource(time.Now().Unix())).Read(randBytes)

	decoder := BorrowDecoder(strings.NewReader(`{"base64_stream":"` + base64.StdEncoding.EncodeToString(randBytes) + `","int":20}`))
	defer decoder.Release()

	var payload PayloadForDecoding
	if assert.NoError(t, decoder.DecodeObject(&payload)) {
		assert.Equal(t, PayloadForDecoding{
			Int:          20,
			Base64Stream: *bytes.NewBuffer(randBytes),
		}, payload)
	}
}
