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

	validSchema := Struct[User](map[string]Schema[any]{
		"name": String().MinLength(4).MaxLength(32),
		"age":  Number[int]().Max(44),
	})
	invalidSchema := Struct[User](map[string]Schema[any]{
		"name": String().MaxLength(4),
		"age":  Number[int]().Min(44),
	})

	err1 := validSchema.Validate(u)
	t.Log(err1)
	err2 := invalidSchema.Validate(u)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}

type Group struct {
	name   string
	member *User
}

func (g Group) SchemaJSON() map[string]any {
	return map[string]any{
		"name":   g.name,
		"member": g.member,
	}
}

func Test_StructInsideStruct(t *testing.T) {
	u := User{
		name: "Bob Smith",
		age:  23,
	}

	g := Group{
		name:   "Worker",
		member: &u,
	}

	validUserSchema := Struct[User](map[string]Schema[any]{
		"name": String().MaxLength(32),
		"age":  Number[int]().Min(18),
	})
	invalidUserSchema := Struct[User](map[string]Schema[any]{
		"name": String().MinLength(32),
		"age":  Number[int]().Max(21),
	})

	validSchema := Struct[Group](map[string]Schema[any]{
		"name":   String().MaxLength(12),
		"member": validUserSchema,
	})
	invalidSchema := Struct[Group](map[string]Schema[any]{
		"name":   String().MaxLength(4),
		"member": invalidUserSchema,
	})

	err1 := validSchema.Validate(g)
	t.Log(err1)
	err2 := invalidSchema.Validate(g)
	t.Log(err2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}

type regexTestCase struct {
	value  string
	valid  bool
	schema *stringSchema
}

func Test_Regex(t *testing.T) {
	emailSchema := String().Email()
	uuidSchema := String().UUID()
	urlSchema := String().URL()

	var tests = []regexTestCase{
		// Email tests
		{"user@example.com", true, emailSchema},
		{"first.last@domain.co.uk", true, emailSchema},
		{"user+tag@sub.domain.org", true, emailSchema},
		{"user@", false, emailSchema},
		{"@example.com", false, emailSchema},
		{"user@.com", false, emailSchema},
		{"user@domain", false, emailSchema},
		{"plainaddress", false, emailSchema},
		{"", false, emailSchema},

		// UUID tests
		{"123e4567-e89b-12d3-a456-426614174000", true, uuidSchema},
		{"550e8400-e29b-41d4-a716-446655440000", true, uuidSchema},
		{"550e8400e29b41d4a716446655440000", false, uuidSchema},
		{"g23e4567-e89b-12d3-a456-426614174000", false, uuidSchema},
		{"123e4567-e89b-12d3-a456", false, uuidSchema},
		{"", false, uuidSchema},

		// URL tests
		{"http://example.com", true, urlSchema},
		{"https://sub.domain.com/path?query=1#fragment", true, urlSchema},
		{"ftp://files.example.com/resource.zip", true, urlSchema},
		{"http://localhost:8080/path", true, urlSchema},
		{"https://192.168.1.1:8080", true, urlSchema},
		{"example.com", false, urlSchema},
		{"http:/example.com", false, urlSchema},
		{"://example.com", false, urlSchema},
		{"", false, urlSchema},
	}
	for _, tt := range tests {
		isValid := tt.schema.Validate(tt.value) == nil

		assert.Equal(t, isValid, tt.valid)
	}
}
