package gojay

import "regexp"

// Schemas is a type representing a slice of schemas
type Schemas []*Schema

func (ss *Schemas) UnmarshalJSONArray(dec *Decoder) error {
	s := &Schema{}
	if err := dec.Object(s); err != nil {
		return err
	}
	*ss = append(*ss, s)
	return nil
}

// Schema is the type representing a JSON Schema
type Schema struct {
	ID          string
	Ref         string
	Schema      string
	Description string
	Const       *EmbeddedJSON
	Enum        *EmbeddedJSON
	// here we use EmbeddedJSON as it could be an object or an array
	// we will determine this later
	TypeJSON *EmbeddedJSON
	Type     string
	Types    *Strings
	Title    string
	// here we use EmbeddedJSON as it could be an object or an array
	// we will determine this later
	ItemsJSON    *EmbeddedJSON
	ItemsSchemas *Schemas
	ItemsSchema  *Schema
	OneOf        *Schemas
	Required     *Strings
	Properties   Properties
	Definitions  Definitions
	// string related
	MinLength int
	MaxLength int
	Pattern   *regexp.Regexp
	// number related
	MultipleOf int
	MinBytes   *EmbeddedJSON
	Min        *int
	ExcMin     bool
	MaxBytes   *EmbeddedJSON
	Max        *int
	ExcMax     bool
}

func (s *Schema) UnmarshalJSONObject(dec *Decoder, k string) error {
	switch k {
	case "id":
		return dec.String(&s.ID)
	case "title":
		return dec.String(&s.Title)
	case "$ref":
		return dec.String(&s.ID)
	case "$schema":
		return dec.String(&s.Schema)
	case "description":
		return dec.String(&s.Description)
	case "type":
		s.TypeJSON = &EmbeddedJSON{}
		return dec.AddEmbeddedJSON(s.TypeJSON)
	case "items":
		s.ItemsJSON = &EmbeddedJSON{}
		return dec.AddEmbeddedJSON(s.ItemsJSON)
	case "oneOf":
		schemas := make(Schemas, 0)
		s.OneOf = &schemas
		return dec.Array(s.OneOf)
	case "required":
		required := make(Strings, 0)
		s.Required = &required
		return dec.Array(s.Required)
	case "properties":
		s.Properties = make(Properties)
		return dec.Object(s.Properties)
	case "definitions":
		s.Definitions = make(Definitions)
		return dec.Object(s.Definitions)
	case "minLength":
		return dec.Int(&s.MinLength)
	case "maxLength":
		return dec.Int(&s.MaxLength)
	case "pattern":
		str := ""
		if err := dec.String(&str); err != nil {
			return err
		}
		r, err := regexp.Compile(str)
		if err != nil {
			return err
		}
		s.Pattern = r
		return nil
	case "multipleOf":
		return dec.Int(&s.MultipleOf)
	case "minimum":
		s.MinBytes = &EmbeddedJSON{}
		return dec.AddEmbeddedJSON(s.MinBytes)
	case "exclusiveMinimum":
		return dec.Bool(&s.ExcMin)
	case "maximum":
		s.MaxBytes = &EmbeddedJSON{}
		return dec.AddEmbeddedJSON(s.MaxBytes)
	case "exclusiveMaximum":
		return dec.Bool(&s.ExcMax)
	}
	return nil
}

func (s *Schema) NKeys() int {
	return 0
}

// Properties
type Properties map[string]*Schema

func (p Properties) UnmarshalJSONObject(dec *Decoder, k string) error {
	s := &Schema{}
	if err := dec.Object(s); err != nil {
		return err
	}
	p[k] = s
	return nil
}

func (p Properties) NKeys() int {
	return 0
}

// Definitions
type Definitions map[string]*Schema

func (p Definitions) UnmarshalJSONObject(dec *Decoder, k string) error {
	s := &Schema{}
	if err := dec.Object(s); err != nil {
		return err
	}
	p[k] = s
	return nil
}

func (p Definitions) NKeys() int {
	return 0
}
