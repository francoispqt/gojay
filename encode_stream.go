package gojay

import "strconv"

// MarshalerStream is the interface to implement
// to continuously encode of stream of data.
type MarshalerStream interface {
	MarshalStream(enc *StreamEncoder) error
}

// A StreamEncoder reads and encodes values to JSON from an input stream.
//
// It implements conext.Context and provide a channel to notify interruption.
type StreamEncoder struct {
	*Encoder
	err       error
	nConsumer int
	delimiter byte
	done      chan struct{}
}

// EncodeStream spins up a defined number of non blocking consumers of the MarshalerStream m.
//
// m must implement MarshalerStream. Ideally m is a channel. See example for implementation.
//
// See the documentation for Marshal for details about the conversion of Go value to JSON.
func (s *StreamEncoder) EncodeStream(m MarshalerStream) {
	// if a single consumer, just use this encoder
	if s.nConsumer == 1 {
		go consume(s, s, m)
		return
	}
	// else use this Encoder only for first consumer
	// and use new encoders for other consumers
	// this is to avoid concurrent writing to same buffer
	// resulting in a weird JSON
	go consume(s, s, m)
	for i := 1; i < s.nConsumer; i++ {
		ss := Stream.borrowEncoder(s.w)
		ss.done = s.done
		ss.buf = make([]byte, 0, 512)
		ss.delimiter = s.delimiter
		go consume(s, ss, m)
	}
	return
}

// LineDelimited sets the delimiter to a new line character.
//
// It will add a new line after each JSON marshaled by the MarshalerStream
func (s *StreamEncoder) LineDelimited() *StreamEncoder {
	s.delimiter = '\n'
	return s
}

// CommaDelimited sets the delimiter to a comma.
//
// It will add a new line after each JSON marshaled by the MarshalerStream
func (s *StreamEncoder) CommaDelimited() *StreamEncoder {
	s.delimiter = ','
	return s
}

// NConsumer sets the number of non blocking go routine to consume the stream.
func (s *StreamEncoder) NConsumer(n int) *StreamEncoder {
	s.nConsumer = n
	return s
}

// Release sends back a Decoder to the pool.
// If a decoder is used after calling Release
// a panic will be raised with an InvalidUsagePooledDecoderError error.
func (s *StreamEncoder) Release() {
	s.Encoder.isPooled = 1
	select {
	case streamEncPool <- s:
	default:
	}
}

// Done returns a channel that's closed when work is done.
// It implements context.Context
func (s *StreamEncoder) Done() <-chan struct{} {
	return s.done
}

// Err returns nil if Done is not yet closed.
// If Done is closed, Err returns a non-nil error explaining why.
// It implements context.Context
func (s *StreamEncoder) Err() error {
	return s.err
}

// Cancel cancels the consumers of the stream, interrupting the stream encoding.
//
// After calling cancel, Done() will return a closed channel.
func (s *StreamEncoder) Cancel(err error) {
	select {
	case <-s.done:
	default:
		s.err = err
		close(s.done)
	}
}

// AddObject adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerObject
func (s *StreamEncoder) AddObject(v MarshalerObject) error {
	if v.IsNil() {
		return nil
	}
	s.Encoder.writeByte('{')
	v.MarshalObject(s.Encoder)
	s.Encoder.writeByte('}')
	s.Encoder.writeByte(s.delimiter)
	return nil
}

// AddInt adds an int to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (s *StreamEncoder) AddInt(value int) error {
	s.buf = strconv.AppendInt(s.buf, int64(value), 10)
	s.Encoder.writeByte(s.delimiter)
	return nil
}

// AddFloat adds a float64 to be encoded, must be used inside a slice or array encoding (does not encode a key)
func (s *StreamEncoder) AddFloat(value float64) error {
	s.buf = strconv.AppendFloat(s.buf, value, 'f', -1, 64)
	s.Encoder.writeByte(s.delimiter)
	return nil
}

// Non exposed

func consume(init *StreamEncoder, s *StreamEncoder, m MarshalerStream) {
	defer s.Release()
	for {
		select {
		case <-init.Done():
			return
		default:
			err := m.MarshalStream(s)
			if err != nil {
				init.Cancel(err)
				return
			}
			i, err := s.Encoder.write()
			if err != nil || i == 0 {
				init.Cancel(err)
				return
			}
		}
	}
}
