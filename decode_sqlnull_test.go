package gojay

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeSQLNullString(t *testing.T) {
	testCases := []struct {
		name               string
		json               string
		expectedNullString sql.NullString
		err                bool
	}{
		{
			name:               "basic",
			json:               `"test"`,
			expectedNullString: sql.NullString{String: "test", Valid: true},
		},
		{
			name:               "basic",
			json:               `"test`,
			expectedNullString: sql.NullString{String: "test", Valid: true},
			err:                true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			nullString := sql.NullString{}
			dec := NewDecoder(strings.NewReader(testCase.json))
			err := dec.DecodeSQLNullString(&nullString)
			if testCase.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedNullString, nullString)
			}
		})
	}
	t.Run(
		"should panic because decoder is pooled",
		func(t *testing.T) {
			dec := NewDecoder(nil)
			dec.Release()
			defer func() {
				err := recover()
				assert.NotNil(t, err, "err shouldnt be nil")
				assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
			}()
			_ = dec.DecodeSQLNullString(&sql.NullString{})
			assert.True(t, false, "should not be called as decoder should have panicked")
		},
	)
}

func TestDecodeSQLNullInt64(t *testing.T) {
	testCases := []struct {
		name              string
		json              string
		expectedNullInt64 sql.NullInt64
		err               bool
	}{
		{
			name:              "basic",
			json:              `1`,
			expectedNullInt64: sql.NullInt64{Int64: 1, Valid: true},
		},
		{
			name:              "basic",
			json:              `"test`,
			expectedNullInt64: sql.NullInt64{Int64: 1, Valid: true},
			err:               true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			nullInt64 := sql.NullInt64{}
			dec := NewDecoder(strings.NewReader(testCase.json))
			err := dec.DecodeSQLNullInt64(&nullInt64)
			if testCase.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedNullInt64, nullInt64)
			}
		})
	}
	t.Run(
		"should panic because decoder is pooled",
		func(t *testing.T) {
			dec := NewDecoder(nil)
			dec.Release()
			defer func() {
				err := recover()
				assert.NotNil(t, err, "err shouldnt be nil")
				assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
			}()
			_ = dec.DecodeSQLNullInt64(&sql.NullInt64{})
			assert.True(t, false, "should not be called as decoder should have panicked")
		},
	)
}

func TestDecodeSQLNullFloat64(t *testing.T) {
	testCases := []struct {
		name                string
		json                string
		expectedNullFloat64 sql.NullFloat64
		err                 bool
	}{
		{
			name:                "basic",
			json:                `1`,
			expectedNullFloat64: sql.NullFloat64{Float64: 1, Valid: true},
		},
		{
			name:                "basic",
			json:                `"test`,
			expectedNullFloat64: sql.NullFloat64{Float64: 1, Valid: true},
			err:                 true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			nullFloat64 := sql.NullFloat64{}
			dec := NewDecoder(strings.NewReader(testCase.json))
			err := dec.DecodeSQLNullFloat64(&nullFloat64)
			if testCase.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedNullFloat64, nullFloat64)
			}
		})
	}
	t.Run(
		"should panic because decoder is pooled",
		func(t *testing.T) {
			dec := NewDecoder(nil)
			dec.Release()
			defer func() {
				err := recover()
				assert.NotNil(t, err, "err shouldnt be nil")
				assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
			}()
			_ = dec.DecodeSQLNullFloat64(&sql.NullFloat64{})
			assert.True(t, false, "should not be called as decoder should have panicked")
		},
	)
}

func TestDecodeSQLNullBool(t *testing.T) {
	testCases := []struct {
		name             string
		json             string
		expectedNullBool sql.NullBool
		err              bool
	}{
		{
			name:             "basic",
			json:             `true`,
			expectedNullBool: sql.NullBool{Bool: true, Valid: true},
		},
		{
			name:             "basic",
			json:             `"&`,
			expectedNullBool: sql.NullBool{Bool: true, Valid: true},
			err:              true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			nullBool := sql.NullBool{}
			dec := NewDecoder(strings.NewReader(testCase.json))
			err := dec.DecodeSQLNullBool(&nullBool)
			if testCase.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedNullBool, nullBool)
			}
		})
	}
	t.Run(
		"should panic because decoder is pooled",
		func(t *testing.T) {
			dec := NewDecoder(nil)
			dec.Release()
			defer func() {
				err := recover()
				assert.NotNil(t, err, "err shouldnt be nil")
				assert.IsType(t, InvalidUsagePooledDecoderError(""), err, "err should be of type InvalidUsagePooledDecoderError")
			}()
			_ = dec.DecodeSQLNullBool(&sql.NullBool{})
			assert.True(t, false, "should not be called as decoder should have panicked")
		},
	)
}
