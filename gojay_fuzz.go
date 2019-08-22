// +build gofuzz

package gojay

import (
	"bytes"
	"strings"
)

func FuzzUnmarshal(input []byte) int {
	var result *unmarshalTarget
	err := UnmarshalJSONObject(input, result)
	if err != nil {
		return 0
	}
	return 1
}

func FuzzDecode(input []byte) int {
	var result *interface{}
	decoder := NewDecoder(bytes.NewReader(input))
	defer decoder.Release()
	err := decoder.Decode(result)
	if err != nil {
		return 0
	}
	builder := strings.Builder{}
	encoder := NewEncoder(&builder)
	err = encoder.Encode(result)
	if err != nil {
		return 0
	}
	return 1
}

func FuzzStream(input []byte) int {
	var stream *unmarshalStream
	decoder := Stream.NewDecoder(bytes.NewReader(input))
	defer decoder.Release()
	done := decoder.Done()
	for {
		err := decoder.DecodeStream(stream)
		if err != nil {
			return 0
		}
		select {
		case <-done:
			return 1
		default:
		}
	}
}

type unmarshalTarget struct{}

func (*unmarshalTarget) UnmarshalJSONObject(dec *Decoder, key string) error {
	return nil
}

func (*unmarshalTarget) NKeys() int {
	return 0
}

type unmarshalStream struct{}

func (*unmarshalStream) UnmarshalStream(decoder *StreamDecoder) error {
	var result *interface{}
	return decoder.Decode(result)
}
