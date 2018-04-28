[![Build Status](https://travis-ci.org/francoispqt/gojay.svg?branch=master)](https://travis-ci.org/francoispqt/gojay)
[![codecov](https://codecov.io/gh/francoispqt/gojay/branch/master/graph/badge.svg)](https://codecov.io/gh/francoispqt/gojay)
[![Go Report Card](https://goreportcard.com/badge/github.com/francoispqt/gojay)](https://goreportcard.com/report/github.com/francoispqt/gojay)
[![Go doc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square
)](https://godoc.org/github.com/francoispqt/gojay)
![MIT License](https://img.shields.io/badge/license-mit-blue.svg?style=flat-square)

# GoJay
**Package is currently at version 0.9.1 and still under development**

GoJay is a performant JSON encoder/decoder for Golang (currently the most performant, [see benchmarks](#benchmark-results)). 

It has a simple API and doesn't use reflection. It relies on small interfaces to decode/encode structures and slices. 

Gojay also comes with powerful stream decoding features and an even faster [Unsafe](#unsafe-api) API.

# Why another JSON parser?

I looked at other fast decoder/encoder and realised it was mostly hardly readable static code generation or a lot of reflection, poor streaming features, and not so fast in the end. 

Also, I wanted to build a decoder that could consume an io.Reader of line or comma delimited JSON, in a JIT way. To consume a flow of JSON objects from a TCP connection for example or from a standard output. Same way I wanted to build an encoder that could encode a flow of data to a io.Writer.

This is how GoJay aims to be a very fast, JIT stream parser with 0 reflection, low allocation with a friendly API.

# Get started

```bash
go get github.com/francoispqt/gojay
```

## Decoding

Decoding is done through two different API similar to standard `encoding/json`:
* [Unmarshal](#unmarshal-api)
* [Decode](#decode-api)


Example of basic stucture decoding with Unmarshal:
```go 
import "github.com/francoispqt/gojay"

type user struct {
    id int
    name string
    email string
}
// implement UnmarshalerObject
func (u *user) UnmarshalObject(dec *gojay.Decoder, key string) error {
    switch key {
    case "id":
        return dec.AddInt(&u.id)
    case "name":
        return dec.AddString(&u.name)
    case "email":
        return dec.AddString(&u.email)
    }
    return nil
}
func (u *user) NKeys() int {
    return 3
}

func main() {
    u := &user{}
    d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
    err := gojay.UnmarshalObject(d, u)
    if err != nil {
        log.Fatal(err)
    }
}
```

with Decode:
```go
func main() {
    u := &user{}
    dec := gojay.NewDecoder(bytes.NewReader([]byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)))
    err := dec.DecodeObject(d, u)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Unmarshal API

Unmarshal API decodes a `[]byte` to a given pointer with a single function.

Behind the doors, Unmarshal API borrows a `*gojay.Decoder` resets its settings and decodes the data to the given pointer and releases the `*gojay.Decoder` to the pool when it finishes, whether it encounters an error or not. 

If it cannot find the right Decoding strategy for the type of the given pointer, it returns an `InvalidUnmarshalError`. You can test the error returned by doing `if ok := err.(InvalidUnmarshalError); ok {}`.

Unmarshal API comes with three functions:
* Unmarshal
```go
func Unmarshal(data []byte, v interface{}) error
```

* UnmarshalObject 
```go
func UnmarshalObject(data []byte, v UnmarshalerObject) error
```

* UnmarshalArray
```go
func UnmarshalArray(data []byte, v UnmarshalerArray) error
```


### Decode API

Decode API decodes a `[]byte` to a given pointer by creating or borrowing a `*gojay.Decoder` with an `io.Reader` and calling `Decode` methods. 

__Getting a *gojay.Decoder or Borrowing__

You can either get a fresh `*gojay.Decoder` calling `dec := gojay.NewDecoder(io.Reader)` or borrow one from the pool by calling `dec := gojay.BorrowDecoder(io.Reader)`.

After using a decoder, you can release it by calling `dec.Release()`. Beware, if you reuse the decoder after releasing it, it will panic with an error of type `InvalidUsagePooledDecoderError`. If you want to fully benefit from the pooling, you must release your decoders after using.

`*gojay.Decoder` has multiple methods to decode to specific types:
* Decode
```go
func (dec *Decoder) DecodeInt(v *int) error
```
* DecodeObject
```go
func (dec *Decoder) DecodeObject(v UnmarshalerObject) error
```
* DecodeArray
```go
func (dec *Decoder) DecodeArray(v UnmarshalerArray) error
```
* DecodeInt
```go
func (dec *Decoder) DecodeInt(v *int) error
```
* DecodeBool
```go
func (dec *Decoder) DecodeBool(v *bool) error
```
* DecodeString
```go
func (dec *Decoder) DecodeString(v *string) error
```


### Structs
#### UnmarshalerObject Interface

To unmarshal a JSON object to a structure, the structure must implement the UnmarshalerObject interface:
```go
type UnmarshalerObject interface {
	UnmarshalObject(*Decoder, string) error
	NKeys() int
}
``` 
UnmarshalObject method takes two arguments, the first one is a pointer to the Decoder (*gojay.Decoder) and the second one is the string value of the current key being parsed. If the JSON data is not an object, the UnmarshalObject method will never be called. 

NKeys method must return the number of keys to Unmarshal in the JSON object. 

Example of implementation: 
```go 
type user struct {
    id int
    name string
    email string
}
// implement UnmarshalerObject
func (u *user) UnmarshalObject(dec *gojay.Decoder, key string) error {
    switch k {
    case "id":
        return dec.AddInt(&u.id)
    case "name":
        return dec.AddString(&u.name)
    case "email":
        return dec.AddString(&u.email)
    }
    return nil
}
func (u *user) NKeys() int {
    return 3
}
```


### Arrays, Slices and Channels

To unmarshal a JSON object to a slice an array or a channel, it must implement the UnmarshalerArray interface:
```go
type UnmarshalerArray interface {
	UnmarshalArray(*Decoder) error
}
```
UnmarshalArray method takes one argument, a pointer to the Decoder (*gojay.Decoder). If the JSON data is not an array, the Unmarshal method will never be called. 

Example of implementation with a slice: 
```go
type testSlice []string
// implement UnmarshalerArray
func (t *testStringArr) UnmarshalArray(dec *gojay.Decoder) error {
	str := ""
	if err := dec.AddString(&str); err != nil {
		return err
	}
	*t = append(*t, str)
	return nil
}
```

Example of implementation with a channel: 
```go
type ChannelString chan string
// implement UnmarshalerArray
func (c ChannelArray) UnmarshalArray(dec *gojay.Decoder) error {
	str := ""
	if err := dec.AddString(&str); err != nil {
		return err
	}
	c <- str
	return nil
}
```

### Stream Decoding
GoJay ships with a powerful stream decoder.

It allows to read continuously from an io.Reader stream and do JIT decoding writing unmarshalled JSON to a channel to allow async consuming. 

When using the Stream API, the Decoder implements context.Context to provide graceful cancellation. 

Example: 
```go
type ChannelStream chan *TestObj
// implement UnmarshalerStream
func (c ChannelStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	obj := &TestObj{}
	if err := dec.AddObject(obj); err != nil {
		return err
	}
	c <- obj
	return nil
}

func main() {
    // create our channel which will receive our objects
    streamChan := ChannelStream(make(chan *TestObj))
    // get a reader implementing io.Reader
    reader := getAnIOReaderStream()
    dec := gojay.Stream.NewDecoder(reader)
    // start decoding (will block the goroutine until something is written to the ReadWriter)
    go dec.DecodeStream(streamChan)
    for {
        select {
        case v := <-streamChan:
            // do something with my TestObj
        case <-dec.Done():
            os.Exit("finished reading stream")
        }
    }
}
```

### Other types
To decode other types (string, int, int32, int64, uint32, uint64, float, booleans), you don't need to implement any interface. 

Example of encoding strings:
```go
func main() {
    json := []byte(`"Jay"`)
    var v string
    err := Unmarshal(json, &v)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(v) // Jay
}
```

## Encoding

Example of basic structure encoding:
```go 
import "github.com/francoispqt/gojay"

type user struct {
    id int
    name string
    email string
}
// implement MarshalerObject
func (u *user) MarshalObject(enc *gojay.Encoder) {
    enc.AddIntKey("id", u.id)
    enc.AddStringKey("name", u.name)
    enc.AddStringKey("email", u.email)
}
func (u *user) IsNil() bool {
    return u == nil
}

func main() {
    u := &user{1, "gojay", "gojay@email.com"}
    b, _ := gojay.MarshalObject(user)
    fmt.Println(string(b)) // {"id":1,"name":"gojay","email":"gojay@email.com"}
}
```

### Structs

To encode a structure, the structure must implement the MarshalerObject interface:
```go
type MarshalerObject interface {
	MarshalObject(enc *Encoder)
	IsNil() bool
}
```
MarshalObject method takes one argument, a pointer to the Encoder (*gojay.Encoder). The method must add all the keys in the JSON Object by calling Decoder's methods. 

IsNil method returns a boolean indicating if the interface underlying value is nil or not. It is used to safely ensure that the underlying value is not nil without using Reflection. 

Example of implementation: 
```go 
type user struct {
    id int
    name string
    email string
}
// implement MarshalerObject
func (u *user) MarshalObject(dec *gojay.Decoder, key string) {
    dec.AddIntKey("id", u.id)
    dec.AddStringKey("name", u.name)
    dec.AddStringKey("email", u.email)
}
func (u *user) IsNil() bool {
    return u == nil
}
```

### Arrays and Slices
To encode an array or a slice, the slice/array must implement the MarshalerArray interface:
```go
type MarshalerArray interface {
	MarshalArray(enc *Encoder)
}
```
MarshalArray method takes one argument, a pointer to the Encoder (*gojay.Encoder). The method must add all element in the JSON Array by calling Decoder's methods. 

Example of implementation: 
```go
type users []*user
// implement MarshalerArray
func (u *users) MarshalArray(dec *Decoder) error {
	for _, e := range u {
        err := enc.AddObject(e)
        if err != nil {
            return err
        }
    }
    return nil
}
```

### Other types
To encode other types (string, int, float, booleans), you don't need to implement any interface. 

Example of encoding strings:
```go
func main() {
    name := "Jay"
    b, err := gojay.Marshal(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(b)) // "Jay"
}
```

# Unsafe API

Unsafe API has the same functions than the regular API, it only has `Unmarshal API` for now. It is unsafe because it makes assumptions on the quality of the given JSON. 

If you are not sure if you're JSON is valid, don't use the Unsafe API. 

Also, the `Unsafe` API does not copy the buffer when using Unmarshal API, which, in case of string decoding, can lead to data corruption if a byte buffer is reused. Using the `Decode` API makes `Unsafe` API safer as the io.Reader relies on `copy` builtin method and `Decoder` will have its own internal buffer :) 

Access the `Unsafe` API this way:
```go
gojay.Unsafe.Unmarshal(b, v) 
```


# Benchmarks

Benchmarks encode and decode three different data based on size (small, medium, large). 

To run benchmark for decoder:
```bash
cd $GOPATH/src/github.com/francoispqt/gojay/benchmarks/decoder && make bench
```

To run benchmark for encoder:
```bash
cd $GOPATH/src/github.com/francoispqt/gojay/benchmarks/encoder && make bench
```

# Benchmark Results
## Decode

<img src="https://images2.imgbox.com/78/01/49OExcPh_o.png" width="500px">

### Small Payload
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/decoder/decoder_bench_small_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_small.go)

|             | ns/op | bytes/op | allocs/op |
|-------------|-------|--------------|-----------|
| Std Library | 4661  | 496          | 12        |
| JsonParser  | 1313  | 0            | 0         |
| JsonIter    | 899   | 192          | 5         |
| EasyJson    | 929   | 240          | 2         |
| GoJay       | 662   | 112          | 1         |

### Medium Payload
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/decoder/decoder_bench_medium_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_medium.go)

|             | ns/op | bytes/op | allocs/op |
|-------------|-------|--------------|-----------|
| Std Library | 30148 | 2152         | 496       |
| JsonParser  | 7793  | 0            | 0         |
| EasyJson    | 7957  | 232          | 6         |
| JsonIter    | 5967  | 496          | 44        |
| GoJay       | 3914  | 128          | 7         |

### Large Payload
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/decoder/decoder_bench_large_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_large.go)

|             | ns/op | bytes/op     | allocs/op |
|-------------|-------|--------------|-----------|
| EasyJson    | 106626| 160          | 2         |
| JsonParser  | 66813 | 0            | 0         |
| JsonIter    | 87994 | 6738         | 329       |
| GoJay       | 43402 | 1408         | 76        |

## Encode

<img src="https://images2.imgbox.com/e9/cc/pnM8c7Gf_o.png" width="500px">

### Small Struct
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/encoder/encoder_bench_small_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_small.go)

|             | ns/op | bytes/op     | allocs/op |
|-------------|-------|--------------|-----------|
| Std Library | 1280  | 464          | 3         |
| EasyJson    | 871   | 944          | 6         |
| JsonIter    | 866   | 272          | 3         |
| GoJay       | 484   | 320          | 2         |

### Medium Struct
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/encoder/encoder_bench_medium_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_medium.go)

|             | ns/op | bytes/op     | allocs/op |
|-------------|-------|--------------|-----------|
| Std Library | 3325  | 1496         | 18        |
| EasyJson    | 1997  | 1320         | 19        |
| JsonIter    | 1939  | 648          | 16        |
| GoJay       | 1196  | 936          | 16        |

### Large Struct
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/encoder/encoder_bench_large_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_large.go)

|             | ns/op | bytes/op     | allocs/op |
|-------------|-------|--------------|-----------|
| Std Library | 51317 | 28704        | 326       |
| JsonIter    | 35247 | 14608        | 320       |
| EasyJson    | 32053 | 15474        | 327       |
| GoJay       | 27847 | 27888        | 326       |

# Contributing

Contributions are welcome :) 

If you encounter issues please report it in Github and/or send an email at [francois@parquet.ninja](mailto:francois@parquet.ninja)

