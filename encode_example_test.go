package gojay_test

import (
	"fmt"
	"log"

	"github.com/francoispqt/gojay"
)

func ExampleMarshal_string() {
	str := "gojay"
	d, err := gojay.Marshal(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d)) // "gojay"
}

func ExampleMarshal_bool() {
	b := true
	d, err := gojay.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d)) // true
}
