// Package gojay implements encoding and decoding of JSON as defined in RFC 7159.
// The mapping between JSON and Go values is described
// in the documentation for the Marshal and Unmarshal functions.
//
// It aims at performance and usability by relying on simple interfaces
// to decode and encode structures, slices, arrays and even channels.
package gojay
