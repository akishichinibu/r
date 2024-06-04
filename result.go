package r

import (
	mo "github.com/samber/mo"
)

type Result1[T any, E1 error] struct {
	i mo.Either[T, Error1[E1]]
}

func (r *Result1[T, E1]) FromSuccess(value T) {
	r.i = mo.Left[T, Error1[E1]](value)
}

func (r *Result1[T, E1]) FromFailure(err E1) {
	r.i = mo.Right[T, Error1[E1]](Error1[E1]{err})
}

func (r Result1[T, E1]) IsSuccess() bool {
	return r.i.IsLeft()
}

func (r Result1[T, E1]) IsFailure() bool {
	return r.i.IsRight()
}

func (r Result1[T, E1]) Unpack() (T, Error1[E1], bool) {
	ok := r.IsSuccess()
	result, err := r.i.Unpack()
	return result, err, ok
}
