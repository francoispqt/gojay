// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gojay

// grow grows b's capacity, if necessary, to guarantee space for
// another n bytes. After grow(n), at least n bytes can be written to b
// without another allocation. If n is negative, grow panics.
func (enc *Encoder) grow(n int) {
	if n < 0 {
		panic("Builder.grow: negative count")
	}
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
