package fuzz

import (
	"github.com/francoispqt/gojay"
)

type user struct {
	id      int
	created uint64
	age     float64
	name    string
	email   string
	friend  *user
}

// implement gojay.UnmarshalerJSONObject
func (u *user) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.Int(&u.id)
	case "created":
		return dec.Uint64(&u.created)
	case "age":
		return dec.Float(&u.age)
	case "name":
		return dec.String(&u.name)
	case "email":
		return dec.String(&u.email)
	case "friend":
		uu := &user{}
		return dec.Object(uu)
	}
	return nil
}
func (u *user) NKeys() int {
	return 3
}

func Fuzz(input []byte) int {
	u := &user{}
	err := gojay.UnmarshalJSONObject(input, u)
	if err != nil {
		return 0
	}
	return 1
}
