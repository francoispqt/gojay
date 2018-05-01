package gojay

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StreamChanObject chan *testObject

func (s StreamChanObject) MarshalStream(enc *StreamEncoder) error {
	select {
	case <-enc.Done():
		return enc.Err()
	case o := <-s:
		return enc.AddObject(o)
	}
}

type StreamChanInt chan int

func (s StreamChanInt) MarshalStream(enc *StreamEncoder) error {
	select {
	case <-enc.Done():
		return enc.Err()
	case o := <-s:
		return enc.AddInt(o)
	}
}

type StreamChanFloat chan float64

func (s StreamChanFloat) MarshalStream(enc *StreamEncoder) error {
	select {
	case <-enc.Done():
		return enc.Err()
	case o := <-s:
		return enc.AddFloat(o)
	}
}

type StreamChanError chan *testObject

func (s StreamChanError) MarshalStream(enc *StreamEncoder) error {
	select {
	case <-enc.Done():
		return enc.Err()
	case <-s:
		return errors.New("Test Error")
	}
}

func TestEncodeStreamSingleConsumer(t *testing.T) {
	expectedStr :=
		`{"testStr":"","testInt":0,"testInt64":0,"testInt32":0,"testInt16":0,"testInt8":0,"testUint64":0,"testUint32":0,"testUint16":0,"testUint8":0,"testFloat64":0,"testFloat32":0,"testBool":false}
`
	// create our writer
	w := &TestWriter{target: 100, mux: &sync.RWMutex{}}
	enc := Stream.NewEncoder(w).LineDelimited()
	w.enc = enc
	s := StreamChanObject(make(chan *testObject))
	go enc.EncodeStream(s)
	go feedStream(s, 100)
	select {
	case <-enc.Done():
		assert.Len(t, w.result, 100, "w.result should be 100")
		for _, b := range w.result {
			assert.Equal(t, expectedStr, string(b), "every byte buffer should be equal to expected string")
		}
	}
}
func TestEncodeStreamSingleConsumerInt(t *testing.T) {
	// create our writer
	w := &TestWriter{target: 100, mux: &sync.RWMutex{}}
	enc := Stream.NewEncoder(w).LineDelimited()
	w.enc = enc
	s := StreamChanInt(make(chan int))
	go enc.EncodeStream(s)
	go feedStreamInt(s, 100)
	select {
	case <-enc.Done():
		assert.Len(t, w.result, 100, "w.result should be 100")
	}
}
func TestEncodeStreamSingleConsumerFloat(t *testing.T) {
	// create our writer
	w := &TestWriter{target: 100, mux: &sync.RWMutex{}}
	enc := Stream.NewEncoder(w).LineDelimited()
	w.enc = enc
	s := StreamChanFloat(make(chan float64))
	go enc.EncodeStream(s)
	go feedStreamFloat(s, 100)
	select {
	case <-enc.Done():
		assert.Len(t, w.result, 100, "w.result should be 100")
	}
}
func TestEncodeStreamSingleConsumerMarshalError(t *testing.T) {
	// create our writer
	w := &TestWriter{target: 100, mux: &sync.RWMutex{}}
	enc := Stream.NewEncoder(w).LineDelimited()
	w.enc = enc
	s := StreamChanError(make(chan *testObject))
	go enc.EncodeStream(s)
	go feedStream(s, 100)
	select {
	case <-enc.Done():
		assert.NotNil(t, enc.Err(), "enc.Err() should not be nil")
	}
}

func TestEncodeStreamSingleConsumerWriteError(t *testing.T) {
	// create our writer
	w := TestWriterError("")
	enc := Stream.NewEncoder(w).LineDelimited()
	s := StreamChanObject(make(chan *testObject))
	go enc.EncodeStream(s)
	go feedStream(s, 100)
	select {
	case <-enc.Done():
		assert.NotNil(t, enc.Err(), "enc.Err() should not be nil")
	}
}
func TestEncodeStreamSingleConsumerCommaDelimited(t *testing.T) {
	expectedStr :=
		`{"testStr":"","testInt":0,"testInt64":0,"testInt32":0,"testInt16":0,"testInt8":0,"testUint64":0,"testUint32":0,"testUint16":0,"testUint8":0,"testFloat64":0,"testFloat32":0,"testBool":false},`
	// create our writer
	w := &TestWriter{target: 5000, mux: &sync.RWMutex{}}
	enc := Stream.BorrowEncoder(w).NConsumer(50).CommaDelimited()
	w.enc = enc
	s := StreamChanObject(make(chan *testObject))
	go enc.EncodeStream(s)
	go feedStream(s, 5000)
	select {
	case <-enc.Done():
		assert.Len(t, w.result, 5000, "w.result should be 100")
		for _, b := range w.result {
			assert.Equal(t, expectedStr, string(b), "every byte buffer should be equal to expected string")
		}
	}
}

func TestEncodeStreamMultipleConsumer(t *testing.T) {
	expectedStr :=
		`{"testStr":"","testInt":0,"testInt64":0,"testInt32":0,"testInt16":0,"testInt8":0,"testUint64":0,"testUint32":0,"testUint16":0,"testUint8":0,"testFloat64":0,"testFloat32":0,"testBool":false}
`
	// create our writer
	w := &TestWriter{target: 5000, mux: &sync.RWMutex{}}
	enc := Stream.NewEncoder(w).NConsumer(50).LineDelimited()
	w.enc = enc
	s := StreamChanObject(make(chan *testObject))
	go enc.EncodeStream(s)
	go feedStream(s, 5000)
	select {
	case <-enc.Done():
		assert.Len(t, w.result, 5000, "w.result should be 100")
		for _, b := range w.result {
			assert.Equal(t, expectedStr, string(b), "every byte buffer should be equal to expected string")
		}
	}
}

// TestWriter to assert result
type TestWriter struct {
	nWrite *int
	target int
	enc    *StreamEncoder
	result [][]byte
	mux    *sync.RWMutex
}

func (w *TestWriter) Write(b []byte) (int, error) {
	if len(b) > 0 {
		w.mux.Lock()
		w.result = append(w.result, b)
		if len(w.result) == w.target {
			w.enc.Cancel(nil)
		}
		w.mux.Unlock()
	}
	return len(b), nil
}

func feedStream(s chan *testObject, target int) {
	for i := 0; i < target; i++ {
		s <- &testObject{}
	}
}

func feedStreamInt(s chan int, target int) {
	for i := 0; i < target; i++ {
		s <- i
	}
}

func feedStreamFloat(s chan float64, target int) {
	for i := 0; i < target; i++ {
		s <- float64(i)
	}
}
