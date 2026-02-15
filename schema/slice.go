package schema

import (
	"fmt"
	"slices"
)

type SchemaSlice[T comparable] struct {
	baseSchema[[]T]
}

func Slice[T comparable]() *SchemaSlice[T] {
	return &SchemaSlice[T]{
		baseSchema: newBaseSchema[[]T](),
	}
}

func (ss *SchemaSlice[T]) Custom(fn Validator[[]T]) *SchemaSlice[T] {
	ss.appendValidator(fn)

	return ss
}

func (ss *SchemaSlice[T]) LengthMin(minLen int) *SchemaSlice[T] {
	ss.appendValidator(func(a []T) error {
		if len(a) < minLen {
			return fmt.Errorf("required min length: %v", minLen)
		}

		return nil
	})

	return ss
}

func (ss *SchemaSlice[T]) LengthMax(maxLen int) *SchemaSlice[T] {
	ss.appendValidator(func(a []T) error {
		if len(a) > maxLen {
			return fmt.Errorf("required max length: %v", maxLen)
		}

		return nil
	})

	return ss
}

func (ss *SchemaSlice[T]) Contains(target T) *SchemaSlice[T] {
	ss.appendValidator(func(t []T) error {
		if !slices.Contains(t, target) {
			return fmt.Errorf("slice must contain value: %v", target)
		}

		return nil
	})

	return ss
}
