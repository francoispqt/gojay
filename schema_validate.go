package gojay

import (
	"errors"
	"log"
)

func (s *Schema) ValidatePathString(path, v string) error {
	if path == "." {
		return s.Validate(v)
	}
	return nil
}

func (s *Schema) ValidatePath(path, v interface{}) error {
	if path == "." {
		log.Print("validating")
		return s.Validate(v)
	}
	return nil
}

// ValidateConst validates a json schema const against the value v given
func (s *Schema) ValidateConst(v interface{}) error {
	return nil
}

// ValidateConst validates a json schema const against the value v given
func (s *Schema) ValidateEnum(v interface{}) error {
	return nil
}

// ValidateType validates a json schema type against the value v given
func (s *Schema) ValidateType(v interface{}) error {
	if s.Types != nil {
		types := *s.Types
		var err error
		for _, t := range types {
			// if err is nil type is valid, so we break
			if err := s.validateType(t, v); err == nil {
				break
			}
		}
		return err
	}
	t := s.Type
	return s.validateType(t, v)
}

func (s *Schema) validateType(t string, v interface{}) error {
	switch t {
	case "string":
		return s.validateStringType(v)
	case "number":
		return s.validateNumberType(v)
	case "integer":
		return s.validateIntegerType(v)
	case "object":
		return validateObject(v)
	case "array":
		return validateArray(v)
	case "boolean":
		return validateBool(v)
	case "null":
		return validateNull(v)
	}
	return errors.New("Invalid type")
}

func validateObject(v interface{}) error {
	if _, ok := v.(ObjectValidator); ok {
		return nil
	}
	return errors.New("Invalid type")
}

func validateArray(v interface{}) error {
	if _, ok := v.(ArrayValidator); ok {
		return nil
	}
	return errors.New("Invalid type")
}

func validateBool(v interface{}) error {
	switch v.(type) {
	case bool:
		return nil
	}
	return errors.New("Invalid type")
}

func validateNull(v interface{}) error {
	switch v.(type) {
	case nil:
		return nil
	}
	return errors.New("Invalid type")
}
