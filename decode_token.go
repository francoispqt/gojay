package gojay

type Token uint8

const (
	TokenUnknown Token = iota
	TokenArray
	TokenBoolean
	TokenNull
	TokenNumber
	TokenObject
	TokenString
)

// NextToken returns a token identifying whats next in the underline data buffer as a token e.g string (") array ([)
// should only be used in the context of satisfying the UnmarshalJSONObject interface
func (dec *Decoder) NextToken() Token {
	switch b := dec.nextChar(); {
	case b == '[':
		return TokenArray
	case b == 't', b == 'f':
		return TokenBoolean
	case b == 'n':
		return TokenNull
	case rune(b) >= 0x30 && rune(b) <= 0x39:
		return TokenNumber
	case b == '{':
		return TokenObject
	case b == '"':
		return TokenString
	default:
		return TokenUnknown
	}
}
