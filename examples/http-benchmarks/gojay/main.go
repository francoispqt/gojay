package main

import (
	"log"
	"net/http"

	"github.com/francoispqt/gojay"
)

func main() {
	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(handler)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body Body
	dec := gojay.BorrowDecoder(r.Body)
	defer dec.Release()
	err := dec.Decode(&body)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	enc := gojay.BorrowEncoder(w)
	defer enc.Release()
	err = enc.Encode(&body)
	if err != nil {
		panic(err)
	}
}

type Body struct {
	Colors *colors `json:"colors"`
}

func (c *Body) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "colors":
		cols := make(colors, 0)
		c.Colors = &cols
		return dec.Array(c.Colors)
	}
	return nil
}
func (b *Body) NKeys() int {
	return 1
}

func (b *Body) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("colors", b.Colors)
}
func (b *Body) IsNil() bool {
	return b == nil
}

type colors []*Color

func (b *colors) UnmarshalJSONArray(dec *gojay.Decoder) error {
	color := &Color{}
	if err := dec.Object(color); err != nil {
		return err
	}
	*b = append(*b, color)
	return nil
}

func (b *colors) MarshalJSONArray(enc *gojay.Encoder) {
	for _, color := range *b {
		enc.Object(color)
	}
}

func (b *colors) IsNil() bool {
	return len(*b) == 0
}

type Color struct {
	Color    string `json:"color,omitempty"`
	Category string `json:"category,omitempty"`
	Type     string `json:"type,omitempty"`
	Code     *Code  `json:"code,omitempty"`
}

func (b *Color) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "color":
		return dec.String(&b.Color)
	case "category":
		return dec.String(&b.Category)
	case "type":
		return dec.String(&b.Type)
	case "code":
		b.Code = &Code{}
		return dec.Object(b.Code)
	}
	return nil
}
func (b *Color) NKeys() int {
	return 4
}

func (b *Color) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ObjectKey("code", b.Code)
	enc.StringKey("color", b.Color)
	enc.StringKey("category", b.Category)
	enc.StringKey("type", b.Type)
}
func (b *Color) IsNil() bool {
	return b == nil
}

type Code struct {
	RGBA *ints  `json:"rgba,omitempty"`
	Hex  string `json:"hex,omitempty"`
}

func (c *Code) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "rgba":
		rgba := make(ints, 0)
		c.RGBA = &rgba
		return dec.Array(&rgba)
	case "hex":
		return dec.String(&c.Hex)
	}
	return nil
}
func (b *Code) NKeys() int {
	return 2
}

func (b *Code) MarshalJSONObject(enc *gojay.Encoder) {
	enc.ArrayKey("rgba", b.RGBA)
	enc.StringKey("hex", b.Hex)
}
func (b *Code) IsNil() bool {
	return b == nil
}

type ints []int

func (v *ints) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i int
	if err := dec.Int(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

func (v *ints) MarshalJSONArray(enc *gojay.Encoder) {
	for _, i := range *v {
		enc.Int(i)
	}
}
func (v *ints) IsNil() bool {
	return v == nil || len(*v) == 0
}
