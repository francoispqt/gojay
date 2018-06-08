package gojay

func isArray(b []byte) bool {
	bLen := len(b)
	for i := 0; i < bLen; i++ {
		switch b[i] {
		case ' ', '\t', '\r', '\n':
			continue
		case '[':
			return true
		default:
			return false
		}
	}
	return false
}
