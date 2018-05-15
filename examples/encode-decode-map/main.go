package main

import (
	"log"
	"strings"

	"github.com/francoispqt/gojay"
)

// define our custom map type implementing MarshalerJSONObject and UnmarshalerJSONObject
type myMap map[string]string

// Implementing Unmarshaler
func (m myMap) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	str := ""
	err := dec.AddString(&str)
	if err != nil {
		return err
	}
	m[k] = str
	return nil
}

// Her we put the number of keys
// If number of keys is unknown return 0, it will parse all keys
func (m myMap) NKeys() int {
	return 0
}

// Implementing Marshaler
func (m myMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range m {
		enc.AddStringKey(k, v)
	}
}

func (m myMap) IsNil() bool {
	return m == nil
}

// Using Marshal / Unmarshal API
func marshalAPI(m myMap) error {
	b, err := gojay.Marshal(m)
	if err != nil {
		return err
	}
	log.Print(string(b))

	nM := myMap(make(map[string]string))
	err = gojay.Unmarshal(b, nM)
	if err != nil {
		return err
	}
	log.Print(nM)
	return nil
}

// Using Encode / Decode API
func encodeAPI(m myMap) error {
	// we use strings.Builder as it implements io.Writer
	builder := &strings.Builder{}
	enc := gojay.BorrowEncoder(builder)
	defer enc.Release()
	// encode
	err := enc.EncodeObject(m)
	if err != nil {
		return err
	}
	log.Print(builder.String())

	// make our new map which will receive the decoded JSON
	nM := myMap(make(map[string]string))
	// get our decoder with an io.Reader
	dec := gojay.BorrowDecoder(strings.NewReader(builder.String()))
	defer dec.Release()
	// decode
	err = dec.DecodeObject(nM)
	if err != nil {
		return err
	}
	log.Print(nM)
	return nil
}

func main() {
	// make our map to be encoded
	m := myMap(map[string]string{
		"test":  "test",
		"test2": "test2",
	})

	err := marshalAPI(m)
	if err != nil {
		log.Fatal(err)
	}
	err = encodeAPI(m)
	if err != nil {
		log.Fatal(err)
	}
}
