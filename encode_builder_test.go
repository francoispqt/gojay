package gojay

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CountingWritesBuffer struct {
	bytes.Buffer
	WriteCalls int
}

func (c *CountingWritesBuffer) Write(p []byte) (int, error) {
	c.WriteCalls++
	return c.Buffer.Write(p)
}

type PayloadForEncode struct {
	St   int
	RS   io.Reader
	Sid  int
	Tt   string
	Gr   int
	Uuid string
	Ip   string
	Ua   string
	Tz   int
	R64  io.Reader
	V    int
}

func (p *PayloadForEncode) MarshalJSONObject(enc *Encoder) {
	enc.AddIntKey("st", p.St)
	enc.AddReaderToEscapedKey("rs", p.RS)
	enc.AddIntKey("sid", p.Sid)
	enc.AddStringKey("tt", p.Tt)
	enc.AddIntKey("gr", p.Gr)
	enc.AddStringKey("uuid", p.Uuid)
	enc.AddStringKey("ip", p.Ip)
	enc.AddStringKey("ua", p.Ua)
	enc.AddIntKey("tz", p.Tz)
	enc.AddReaderToBase64Key("r64", p.R64, base64.StdEncoding)
	enc.AddIntKey("v", p.V)
}

func (p *PayloadForEncode) IsNil() bool { return p == nil }

func TestEncodeWithFlush(t *testing.T) {
	t.Run("buffer must be flushed after threshold reached", func(t *testing.T) {
		var target CountingWritesBuffer

		var randBytes [120]byte
		_, err := io.ReadFull(rand.Reader, randBytes[:])
		assert.NoError(t, err)

		encoder := BorrowEncoder(&target)
		defer encoder.Release()

		const bufferFlushThreshold = 64
		encoder.SetBufFlushThreshold(bufferFlushThreshold)
		assert.NoError(t, encoder.EncodeObject(&PayloadForEncode{
			St:   1,
			RS:   bytes.NewReader(randBytes[:]),
			Sid:  2,
			Tt:   "TestString",
			Gr:   4,
			Uuid: "8f9a65eb-4807-4d57-b6e0-bda5d62f1429",
			Ip:   "127.0.0.1",
			Ua:   "Mozilla",
			Tz:   8,
			R64:  bytes.NewReader([]byte{1, 2, 3, 4}),
			V:    6,
		}))

		wroteBytes := len(target.Bytes())
		expectedWriteCalls := 1 + (wroteBytes-1)/bufferFlushThreshold
		assert.Equal(t, expectedWriteCalls, target.WriteCalls)
	})

	t.Run("ensure that output is valid", func(t *testing.T) {
		var target bytes.Buffer

		encoder := BorrowEncoder(&target)
		defer encoder.Release()

		const bufferFlushThreshold = 64
		encoder.SetBufFlushThreshold(bufferFlushThreshold)
		assert.NoError(t, encoder.EncodeObject(&PayloadForEncode{
			St:   1,
			RS:   strings.NewReader(`wkofowk[grlmgaemriogjjgivsinfvna/snbgaipw43jgh'jnsprnbigphrjizsjo;ijb;osdjtnbs'`),
			Sid:  2,
			Tt:   "TestString",
			Gr:   4,
			Uuid: "8f9a65eb-4807-4d57-b6e0-bda5d62f1429",
			Ip:   "127.0.0.1",
			Ua:   "Mozilla",
			Tz:   8,
			R64:  strings.NewReader(`aoksdfpos'agpmejriojgp'nirbnatngads lkmsalkemflsapkdfpoakdospkf`),
			V:    6,
		}))

		assert.JSONEq(t, `
{
	"st": 1,
	"rs": "wkofowk[grlmgaemriogjjgivsinfvna/snbgaipw43jgh'jnsprnbigphrjizsjo;ijb;osdjtnbs'",
	"sid": 2,
	"tt": "TestString",
	"gr": 4,
	"uuid": "8f9a65eb-4807-4d57-b6e0-bda5d62f1429",
	"ip": "127.0.0.1",
	"ua": "Mozilla",
	"tz": 8,
	"r64": "YW9rc2RmcG9zJ2FncG1lanJpb2pncCduaXJibmF0bmdhZHMgbGttc2Fsa2VtZmxzYXBrZGZwb2FrZG9zcGtm",
	"v": 6
}
`, target.String())
	})
}
