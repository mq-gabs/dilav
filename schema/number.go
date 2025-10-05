package schema

import (
	"errors"
	"fmt"
)

type numberSchema struct {
	baseSchema[int]
}

func Number() *numberSchema {
	return &numberSchema{
		baseSchema: newBaseSchema[int](),
	}
}

func (ns *numberSchema) Custom(fn Validator[int]) {
	ns.appendValidator(fn)
}

func (is *numberSchema) Min(min int) *numberSchema {
	is.appendValidator(func(i int) error {
		if i < min {
			return fmt.Errorf("required min value: %v", min)
		}

		return nil
	})

	return is
}

func (is *numberSchema) Max(max int) *numberSchema {
	is.appendValidator(func(i int) error {
		if i > max {
			return fmt.Errorf("required max value: %v", i)
		}

		return nil
	})

	return is
}

func (is *numberSchema) Equals(target int) *numberSchema {
	is.appendValidator(func(i int) error {
		if i != target {
			return fmt.Errorf("value must be equal to: %v", target)
		}

		return nil
	})

	return is
}

func (is *numberSchema) NonZero() *numberSchema {
	is.appendValidator(func(i int) error {
		if i == 0 {
			return errors.New("value must be non zero")
		}

		return nil
	})

	return is
}

func (is *numberSchema) Positive() *numberSchema {
	is.appendValidator(func(i int) error {
		if i < 0 {
			return errors.New("value must be positive")
		}

		return nil
	})

	return is
}

func (is *numberSchema) Negative() *numberSchema {
	is.appendValidator(func(i int) error {
		if i > 0 {
			return errors.New("value must be negative")
		}

		return nil
	})

	return is
}
