package r

import (
	mo "github.com/samber/mo"
)

// Result1 is a type that can be used to represent a result
// that contains a value in success case and one error in failure case.
type Result1[T any, E1 error] struct {
	i mo.Either[T, Error1[E1]]
}

// Overwrite the value of the result with the given value into the success case.
func (r *Result1[T, E1]) FromSuccess(value T) {
	r.i = mo.Left[T, Error1[E1]](value)
}

// Overwrite the error of the result with the given error into the first failure case.
func (r *Result1[T, E1]) FromFailure(err E1) {
	e := Error1[E1]{}
	e.FromFailure(err)
	r.i = mo.Right[T, Error1[E1]](e)
}

// Return whether the result is the success case.
func (r Result1[T, E1]) IsSuccess() bool {
	return r.i.IsLeft()
}

// Return whether the result is the failure case.
func (r Result1[T, E1]) IsFailure() bool {
	return r.i.IsRight()
}

// Unpack the result into the value, the error, and whether the result is the success case.
func (r Result1[T, E1]) Unpack() (value T, err Error1[E1], ok bool) {
	ok = r.IsSuccess()
	value, err = r.i.Unpack()
	return value, err, ok
}
