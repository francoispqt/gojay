package gojay

import "fmt"

func (s *Schema) validateStringType(v interface{}) error {
	switch t := v.(type) {
	case string:
		return s.validateString(t)
	case *string: // is correct type
		return s.validateString(*t)
	}
	return SchemaValidationError("Invalid type provided")
}

func (s *Schema) validateString(str string) error {
	if s.MaxLength != 0 {
		if len(str) > s.MaxLength {
			return SchemaValidationError("String length is higher than maxLength")
		}
	}
	if s.MinLength != 0 {
		if len(str) < s.MinLength {
			return SchemaValidationError("String length is higher than maxLength")
		}
	}
	if s.Pattern != nil {
		if !s.Pattern.MatchString(str) {
			return SchemaValidationError(fmt.Sprintf("String does not match pattern %s", s.Pattern))
		}
	}
	return nil
}
