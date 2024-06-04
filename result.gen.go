package r

import mo "github.com/samber/mo"

type Result2[T any, E1 error, E2 error] struct {
	i mo.Either[T, Error2[E1, E2]]
}

func (r Result2[T, E1, E2]) IsSuccess() bool {
	return r.i.IsLeft()
}
func (r Result2[T, E1, E2]) IsFailure() bool {
	return r.i.IsRight()
}
func (r *Result2[T, E1, E2]) FromSuccess(value T) {
	r.i = mo.Left[T, Error2[E1, E2]](value)
}
func (r *Result2[T, E1, E2]) FromFailure1(err E1) {
	var e Error2[E1, E2]
	e.FromFailure1(err)
	r.i = mo.Right[T, Error2[E1, E2]](e)
}
func (r *Result2[T, E1, E2]) FromFailure2(err E2) {
	var e Error2[E1, E2]
	e.FromFailure2(err)
	r.i = mo.Right[T, Error2[E1, E2]](e)
}
func (r Result2[T, E1, E2]) Unpack() (T, Error2[E1, E2], bool) {
	ok := r.IsSuccess()
	result, err := r.i.Unpack()
	return result, err, ok
}

type Result3[T any, E1 error, E2 error, E3 error] struct {
	i mo.Either[T, Error3[E1, E2, E3]]
}

func (r Result3[T, E1, E2, E3]) IsSuccess() bool {
	return r.i.IsLeft()
}
func (r Result3[T, E1, E2, E3]) IsFailure() bool {
	return r.i.IsRight()
}
func (r *Result3[T, E1, E2, E3]) FromSuccess(value T) {
	r.i = mo.Left[T, Error3[E1, E2, E3]](value)
}
func (r *Result3[T, E1, E2, E3]) FromFailure1(err E1) {
	var e Error3[E1, E2, E3]
	e.FromFailure1(err)
	r.i = mo.Right[T, Error3[E1, E2, E3]](e)
}
func (r *Result3[T, E1, E2, E3]) FromFailure2(err E2) {
	var e Error3[E1, E2, E3]
	e.FromFailure2(err)
	r.i = mo.Right[T, Error3[E1, E2, E3]](e)
}
func (r *Result3[T, E1, E2, E3]) FromFailure3(err E3) {
	var e Error3[E1, E2, E3]
	e.FromFailure3(err)
	r.i = mo.Right[T, Error3[E1, E2, E3]](e)
}
func (r Result3[T, E1, E2, E3]) Unpack() (T, Error3[E1, E2, E3], bool) {
	ok := r.IsSuccess()
	result, err := r.i.Unpack()
	return result, err, ok
}

type Result4[T any, E1 error, E2 error, E3 error, E4 error] struct {
	i mo.Either[T, Error4[E1, E2, E3, E4]]
}

func (r Result4[T, E1, E2, E3, E4]) IsSuccess() bool {
	return r.i.IsLeft()
}
func (r Result4[T, E1, E2, E3, E4]) IsFailure() bool {
	return r.i.IsRight()
}
func (r *Result4[T, E1, E2, E3, E4]) FromSuccess(value T) {
	r.i = mo.Left[T, Error4[E1, E2, E3, E4]](value)
}
func (r *Result4[T, E1, E2, E3, E4]) FromFailure1(err E1) {
	var e Error4[E1, E2, E3, E4]
	e.FromFailure1(err)
	r.i = mo.Right[T, Error4[E1, E2, E3, E4]](e)
}
func (r *Result4[T, E1, E2, E3, E4]) FromFailure2(err E2) {
	var e Error4[E1, E2, E3, E4]
	e.FromFailure2(err)
	r.i = mo.Right[T, Error4[E1, E2, E3, E4]](e)
}
func (r *Result4[T, E1, E2, E3, E4]) FromFailure3(err E3) {
	var e Error4[E1, E2, E3, E4]
	e.FromFailure3(err)
	r.i = mo.Right[T, Error4[E1, E2, E3, E4]](e)
}
func (r *Result4[T, E1, E2, E3, E4]) FromFailure4(err E4) {
	var e Error4[E1, E2, E3, E4]
	e.FromFailure4(err)
	r.i = mo.Right[T, Error4[E1, E2, E3, E4]](e)
}
func (r Result4[T, E1, E2, E3, E4]) Unpack() (T, Error4[E1, E2, E3, E4], bool) {
	ok := r.IsSuccess()
	result, err := r.i.Unpack()
	return result, err, ok
}

type Result5[T any, E1 error, E2 error, E3 error, E4 error, E5 error] struct {
	i mo.Either[T, Error5[E1, E2, E3, E4, E5]]
}

func (r Result5[T, E1, E2, E3, E4, E5]) IsSuccess() bool {
	return r.i.IsLeft()
}
func (r Result5[T, E1, E2, E3, E4, E5]) IsFailure() bool {
	return r.i.IsRight()
}
func (r *Result5[T, E1, E2, E3, E4, E5]) FromSuccess(value T) {
	r.i = mo.Left[T, Error5[E1, E2, E3, E4, E5]](value)
}
func (r *Result5[T, E1, E2, E3, E4, E5]) FromFailure1(err E1) {
	var e Error5[E1, E2, E3, E4, E5]
	e.FromFailure1(err)
	r.i = mo.Right[T, Error5[E1, E2, E3, E4, E5]](e)
}
func (r *Result5[T, E1, E2, E3, E4, E5]) FromFailure2(err E2) {
	var e Error5[E1, E2, E3, E4, E5]
	e.FromFailure2(err)
	r.i = mo.Right[T, Error5[E1, E2, E3, E4, E5]](e)
}
func (r *Result5[T, E1, E2, E3, E4, E5]) FromFailure3(err E3) {
	var e Error5[E1, E2, E3, E4, E5]
	e.FromFailure3(err)
	r.i = mo.Right[T, Error5[E1, E2, E3, E4, E5]](e)
}
func (r *Result5[T, E1, E2, E3, E4, E5]) FromFailure4(err E4) {
	var e Error5[E1, E2, E3, E4, E5]
	e.FromFailure4(err)
	r.i = mo.Right[T, Error5[E1, E2, E3, E4, E5]](e)
}
func (r *Result5[T, E1, E2, E3, E4, E5]) FromFailure5(err E5) {
	var e Error5[E1, E2, E3, E4, E5]
	e.FromFailure5(err)
	r.i = mo.Right[T, Error5[E1, E2, E3, E4, E5]](e)
}
func (r Result5[T, E1, E2, E3, E4, E5]) Unpack() (T, Error5[E1, E2, E3, E4, E5], bool) {
	ok := r.IsSuccess()
	result, err := r.i.Unpack()
	return result, err, ok
}