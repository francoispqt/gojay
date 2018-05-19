package main

import (
	"net/http"

	"github.com/francoispqt/gojay"
)

type message struct {
	foo string
	bar string
}

func (m *message) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "foo":
		return dec.AddString(&m.foo)
	case "bar":
		return dec.AddString(&m.bar)
	}
	return nil
}

func (m *message) NKeys() int {
	return 2
}

func (m *message) MarshalJSONObject(dec *gojay.Encoder) {
	dec.AddStringKey("foo", m.foo)
	dec.AddStringKey("bar", m.bar)
}

func (m *message) IsNil() bool {
	return m == nil
}

func home(w http.ResponseWriter, r *http.Request) {
	// read body using io.Reader
	m := &message{}
	dec := gojay.BorrowDecoder(r.Body)
	defer dec.Release()
	err := dec.DecodeObject(m)
	if err != nil {
		i, err := w.Write([]byte(err.Error()))
		if err != nil || i == 0 {
			panic(err)
		}
		return
	}

	// just transform response slightly
	m.foo += "hey"

	// return response using io.Writer
	enc := gojay.BorrowEncoder(w)
	defer enc.Release()
	err = enc.Encode(m)
	if err != nil {
		i, err := w.Write([]byte(err.Error()))
		if err != nil || i == 0 {
			panic(err)
		}
	}
	return
}

func main() {
	http.HandleFunc("/", home)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
