# Gojay code generator
This package provides a command line tool to generate gojay's marshaling and unmarshing interface implementation for custom struct type(s)


## Get started

```sh
go install github.com/francoispqt/gojay/gojaygen
```

## Generate code 

### Basic command
The basic command is straightforward and easy to use:
```sh
cd $GOPATH/src/github.com/user/project
gojaygen -s . -p true -t MyType -o output.go
```
If you just want to the output to stdout, omit the -o flag. 

### Using flags
- s file/dir path, can be a relative or absolute path
- t root types to generate with all its dependencies (comma separated)
- a annotation tag used to read meta data (default: json)
- o output file (relative or absolute path)
- p reuse object witt sync.Pool

Examples: 

- Specific type in a go package, write to a file:
```sh
gojay -s /tmp/myproj -t SomeType -o output.go
```

- Specific types in a file, with custom tag, write to stdout
```sh
gojay -s somegofile.go -a gojay -t SomeType
```


## Generator tags
You can add tags to your structs to control:

- the JSON key
- skip a struct field
- the use of omitempty methods for marshaling
- timeFormat (java style data format)
- timeLayout (golang time layout)
 

### Example: 
```go
type A struct {
    Str             string  `json:"string"`
    StrOmitEmpty    string  `json:"stringOrEmpty,omitempty"`
    Skip            string  `json:"-"`
	StartTime time.Time `json:"startDate" timeFormat:"yyyy-MM-dd HH:mm:ss"`
	EndTime *time.Time `json:"endDate" timeLayout:"2006-01-02 15:04:05"`
}
```

