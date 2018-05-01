package gojay

import "errors"

type TestWriterError string

func (t TestWriterError) Write(b []byte) (int, error) {
	return 0, errors.New("Test Error")
}
