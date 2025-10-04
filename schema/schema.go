package schema

import (
	"errors"
)

type Validator[T any] func(T) error

type Schema[T any] interface {
	Validate(T) error
}

var initialValidate = func() error {
	return errors.New("not implemented")
}

type baseSchema[T any] struct {
	validators []Validator[T]
}

func newBaseSchema[T any]() baseSchema[T] {
	return baseSchema[T]{}
}

func (bs *baseSchema[T]) Validate(value T) error {
	if len(bs.validators) == 0 {
		initialValidate()
	}

	var err error
	for _, valid := range bs.validators {
		e := valid(value)
		if e != nil {
			err = errors.Join(err, e)
		}
	}

	return err
}

func (bs *baseSchema[T]) appendValidator(newValidator Validator[T]) {
	bs.validators = append(bs.validators, newValidator)
}
