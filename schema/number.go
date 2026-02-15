package schema

import (
	"errors"
	"fmt"
)

type NumberType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type SchemaNumber[T NumberType] struct {
	baseSchema[T]
}

func Number[T NumberType]() *SchemaNumber[T] {
	return &SchemaNumber[T]{
		baseSchema: newBaseSchema[T](),
	}
}

func (ns *SchemaNumber[T]) Custom(fn Validator[T]) *SchemaNumber[T] {
	ns.appendValidator(fn)

	return ns
}

func (ns *SchemaNumber[T]) Min(min T) *SchemaNumber[T] {
	ns.appendValidator(func(i T) error {
		if i < min {
			return fmt.Errorf("required min value: %v", min)
		}

		return nil
	})

	return ns
}

func (is *SchemaNumber[T]) Max(max T) *SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i > max {
			return fmt.Errorf("required max value: %v", i)
		}

		return nil
	})

	return is
}

func (is *SchemaNumber[T]) Equals(target T) *SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i != target {
			return fmt.Errorf("value must be equal to: %v", target)
		}

		return nil
	})

	return is
}

func (is *SchemaNumber[T]) NonZero() *SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i == 0 {
			return errors.New("value must be non zero")
		}

		return nil
	})

	return is
}

func (is *SchemaNumber[T]) Positive() *SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i < 0 {
			return errors.New("value must be positive")
		}

		return nil
	})

	return is
}

func (is *SchemaNumber[T]) Negative() *SchemaNumber[T] {
	is.appendValidator(func(i T) error {
		if i > 0 {
			return errors.New("value must be negative")
		}

		return nil
	})

	return is
}
