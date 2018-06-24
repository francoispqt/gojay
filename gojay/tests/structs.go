package tests

//gojay:json
type A struct {
	Str    string `gojay:"string"`
	Bool   bool
	Int    int
	Int64  int64
	Int32  int32
	Int16  int16
	Int8   int8
	Uint64 uint64
	Uint32 uint32
	Uint16 uint16
	Uint8  uint8
	Bval   *B
	Arrval *StrSlice
}

//gojay:json
type B struct {
	Str       string
	Bool      bool
	Int       int
	Int64     int64
	Int32     int32
	Int16     int16
	Int8      int8
	Uint64    uint64
	Uint32    uint32
	Uint16    uint16
	Uint8     uint8
	StrPtr    *string
	BoolPtr   *bool
	IntPtr    *int
	Int64Ptr  *int64
	Int32Ptr  *int32
	Int16Ptr  *int16
	Int8Ptr   *int8
	Uint64Ptr *uint64
	Uint32Ptr *uint32
	Uint16Ptr *uint16
	Uint8PTr  *uint8
}
