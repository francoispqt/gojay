package gojay

import (
	"encoding/base64"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeToBase64(t *testing.T) {
	type testCase struct {
		source       io.Reader
		baseJSON     string
		expectedJSON string
	}

	f := func(tc testCase) {
		t.Helper()

		var b strings.Builder
		var enc = NewEncoder(&b)
		enc.writeString(tc.baseJSON)
		enc.AddReaderToBase64(tc.source, base64.StdEncoding)
		enc.Write()
		assert.Equal(t, tc.expectedJSON, b.String())
	}

	f(testCase{
		source:       strings.NewReader("some long string to be encoded\x00\x01\x02\x03"),
		baseJSON:     `[`,
		expectedJSON: `["c29tZSBsb25nIHN0cmluZyB0byBiZSBlbmNvZGVkAAECAw=="`,
	})
	f(testCase{
		source:       strings.NewReader("some long string to be encoded\x00\x01\x02\x03"),
		baseJSON:     `["aaa",123`,
		expectedJSON: `["aaa",123,"c29tZSBsb25nIHN0cmluZyB0byBiZSBlbmNvZGVkAAECAw=="`,
	})
}

func TestEncodeToBase64Key(t *testing.T) {
	type testCase struct {
		source       io.Reader
		baseJSON     string
		expectedJSON string
	}

	f := func(tc testCase) {
		t.Helper()

		var b strings.Builder
		var enc = NewEncoder(&b)
		enc.writeString(tc.baseJSON)
		enc.AddReaderToBase64Key("key", tc.source, base64.StdEncoding)
		enc.Write()
		assert.Equal(t, tc.expectedJSON, b.String())
	}

	f(testCase{
		source:       strings.NewReader("some long string to be encoded\x00\x01\x02\x03"),
		baseJSON:     `{`,
		expectedJSON: `{"key":"c29tZSBsb25nIHN0cmluZyB0byBiZSBlbmNvZGVkAAECAw=="`,
	})
	f(testCase{
		source:       strings.NewReader("some long string to be encoded\x00\x01\x02\x03"),
		baseJSON:     `{"a":"b"`,
		expectedJSON: `{"a":"b","key":"c29tZSBsb25nIHN0cmluZyB0byBiZSBlbmNvZGVkAAECAw=="`,
	})
}

func TestEncodeToEscaped(t *testing.T) {
	type testCase struct {
		source       io.Reader
		baseJSON     string
		expectedJSON string
	}

	f := func(tc testCase) {
		t.Helper()

		var b strings.Builder
		var enc = NewEncoder(&b)
		enc.writeString(tc.baseJSON)
		enc.AddReaderToEscaped(tc.source)
		enc.Write()
		assert.Equal(t, tc.expectedJSON, b.String())
	}

	f(testCase{
		source:       strings.NewReader("some long string to be encoded\x00\x01\x02\x03テュールスト マ\\ーテ\nィン ヤコブ 😁\t"),
		baseJSON:     `[`,
		expectedJSON: `["some long string to be encoded\u0000\u0001\u0002\u0003テュールスト マ\\ーテ\nィン ヤコブ 😁\t"`,
	})
	f(testCase{
		source:       strings.NewReader("some long string to be encoded\x00\x01\x02\x03テュールスト マ\\ーテ\nィン ヤコブ 😁\t"),
		baseJSON:     `["aaa",123`,
		expectedJSON: `["aaa",123,"some long string to be encoded\u0000\u0001\u0002\u0003テュールスト マ\\ーテ\nィン ヤコブ 😁\t"`,
	})
}

func TestEncodeToEscapedKey(t *testing.T) {
	type testCase struct {
		source       io.Reader
		baseJSON     string
		expectedJSON string
	}

	f := func(tc testCase) {
		t.Helper()

		var b strings.Builder
		var enc = NewEncoder(&b)
		enc.writeString(tc.baseJSON)
		enc.AddReaderToEscapedKey("key", tc.source)
		enc.Write()
		assert.Equal(t, tc.expectedJSON, b.String())
	}

	f(testCase{
		source:       strings.NewReader("some long string to be encoded\x00\x01\x02\x03テュールスト マ\\ーテ\nィン ヤコブ 😁\t"),
		baseJSON:     `{`,
		expectedJSON: `{"key":"some long string to be encoded\u0000\u0001\u0002\u0003テュールスト マ\\ーテ\nィン ヤコブ 😁\t"`,
	})
	f(testCase{
		source:       strings.NewReader("some long string to be encoded\x00\x01\x02\x03テュールスト マ\\ーテ\nィン ヤコブ 😁\t"),
		baseJSON:     `{"a":"b"`,
		expectedJSON: `{"a":"b","key":"some long string to be encoded\u0000\u0001\u0002\u0003テュールスト マ\\ーテ\nィン ヤコブ 😁\t"`,
	})
}
