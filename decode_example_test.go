package gojay_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/francoispqt/gojay"
)

func ExampleUnmarshal_string() {
	data := []byte(`"gojay"`)
	var str string
	err := gojay.Unmarshal(data, &str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str) // true
}

func ExampleUnmarshal_bool() {
	data := []byte(`true`)
	var b bool
	err := gojay.Unmarshal(data, &b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b) // true
}

func ExampleUnmarshal_invalidType() {
	data := []byte(`"gojay"`)
	someStruct := struct{}{}
	err := gojay.Unmarshal(data, &someStruct)

	fmt.Println(err) // "Cannot unmarshal JSON to type '*struct{}'"
}

func ExampleDecode_string() {
	var str string
	dec := gojay.BorrowDecoder(strings.NewReader(`"gojay"`))
	err := dec.Decode(&str)
	dec.Release()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str) // "gojay"
}
