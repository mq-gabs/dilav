package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	name string
	age  int
}

func (u User) SchemaJSON() map[string]any {
	return map[string]any{
		"name": u.name,
		"age":  u.age,
	}
}

func Test_Struct(t *testing.T) {

	u := User{
		name: "John Doe",
		age:  33,
	}

	validSchema := Struct(map[string]Schema[any]{
		"name": String().MinLength(4).MaxLength(32),
		"age":  Number().Max(44),
	})
	invalidSchema := Struct(map[string]Schema[any]{
		"name": String().MaxLength(4),
		"age":  Number().Min(44),
	})

	err1 := validSchema.Validate(u)
	t.Log(err1)
	err2 := invalidSchema.Validate(u)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
