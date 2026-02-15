package schema

import (
	"errors"
	"fmt"
)

type SchemaMap[T comparable, U any] struct {
	baseSchema[map[T]U]
}

func Map[T comparable, U any]() *SchemaMap[T, U] {
	return &SchemaMap[T, U]{
		baseSchema: newBaseSchema[map[T]U](),
	}
}

func (ms *SchemaMap[T, U]) Custom(fn Validator[map[T]U]) *SchemaMap[T, U] {
	ms.appendValidator(fn)

	return ms
}

func (ms *SchemaMap[T, U]) LengthMax(max int) *SchemaMap[T, U] {
	ms.appendValidator(func(m map[T]U) error {
		if len(m) > max {
			return fmt.Errorf("required max length: %v", max)
		}

		return nil
	})

	return ms
}

func (ms *SchemaMap[T, U]) LengthMin(min int) *SchemaMap[T, U] {
	ms.appendValidator(func(m map[T]U) error {
		if len(m) < min {
			return fmt.Errorf("required min length: %v", min)
		}

		return nil
	})

	return ms
}

func (ms *SchemaMap[T, U]) Child(schema Schema[U]) *SchemaMap[T, U] {
	ms.appendValidator(func(m map[T]U) error {
		var err error
		for key, value := range m {
			schemaErr := schema.Validate(value)
			if schemaErr != nil {
				err = errors.Join(err, fmt.Errorf("[%v]: %v", key, schemaErr.Error()))
			}
		}

		return err
	})

	return ms
}
