package gojay

import "database/sql"

// DecodeSQLNullString decodes a sql.NullString
func (dec *Decoder) DecodeSQLNullString(v *sql.NullString) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeSQLNullString(v)
}

func (dec *Decoder) decodeSQLNullString(v *sql.NullString) error {
	var str string
	if err := dec.decodeString(&str); err != nil {
		return err
	}
	v.String = str
	v.Valid = true
	return nil
}

// DecodeSQLNullInt64 decodes a sql.NullInt64
func (dec *Decoder) DecodeSQLNullInt64(v *sql.NullInt64) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeSQLNullInt64(v)
}

func (dec *Decoder) decodeSQLNullInt64(v *sql.NullInt64) error {
	var i int64
	if err := dec.decodeInt64(&i); err != nil {
		return err
	}
	v.Int64 = i
	v.Valid = true
	return nil
}

// DecodeSQLNullFloat64 decodes a sql.NullString with the given format
func (dec *Decoder) DecodeSQLNullFloat64(v *sql.NullFloat64) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeSQLNullFloat64(v)
}

func (dec *Decoder) decodeSQLNullFloat64(v *sql.NullFloat64) error {
	var i float64
	if err := dec.decodeFloat64(&i); err != nil {
		return err
	}
	v.Float64 = i
	v.Valid = true
	return nil
}

// DecodeSQLNullBool decodes a sql.NullString with the given format
func (dec *Decoder) DecodeSQLNullBool(v *sql.NullBool) error {
	if dec.isPooled == 1 {
		panic(InvalidUsagePooledDecoderError("Invalid usage of pooled decoder"))
	}
	return dec.decodeSQLNullBool(v)
}

func (dec *Decoder) decodeSQLNullBool(v *sql.NullBool) error {
	var b bool
	if err := dec.decodeBool(&b); err != nil {
		return err
	}
	v.Bool = b
	v.Valid = true
	return nil
}
