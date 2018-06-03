package tests 

import "github.com/francoispqt/gojay"

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *BoolSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var b bool
	if err := dec.Bool(&b); err != nil {
		return err
	}
	*v = append(*v, b)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *BoolSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Bool(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *BoolSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *StructSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var s = &A{}
	if err := dec.Object(s); err != nil {
		return err
	}
	*v = append(*v, s)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *StructSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Object(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *StructSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *SliceSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var s = make(StrSlice, 0)
	if err := dec.Array(&s); err != nil {
		return err
	}
	*v = append(*v, &s)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *SliceSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Array(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *SliceSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *StrSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var str string
	if err := dec.String(&str); err != nil {
		return err
	}
	*v = append(*v, str)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *StrSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.String(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *StrSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}

// UnmarshalJSONArray implements gojay's UnmarshalerJSONArray
func (v *IntSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var i int
	if err := dec.Int(&i); err != nil {
		return err
	}
	*v = append(*v, i)
	return nil
}

// MarshalJSONArray implements gojay's MarshalerJSONArray
func (v *IntSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, s := range *v {
		enc.Int(s)
	}
}

// IsNil implements gojay's MarshalerJSONArray
func (v *IntSlice) IsNil() bool {
	return *v == nil || len(*v) == 0
}
