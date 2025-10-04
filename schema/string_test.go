package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	name := "John"

	validSchema := String().MinLength(2).MaxLength(8)
	invalidSchema := String().MinLength(5).MaxLength(8)

	err1 := validSchema.Validate(name)
	t.Log(err1)
	err2 := invalidSchema.Validate(name)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
