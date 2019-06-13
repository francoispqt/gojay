package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"runtime/debug"
	"time"

	"github.com/francoispqt/gojay"
)

func main() {
	debug.SetGCPercent(5)

	b1 := make([]byte, 0, 512)
	b2 := make([]byte, 0, 512)
	r := bytes.NewBuffer(b1)
	w := bytes.NewBuffer(b2)

	rw := bufio.NewReadWriter(bufio.NewReader(r), bufio.NewWriter(w))
	dec := gojay.NewDecoder(rw)

	go write(r)
	go read(dec, rw)

	c := make(chan struct{})
	<-c
}

func write(w io.Writer) {
	for {
		w.Write([]byte(`{"foo":"bar"}`))
		time.Sleep(1 * time.Millisecond)
	}
}

func read(dec *gojay.Decoder, r io.Reader) {
	var msg gojay.EmbeddedJSON
	for {
		msg = msg[:0]
		err := dec.Decode(&msg)
		if err != nil {
			dec = gojay.NewDecoder(r)
			fmt.Println(err)
		} else {
			fmt.Println(string(msg))
		}
		time.Sleep(1 * time.Millisecond)
	}
}
