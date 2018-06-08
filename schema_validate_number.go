package gojay

import (
	"errors"
	"fmt"
)

func (s *Schema) validateNumberType(v interface{}) error {
	switch t := v.(type) {
	case int:
		return s.validateNumberInt(t)
	case int8:
		return s.validateNumberInt(int(t))
	case int16:
		return s.validateNumberInt(int(t))
	case int32:
		return s.validateNumberInt(int(t))
	case int64:
		return s.validateNumberInt(int(t))
	case uint8:
		return s.validateNumberUint(uint64(t))
	case uint16:
		return s.validateNumberUint(uint64(t))
	case uint32:
		return s.validateNumberUint(uint64(t))
	case uint64:
		return s.validateNumberUint(t)
	case float32:
		return s.validateNumberFloat(float64(t))
	case float64:
		return s.validateNumberFloat(t)
	case *int:
		return s.validateNumberInt(*t)
	case *int8:
		return s.validateNumberInt(int(*t))
	case *int16:
		return s.validateNumberInt(int(*t))
	case *int32:
		return s.validateNumberInt(int(*t))
	case *int64:
		return s.validateNumberInt(int(*t))
	case *uint8:
		return s.validateNumberUint(uint64(*t))
	case *uint16:
		return s.validateNumberUint(uint64(*t))
	case *uint32:
		return s.validateNumberUint(uint64(*t))
	case *uint64:
		return s.validateNumberUint(uint64(*t))
	case *float32:
		return s.validateNumberFloat(float64(*t))
	case *float64:
		return s.validateNumberFloat(*t)
	}
	return errors.New("Invalid type")
}

func (s *Schema) validateIntegerType(v interface{}) error {
	switch t := v.(type) {
	case int:
		return s.validateNumberInt(t)
	case int8:
		return s.validateNumberInt(int(t))
	case int16:
		return s.validateNumberInt(int(t))
	case int32:
		return s.validateNumberInt(int(t))
	case int64:
		return s.validateNumberInt(int(t))
	case uint8:
		return s.validateNumberUint(uint64(t))
	case uint16:
		return s.validateNumberUint(uint64(t))
	case uint32:
		return s.validateNumberUint(uint64(t))
	case uint64:
		return s.validateNumberUint(t)
	case *int:
		return s.validateNumberInt(*t)
	case *int8:
		return s.validateNumberInt(int(*t))
	case *int16:
		return s.validateNumberInt(int(*t))
	case *int32:
		return s.validateNumberInt(int(*t))
	case *int64:
		return s.validateNumberInt(int(*t))
	case *uint8:
		return s.validateNumberUint(uint64(*t))
	case *uint16:
		return s.validateNumberUint(uint64(*t))
	case *uint32:
		return s.validateNumberUint(uint64(*t))
	case *uint64:
		return s.validateNumberUint(uint64(*t))
	}
	return errors.New("Invalid type")
}

func (s *Schema) validateNumberInt(n int) error {
	if s.MultipleOf != 0 {
		if n%s.MultipleOf != 0 {
			return SchemaValidationError(fmt.Sprintf("Number is not multiple of %d", s.MultipleOf))
		}
	}
	if s.Min != nil {
		if s.ExcMin {
			if n <= *s.Min {
				return SchemaValidationError(fmt.Sprintf("Number is lower than %d", *s.Min))
			}
		} else {
			if n < *s.Min {
				return SchemaValidationError(fmt.Sprintf("Number is lower than %d", *s.Min))
			}
		}
	}
	if s.Max != nil {
		if s.ExcMax {
			if n >= *s.Max {
				return SchemaValidationError(fmt.Sprintf("Number is greater than %d", *s.Max))
			}
		} else {
			if n > *s.Max {
				return SchemaValidationError(fmt.Sprintf("Number is greater than %d", *s.Max))
			}
		}
	}
	return nil
}

func (s *Schema) validateNumberUint(n uint64) error {
	if s.MultipleOf != 0 {
		if int(n)%s.MultipleOf != 0 {
			return SchemaValidationError(fmt.Sprintf("Number is not multiple of %d", s.MultipleOf))
		}
	}
	if s.Min != nil {
		if s.ExcMin {
			if n < uint64(*s.Min) {
				return SchemaValidationError(fmt.Sprintf("Number is lower than %d", *s.Min))
			}
		} else {
			if n <= uint64(*s.Min) {
				return SchemaValidationError(fmt.Sprintf("Number is lower than %d", *s.Min))
			}
		}
	}
	if s.Max != nil {
		if s.ExcMax {
			if n > uint64(*s.Max) {
				return SchemaValidationError(fmt.Sprintf("Number is greater than %d", *s.Min))
			}
		} else {
			if n >= uint64(*s.Max) {
				return SchemaValidationError(fmt.Sprintf("Number is greater than %d", *s.Min))
			}
		}
	}
	return nil
}

func (s *Schema) validateNumberFloat(n float64) error {
	if s.MultipleOf != 0 {
		if int(n)%s.MultipleOf != 0 {
			return SchemaValidationError(fmt.Sprintf("Number is not multiple of %d", s.MultipleOf))
		}
	}
	if s.Min != nil {
		if s.ExcMin {
			if n < float64(*s.Min) {
				return SchemaValidationError(fmt.Sprintf("Number is lower than %d", *s.Min))
			}
		} else {
			if n <= float64(*s.Min) {
				return SchemaValidationError(fmt.Sprintf("Number is lower than %d", *s.Min))
			}
		}
	}
	if s.Max != nil {
		if s.ExcMax {
			if n > float64(*s.Max) {
				return SchemaValidationError(fmt.Sprintf("Number is greater than %d", *s.Min))
			}
		} else {
			if n >= float64(*s.Max) {
				return SchemaValidationError(fmt.Sprintf("Number is greater than %d", *s.Min))
			}
		}
	}
	return nil
}
