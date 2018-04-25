[![Go doc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square
)](https://godoc.org/github.com/francoispqt/gojay)
![MIT License](https://img.shields.io/badge/license-mit-blue.svg?style=flat-square)

# GoJay
GoJay is a performant JSON encoder/decoder for Golang (currently the most performant). 

It has a simple API and doesn't use reflection. It relies on small interfaces to decode/encode structures and slices.

Gojay also comes with powerful stream decoding features.

# Get started

```bash
go get github.com/francoispqt/gojay
```

## Decoding

Example of basic stucture decoding:
```go 
import "github.com/francoispqt/gojay"

type user struct {
    id int
    name string
    email string
}
// implement UnmarshalerObject
func (u *user) UnmarshalObject(dec *gojay.Decoder, key string) {
    switch k {
    case "id":
        return dec.AddInt(&u.id)
    case "name":
        return dec.AddString(&u.name)
    case "email":
        return dec.AddString(&u.email)
    }
}
func (u *user) NKeys() int {
    return 3
}

func main() {
    u := &user{}
    d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
    err := gojay.UnmarshalObject(d, user)
    if err != nil {
        log.Fatal(err)
    }
}
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
func (u *user) UnmarshalObject(dec *gojay.Decoder, key string) {
    switch k {
    case "id":
        return dec.AddInt(&u.id)
    case "name":
        return dec.AddString(&u.name)
    case "email":
        return dec.AddString(&u.email)
    }
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

func (c *ChannelArray) UnmarshalArray(dec *gojay.Decoder) error {
	str := ""
	if err := dec.AddString(&str); err != nil {
		return err
	}
	*c <- str
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

func (c *ChannelStream) UnmarshalStream(dec *gojay.StreamDecoder) error {
	obj := &TestObj{}
	if err := dec.AddObject(obj); err != nil {
		return err
	}
	*c <- obj
	return nil
}

func main() {
    // create our channel which will receive our objects
    streamChan := ChannelStream(make(chan *TestObj))
    // get a reader implementing io.Reader
    reader := getAnIOReaderStream()
    dec := gojay.Stream.NewDecoder(reader)
    // start decoding (will block the goroutine until something is written to the ReadWriter)
    go dec.DecodeStream(&streamChan)
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

## Encoding

Example of basic structure encoding:
```go 
import "github.com/francoispqt/gojay"

type user struct {
    id int
    name string
    email string
}
// implement UnmarshalerObject
func (u *user) MarshalObject(dec *gojay.Decoder, key string) {
    dec.AddIntKey("id", u.id)
    dec.AddStringKey("name", u.name)
    dec.AddStringKey("email", u.email)
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
### Arrays and Slices
### Other types

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
| GoJay       | 662   | 112          | 1         |

### Medium Payload
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/decoder/decoder_bench_medium_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_medium.go)

|             | ns/op | bytes/op | allocs/op |
|-------------|-------|--------------|-----------|
| Std Library | 30148 | 2152         | 496       |
| JsonParser  | 7793  | 0            | 0         |
| JsonIter    | 5967  | 496          | 44        |
| GoJay       | 3914  | 128          | 12        |

### Large Payload
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/decoder/decoder_bench_large_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_large.go)

|             | ns/op | bytes/op     | allocs/op |
|-------------|-------|--------------|-----------|
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
| JsonIter    | 866   | 272          | 3         |
| GoJay       | 484   | 320          | 2         |
### Medium Struct
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/encoder/encoder_bench_medium_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_medium.go)

|             | ns/op | bytes/op     | allocs/op |
|-------------|-------|--------------|-----------|
| Std Library | 3325  | 1496         | 18        |
| JsonIter    | 1939  | 648          | 16        |
| GoJay       | 1196  | 936          | 16        |

### Large Struct
[benchmark code is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/encoder/encoder_bench_large_test.go)

[benchmark data is here](https://github.com/francoispqt/gojay/blob/master/benchmarks/benchmarks_large.go)

|             | ns/op | bytes/op     | allocs/op |
|-------------|-------|--------------|-----------|
| Std Library | 51317 | 28704        | 326       |
| JsonIter    | 35247 | 14608        | 320       |
| GoJay       | 27847 | 27888        | 326       |

# Contributing

Contributions are welcome :) 

If you encounter issues please report it in Github and/or send an email at [francois@parquet.ninja](mailto:francois@parquet.ninja)

