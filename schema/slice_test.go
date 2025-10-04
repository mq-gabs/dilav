package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slice(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	validSchema := Slice[int]().MinLength(2).MaxLength(8)
	invalidSchema := Slice[int]().MinLength(6).MaxLength(10)

	err1 := validSchema.Validate(nums)
	t.Log(err1)
	err2 := invalidSchema.Validate(nums)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
