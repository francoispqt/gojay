# Gojay code generator

This package provides a command line tool to generate gojay's marshaling and unmarshing interface implementation for struct, slice and map types. 

## Get started

```sh
go install github.com/francoispqt/gojay/gojay
```

## Generate code 

- for specific types in a go package, to stdout:
```sh
gojay -s github.com/francoispqt/gojay/gojay/tests -t A,B,StrSlice
```
or simply
```sh
gojay github.com/francoispqt/gojay/gojay/tests A,B,StrSlice
```

- for specific types in a go package, write to a file:
```sh
gojay -s github.com/francoispqt/gojay/gojay/tests -t A,B,StrSlice -o output.go
```

- for all types annotated by a //gojay:json (don't specify any type) in a package: 
```sh
gojay -s github.com/francoispqt/gojay/gojay/tests -o output.go
```

- for types annotated in a specific file
```sh
gojay -s path/to/gofile.go -o output.go
```

## Gojay tags

You can add tags to your structs to control:
- the JSON key
- skip a struct field only for unmarshaling
- skip a struct field only for marshaling
- skip a struct field

### Example: 
```go
type A struct {
    Str             string  `gojay:"string"`
    SkipUnmarshal   string  `gojay:"skipUnmarshal,-u"`
    SkipMarshal     string  `gojay:"skipMarshal,-m"`
    Skip            string  `gojay:"-"`
}
```