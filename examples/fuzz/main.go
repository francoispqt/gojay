package fuzz

import (
	"github.com/francoispqt/gojay"
)

type user struct {
	id    int
	name  string
	email string
}

// implement gojay.UnmarshalerJSONObject
func (u *user) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.Int(&u.id)
	case "name":
		return dec.String(&u.name)
	case "email":
		return dec.String(&u.email)
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
