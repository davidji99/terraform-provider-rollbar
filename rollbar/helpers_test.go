package rollbar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToInt_Basic(t *testing.T) {
	assert.Equal(t, 123, StringToInt("123"))
}

func TestStringToInt_Unique(t *testing.T) {
	assert.Equal(t, 0, StringToInt("12@3"))
}

func TestInt64ToString_Basic(t *testing.T) {
	assert.Equal(t, "355432423", Int64ToString(int64(355432423)))
}

func TestParseCompositeID_Valid(t *testing.T) {
	s1, s2, parseErr := ParseCompositeImportID("hello:moto")

	assert.Nil(t, parseErr)
	assert.Equal(t, "hello", s1)
	assert.Equal(t, "moto", s2)
}

func TestParseCompositeID_Invalid(t *testing.T) {
	_, _, parseErr := ParseCompositeImportID("hello:moto:again")
	assert.NotNil(t, parseErr)
}
