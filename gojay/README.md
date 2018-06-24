# Gojay code generator

This package provides a command line tool to generate gojay's marshaling and unmarshing interface implementation for struct, slice and map types. 

## Get started

```sh
go install github.com/francoispqt/gojay/gojay
```

## Generate code 

### Basic command

The basic command is strait forward and easy to use:
```sh
gojay github.com/some/package TypeA,TypeB,TypeC output.go 
```
If you just want to the output to stdout, omit the third parameter. 

### Using flags

- p package to parse, relative path to $GOPATH/src
- s file/dir to path, can be a relative or absolute path
- t types to generate (comma separated)
- o output file (relative or absolute path)

Examples: 
- Specific types in a go package, to stdout:
```sh
gojay -p github.com/francoispqt/gojay/gojay/tests -t A,B,StrSlice 
```

- Specific types in a go package, write to a file:
```sh
gojay -p github.com/francoispqt/gojay/gojay/tests -t A,B,StrSlice -o output.go
```

- Specific types in a go file, to stdout: 
```sh
gojay -s somegofile.go -t SomeType
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