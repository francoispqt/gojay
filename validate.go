package gojay

type ObjectValidator interface {
	ValidateJSONObject(*Schema)
}
type ArrayValidator interface {
	ValidateJSONArray(*Schema)
}

func (s *Schema) Validate(v interface{}) error {
	// check if const
	if s.Const != nil {
		return s.ValidateConst(v)
	}
	// check if enum
	if s.Enum != nil {
		return s.ValidateEnum(v)
	}
	// then we validate type
	err := s.ValidateType(v)
	if err != nil {
		return err
	}
	return nil
}
