package gojay

import (
	"io"
)

func NewSchema() *Schema {
	return &Schema{}
}

// Bytes unmarshals the given slice of bytes to the schema
func (s *Schema) Bytes(b []byte) (*Schema, error) {
	err := UnmarshalJSONObject(b, s)
	if err != nil {
		return nil, err
	}
	return parseSchema(s)
}

// Reader unmarshals the given io.Reader to the schema
func (s *Schema) Reader(r io.Reader) (*Schema, error) {
	dec := BorrowDecoder(r)
	defer dec.Release()
	err := dec.DecodeObject(s)
	if err != nil {
		return nil, err
	}
	return parseSchema(s)
}

func parseSchema(s *Schema) (*Schema, error) {
	if s.ItemsJSON != nil {
		if isArray(*s.ItemsJSON) {
			itemsSlice := Schemas{}
			s.ItemsSchemas = &itemsSlice
			err := UnmarshalJSONArray(*s.ItemsJSON, s.ItemsSchemas)
			if err != nil {
				return nil, err
			}
			s.ItemsJSON = nil
			return s, nil
		}
		s.ItemsSchema = &Schema{}
		err := UnmarshalJSONObject(*s.ItemsJSON, s.ItemsSchema)
		if err != nil {
			return nil, err
		}
		s.ItemsJSON = nil
	}
	if s.TypeJSON != nil {
		if isArray(*s.TypeJSON) {
			itemsSlice := Schemas{}
			s.ItemsSchemas = &itemsSlice
			err := UnmarshalJSONArray(*s.TypeJSON, s.ItemsSchemas)
			if err != nil {
				return nil, err
			}
			s.TypeJSON = nil
			return s, nil
		}
		t := ""
		err := Unmarshal(*s.TypeJSON, &t)
		if err != nil {
			return nil, err
		}
		s.Type = t
		s.TypeJSON = nil
	}
	if !s.MinBytes.IsNull() {
		n := 0
		err := Unmarshal(*s.MinBytes, &n)
		if err != nil {
			return nil, err
		}
		s.Min = &n
	}
	if !s.MaxBytes.IsNull() {
		n := 0
		err := Unmarshal(*s.MaxBytes, &n)
		if err != nil {
			return nil, err
		}
		s.Max = &n
	}
	return s, nil
}
