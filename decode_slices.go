package gojay

// SliceString is a *[]string implementing gojay.UnmarshalerJSONArray interface
type SliceString []string

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray
func (s *SliceString) UnmarshalJSONArray(dec *Decoder) error {
	var str string
	if err := dec.String(&str); err != nil {
		return err
	}
	*s = append(*s, str)
	return nil
}

// IsNil implements gojay.UnmarshalerJSONArray
func (s *SliceString) IsNil() bool {
	return s == nil || len(*s) == 0
}

// AddSliceString unmarshals the next JSON array of strings to the given *[]string s
func (dec *Decoder) AddSliceString(s *[]string) error {
	return dec.SliceString(s)
}

// SliceString unmarshals the next JSON array of strings to the given *[]string s
func (dec *Decoder) SliceString(s *[]string) error {
	var st = SliceString(*s)
	if err := dec.Array(&st); err != nil {
		return err
	}
	*s = st
	return nil
}

// SliceInt is a *[]int implementing gojay.UnmarshalerJSONArray interface
type SliceInt []int

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray
func (s *SliceInt) UnmarshalJSONArray(dec *Decoder) error {
	var i int
	if err := dec.Int(&i); err != nil {
		return err
	}
	*s = append(*s, i)
	return nil
}

// IsNil implements gojay.UnmarshalerJSONArray
func (s *SliceInt) IsNil() bool {
	return s == nil || len(*s) == 0
}

// AddInt unmarshals the next JSON array of integers to the given *[]int s
func (dec *Decoder) AddSliceInt(s *[]int) error {
	return dec.SliceInt(s)
}

// SliceInt unmarshals the next JSON array of integers to the given *[]int s
func (dec *Decoder) SliceInt(s *[]int) error {
	var st = SliceInt(*s)
	if err := dec.Array(&st); err != nil {
		return err
	}
	*s = st
	return nil
}

// SliceFloat64 is a *[]float64 implementing gojay.UnmarshalerJSONArray interface
type SliceFloat64 []float64

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray
func (s *SliceFloat64) UnmarshalJSONArray(dec *Decoder) error {
	var i float64
	if err := dec.Float64(&i); err != nil {
		return err
	}
	*s = append(*s, i)
	return nil
}

// IsNil implements gojay.UnmarshalerJSONArray
func (s *SliceFloat64) IsNil() bool {
	return s == nil || len(*s) == 0
}

// AddFloat64 unmarshals the next JSON array of floats to the given *[]float64 s
func (dec *Decoder) AddSliceFloat64(s *[]float64) error {
	return dec.SliceFloat64(s)
}

// SliceFloat64 unmarshals the next JSON array of floats to the given *[]float64 s
func (dec *Decoder) SliceFloat64(s *[]float64) error {
	var st = SliceFloat64(*s)
	if err := dec.Array(&st); err != nil {
		return err
	}
	*s = st
	return nil
}

// SliceBool is a *[]bool implementing gojay.UnmarshalerJSONArray boolerface
type SliceBool []bool

// UnmarshalJSONArray implements gojay.UnmarshalerJSONArray
func (s *SliceBool) UnmarshalJSONArray(dec *Decoder) error {
	var i bool
	if err := dec.Bool(&i); err != nil {
		return err
	}
	*s = append(*s, i)
	return nil
}

// IsNil implements gojay.UnmarshalerJSONArray
func (s *SliceBool) IsNil() bool {
	return s == nil || len(*s) == 0
}

// AddBool unmarshals the next JSON array of boolegers to the given *[]bool s
func (dec *Decoder) AddSliceBool(s *[]bool) error {
	return dec.SliceBool(s)
}

// SliceBool unmarshals the next JSON array of boolegers to the given *[]bool s
func (dec *Decoder) SliceBool(s *[]bool) error {
	var st = SliceBool(*s)
	if err := dec.Array(&st); err != nil {
		return err
	}
	*s = st
	return nil
}
