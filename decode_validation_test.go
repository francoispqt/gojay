package gojay

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type validationTestObj struct {
	Test string
}

func (v *validationTestObj) UnmarshalJSONObject(dec *Decoder, k string) error {
	switch k {
	case "test":
		return dec.String(&v.Test)
	}
	return nil
}

func (v *validationTestObj) NKeys() int {
	return 1
}

func strPtr() *string {
	s := ""
	return &s
}

func intPtr() *int {
	s := 0
	return &s
}

func TestGojaySchemaValidationStr(t *testing.T) {
	var testCases = []struct {
		name      string
		schema    string
		json      string
		v         interface{}
		expectedV interface{}
		err       bool
		errType   interface{}
	}{
		{
			name: "basic",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "string",
			}`,
			json:      `"test"`,
			expectedV: "test",
			v:         strPtr(),
		},
		{
			name: "basic-invalid",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "string",
			}`,
			json:      `1`,
			expectedV: "test",
			v:         strPtr(),
			err:       true,
			errType:   InvalidUnmarshalError(""),
		},
		{
			name: "basic-valid-pattern",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "string",
				"pattern": "^test$",
			}`,
			json:      `"test"`,
			expectedV: "test",
			v:         strPtr(),
			err:       false,
		},
		{
			name: "basic-invalid-pattern",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "string",
				"pattern": "^test$",
			}`,
			json:      `""`,
			expectedV: "test",
			v:         strPtr(),
			err:       true,
			errType:   SchemaValidationError(""),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var dec = NewDecoder(strings.NewReader(testCase.json))
			var sch, _ = NewSchema().Bytes([]byte(testCase.schema))
			var err = dec.DecodeAndValidate(testCase.v, sch)
			t.Log(err)
			if testCase.err {
				assert.NotNil(t, err)
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be a of the expected type")
				}
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGojaySchemaValidationInt(t *testing.T) {
	var testCases = []struct {
		name      string
		schema    string
		json      string
		v         interface{}
		expectedV interface{}
		err       bool
		errType   interface{}
	}{
		{
			name: "basic",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set int",
				"type": "number"
			}`,
			json:      `6`,
			expectedV: 6,
			v:         intPtr(),
		},
		{
			name: "basic-invalid",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set int",
				"type": "number"
			}`,
			json:      `"1"`,
			expectedV: 1,
			v:         intPtr(),
			err:       true,
			errType:   InvalidUnmarshalError(""),
		},
		{
			name: "basic-valid-number-max",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"maximum": 10
			}`,
			json:      `4`,
			expectedV: 4,
			v:         intPtr(),
			err:       false,
		},
		{
			name: "basic-valid-number-max-excl",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"maximum": 10,
				"exclusiveMax": false
			}`,
			json:      `10`,
			expectedV: 10,
			v:         intPtr(),
			err:       false,
		},
		{
			name: "basic-invalid-number-max",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"maximum": 10
			}`,
			json:      `11`,
			expectedV: 11,
			v:         intPtr(),
			err:       true,
			errType:   SchemaValidationError(""),
		},
		{
			name: "basic-invalid-number-max-excl",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"maximum": 10,
				"exclusiveMaximum": true
			}`,
			json:      `10`,
			expectedV: 10,
			v:         intPtr(),
			err:       true,
			errType:   SchemaValidationError(""),
		},
		{
			name: "basic-valid-number-min",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"minimum": 10
			}`,
			json:      `11`,
			expectedV: 11,
			v:         intPtr(),
			err:       false,
		},
		{
			name: "basic-valid-number-min-excl",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"minimum": 10
			}`,
			json:      `10`,
			expectedV: 10,
			v:         intPtr(),
			err:       false,
		},
		{
			name: "basic-invalid-number-min",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"minimum": 10
			}`,
			json:      `9`,
			expectedV: 9,
			v:         intPtr(),
			err:       true,
			errType:   SchemaValidationError(""),
		},
		{
			name: "basic-invalid-number-min-excl",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"minimum": 10,
				"exclusiveMinimum": true
			}`,
			json:      `10`,
			expectedV: 10,
			v:         intPtr(),
			err:       true,
			errType:   SchemaValidationError(""),
		},
		{
			name: "basic-multiple-of",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"multipleOf": 2
			}`,
			json:      `10`,
			expectedV: 10,
			v:         intPtr(),
		},
		{
			name: "basic-multiple-of-error",
			schema: `{
				"$schema": "http://json-schema.org/draft-06/schema#",
				"title": "Test set",
				"type": "number",
				"multipleOf": 2
			}`,
			json:      `11`,
			expectedV: 11,
			v:         intPtr(),
			err:       true,
			errType:   SchemaValidationError(""),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var dec = NewDecoder(strings.NewReader(testCase.json))
			var sch, err = NewSchema().Bytes([]byte(testCase.schema))
			if err != nil {
				log.Print(err)
				t.Error(err)
				return
			}
			err = dec.DecodeAndValidate(testCase.v, sch)
			t.Log(err)
			if testCase.err {
				assert.NotNil(t, err)
				if testCase.errType != nil {
					assert.IsType(t, testCase.errType, err, "err should be a of the expected type")
				}
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
