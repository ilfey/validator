package test

import (
	"testing"

	"github.com/ilfey/validator"
	"github.com/stretchr/testify/assert"
)

type Test struct {
	NonPtr   string  `json:"non_ptr"`
	Ptr      *string `json:"ptr,omitempty"`
	LessThan int     `json:"less_than"`
}

func TestValidateField(t *testing.T) {
	st := Test{
		NonPtr:   "test123",
		Ptr:      nil,
		LessThan: 50,
	}

	err, _ := validator.ValidateField(st, "NonPtr", validator.Required())

	assert.Nil(t, err)

	err, _ = validator.ValidateField(st, "Ptr", validator.Required())

	assert.NotNil(t, err)

	err, _ = validator.ValidateField(st, "LessThan", validator.LessThanInt(20))

	assert.NotNil(t, err)

	err, _ = validator.ValidateField(st, "LessThan", validator.LessThanInt(60))

	assert.Nil(t, err)
}

func TestValidateStruct(t *testing.T) {
	st := Test{
		NonPtr:   "test123",
		Ptr:      nil,
		LessThan: 50,
	}

	err := validator.ValidateStruct(st, map[string][]validator.Option{
		"NonPtr":   {validator.Required()},
		"Ptr":      {validator.Required()},
		"LessThan": {validator.LessThanInt(20), validator.LessThanInt(60)},
	})

	assert.NotNil(t, err)
}
