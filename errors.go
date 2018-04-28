package gojay

// InvalidJSONError is a type representing an error returned when
// Decoding encounters invalid JSON.
type InvalidJSONError string

func (err InvalidJSONError) Error() string {
	return string(err)
}

// InvalidTypeError is a type representing an error returned when
// Decoding cannot unmarshal JSON to the receiver type for various reasons.
type InvalidTypeError string

func (err InvalidTypeError) Error() string {
	return string(err)
}

const invalidUnmarshalErrorMsg = "Invalid type %s provided to Unmarshal"

// InvalidUnmarshalError is a type representing an error returned when
// Decoding did not find the proper way to decode
type InvalidUnmarshalError string

func (err InvalidUnmarshalError) Error() string {
	return string(err)
}

const invalidMarshalErrorMsg = "Invalid type %s provided to Marshal"

// InvalidMarshalError is a type representing an error returned when
// Encoding did not find the proper way to encode
type InvalidMarshalError string

func (err InvalidMarshalError) Error() string {
	return string(err)
}

// NoReaderError is a type representing an error returned when
// decoding requires a reader and none was given
type NoReaderError string

func (err NoReaderError) Error() string {
	return string(err)
}

// InvalidUsagePooledDecoderError is a type representing an error returned
// when decoding is called on a still pooled Decoder
type InvalidUsagePooledDecoderError string

func (err InvalidUsagePooledDecoderError) Error() string {
	return string(err)
}

// InvalidUsagePooledEncoderError is a type representing an error returned
// when decoding is called on a still pooled Encoder
type InvalidUsagePooledEncoderError string

func (err InvalidUsagePooledEncoderError) Error() string {
	return string(err)
}
