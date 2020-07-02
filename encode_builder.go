package gojay

const hex = "0123456789abcdef"

// grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, grow panics.
func (enc *Encoder) grow(n int) {
	if cap(enc.buf)-len(enc.buf) < n {
		Buf := make([]byte, len(enc.buf), 2*cap(enc.buf)+n)
		copy(Buf, enc.buf)
		enc.buf = Buf
	}
}

// Write appends the contents of p to b's Buffer.
// Write always returns len(p), nil.
func (enc *Encoder) writeBytes(p []byte) {
	enc.buf = append(enc.buf, p...)
}

func (enc *Encoder) writeTwoBytes(b1 byte, b2 byte) {
	enc.buf = append(enc.buf, b1, b2)
}

// WriteByte appends the byte c to b's Buffer.
// The returned error is always nil.
func (enc *Encoder) writeByte(c byte) {
	enc.buf = append(enc.buf, c)
}

// WriteString appends the contents of s to b's Buffer.
// It returns the length of s and a nil error.
func (enc *Encoder) writeString(s string) {
	enc.buf = append(enc.buf, s...)
}

func (enc *Encoder) writeStringEscape(s string) {
	for i := 0; i < len(s); i++ {
		if isABlackListedControl(s[i]) {
			continue
		}

		c := s[i]
		if c >= 0x20 && c != '\\' && c != '"' {
			enc.writeByte(c)
			continue
		}

		switch c {
		case '\\', '"':
			enc.writeTwoBytes('\\', c)
		case '\n':
			enc.writeTwoBytes('\\', 'n')
		case '\f':
			enc.writeTwoBytes('\\', 'f')
		case '\b':
			enc.writeTwoBytes('\\', 'b')
		case '\r':
			enc.writeTwoBytes('\\', 'r')
		case '\t':
			enc.writeTwoBytes('\\', 't')
		default:
			enc.writeString(`\u00`)
			enc.writeTwoBytes(hex[c>>4], hex[c&0xF])
		}
	}
}

// not all controls should be considered for JSON string representations
func isABlackListedControl(r uint8) bool {
	// fast case all the known controls are below ascii code 20
	if r >= 0x20 {
		return false
	}

	// ignoring Null char, Start of Heading, Start of Text, End of Text, End of Transmission, Enquiry, Acknowledgment, Bell controls
	if r == 0x00 || r <= 0x07 {
		return true
	}
	// ignoring the vertical tab, it's a legacy field and used in printers, has no purpose in JSON
	if r == 0x0B {
		return true
	}

	//ignoring Shift Out / X-On, Shift In / X-Off, Data Line Escape, Device Control 1 (oft. XON), Device Control 2, Device Control 3 (oft. XOFF),
	//Device Control 4, Negative Acknowledgement, Synchronous Idle, End of Transmit Block, Cancel, End of Medium, Substitute,
	//File Separator, Group Separator, Record Separator, Unit Separator controls
	//The escape control was left as a result of it being used in other tests within the application, however it should be noted
	//that the escape control is known to be problematic in JSON
	//https://www.bennadel.com/blog/2576-testing-which-ascii-characters-break-json-javascript-object-notation-parsing.htm
	if r != 0x1B && r >= 0x0E && r <= 0x1F {
		return true
	}

	return false
}
