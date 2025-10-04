package schema

import "fmt"

type sliceSchema[T any] struct {
	baseSchema[[]T]
}

func Slice[T any]() *sliceSchema[T] {
	return &sliceSchema[T]{
		baseSchema: newBaseSchema[[]T](),
	}
}

func (ss *sliceSchema[T]) MinLength(minLen int) *sliceSchema[T] {
	ss.appendValidator(func(a []T) error {
		if len(a) < minLen {
			return fmt.Errorf("required min length: %v", minLen)
		}

		return nil
	})

	return ss
}

func (ss *sliceSchema[T]) MaxLength(maxLen int) *sliceSchema[T] {
	ss.appendValidator(func(a []T) error {
		if len(a) > maxLen {
			return fmt.Errorf("required max length: %v", maxLen)
		}

		return nil
	})

	return ss
}
