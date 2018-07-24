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
	}{
		{
			name:               "basic",
			json:               `"test"`,
			expectedNullString: sql.NullString{String: "test", Valid: true},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			nullString := sql.NullString{}
			dec := NewDecoder(strings.NewReader(testCase.json))
			dec.DecodeSQLNullString(&nullString)
			assert.Equal(t, testCase.expectedNullString, nullString)
		})
	}
}

func TestDecodeSQLNullInt64(t *testing.T) {
	testCases := []struct {
		name              string
		json              string
		expectedNullInt64 sql.NullInt64
	}{
		{
			name:              "basic",
			json:              `1`,
			expectedNullInt64: sql.NullInt64{Int64: 1, Valid: true},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			nullInt64 := sql.NullInt64{}
			dec := NewDecoder(strings.NewReader(testCase.json))
			dec.DecodeSQLNullInt64(&nullInt64)
			assert.Equal(t, testCase.expectedNullInt64, nullInt64)
		})
	}
}
