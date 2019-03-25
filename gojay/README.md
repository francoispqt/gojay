# Gojay code generator
This package provides a command line tool to generate gojay's marshaling and unmarshaling interface implementation for custom struct type(s)


## Get started

```sh
go install github.com/francoispqt/gojay/gojay
```

## Generate code

### Basic command
The basic command is straightforward and easy to use:
```sh
cd $GOPATH/src/github.com/user/project
gojay -s . -p true -t MyType -o output.go
```
If you just want to the output to stdout, omit the -o flag.

### Using flags
- s Source file/dir path, can be a relative or absolute path
- t Types to generate with all its dependencies (comma separated)
- a Annotation tag used to read metadata (default: json)
- o Output file (relative or absolute path)
- p Pool to reuse object (using sync.Pool)

Examples:

- Generate `SomeType` type in `/tmp/myproj` go package, write to file `output.go`:
```sh
gojay -s /tmp/myproj -t SomeType -o output.go
```

- Generate type `SomeType` in file `somegofile.go`, with custom tag `gojay`, write to stdout:
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
	Str          string     `json:"string"`
	StrOmitEmpty string     `json:"stringOrEmpty,omitempty"`
	Skip         string     `json:"-"`
	StartTime    time.Time  `json:"startDate" timeFormat:"yyyy-MM-dd HH:mm:ss"`
	EndTime      *time.Time `json:"endDate" timeLayout:"2006-01-02 15:04:05"`
}
```

