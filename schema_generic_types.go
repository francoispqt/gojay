package gojay

type Strings []string

func (s *Strings) UnmarshalJSONArray(dec *Decoder) error {
	str := ""
	if err := dec.String(&str); err != nil {
		return err
	}
	*s = append(*s, str)
	return nil
}
