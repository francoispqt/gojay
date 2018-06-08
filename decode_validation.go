package gojay

func (dec *Decoder) DecodeAndValidate(v interface{}, sch *Schema) error {
	dec.Schema = sch
	dec.validation = 0x1
	dec.path = "."
	return dec.Decode(v)
}
