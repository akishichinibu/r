package r

import mo "github.com/samber/mo"

type Error2[E1 error, E2 error] struct {
	i mo.Either[E1, E2]
}

func (e Error2[E1, E2]) Failure1() (E1, bool) {
	return e.i.Left()
}
func (e Error2[E1, E2]) Failure2() (E2, bool) {
	return e.i.Right()
}
func (e *Error2[E1, E2]) FromFailure1(err E1) {
	e.i = mo.Left[E1, E2](err)
}
func (e *Error2[E1, E2]) FromFailure2(err E2) {
	e.i = mo.Right[E1, E2](err)
}

type Error3[E1 error, E2 error, E3 error] struct {
	i mo.Either3[E1, E2, E3]
}

func (e Error3[E1, E2, E3]) Failure1() (E1, bool) {
	return e.i.Arg1()
}
func (e Error3[E1, E2, E3]) Failure2() (E2, bool) {
	return e.i.Arg2()
}
func (e Error3[E1, E2, E3]) Failure3() (E3, bool) {
	return e.i.Arg3()
}
func (e *Error3[E1, E2, E3]) FromFailure1(err E1) {
	e.i = mo.NewEither3Arg1[E1, E2, E3](err)
}
func (e *Error3[E1, E2, E3]) FromFailure2(err E2) {
	e.i = mo.NewEither3Arg2[E1, E2, E3](err)
}
func (e *Error3[E1, E2, E3]) FromFailure3(err E3) {
	e.i = mo.NewEither3Arg3[E1, E2, E3](err)
}

type Error4[E1 error, E2 error, E3 error, E4 error] struct {
	i mo.Either4[E1, E2, E3, E4]
}

func (e Error4[E1, E2, E3, E4]) Failure1() (E1, bool) {
	return e.i.Arg1()
}
func (e Error4[E1, E2, E3, E4]) Failure2() (E2, bool) {
	return e.i.Arg2()
}
func (e Error4[E1, E2, E3, E4]) Failure3() (E3, bool) {
	return e.i.Arg3()
}
func (e Error4[E1, E2, E3, E4]) Failure4() (E4, bool) {
	return e.i.Arg4()
}
func (e *Error4[E1, E2, E3, E4]) FromFailure1(err E1) {
	e.i = mo.NewEither4Arg1[E1, E2, E3, E4](err)
}
func (e *Error4[E1, E2, E3, E4]) FromFailure2(err E2) {
	e.i = mo.NewEither4Arg2[E1, E2, E3, E4](err)
}
func (e *Error4[E1, E2, E3, E4]) FromFailure3(err E3) {
	e.i = mo.NewEither4Arg3[E1, E2, E3, E4](err)
}
func (e *Error4[E1, E2, E3, E4]) FromFailure4(err E4) {
	e.i = mo.NewEither4Arg4[E1, E2, E3, E4](err)
}

type Error5[E1 error, E2 error, E3 error, E4 error, E5 error] struct {
	i mo.Either5[E1, E2, E3, E4, E5]
}

func (e Error5[E1, E2, E3, E4, E5]) Failure1() (E1, bool) {
	return e.i.Arg1()
}
func (e Error5[E1, E2, E3, E4, E5]) Failure2() (E2, bool) {
	return e.i.Arg2()
}
func (e Error5[E1, E2, E3, E4, E5]) Failure3() (E3, bool) {
	return e.i.Arg3()
}
func (e Error5[E1, E2, E3, E4, E5]) Failure4() (E4, bool) {
	return e.i.Arg4()
}
func (e Error5[E1, E2, E3, E4, E5]) Failure5() (E5, bool) {
	return e.i.Arg5()
}
func (e *Error5[E1, E2, E3, E4, E5]) FromFailure1(err E1) {
	e.i = mo.NewEither5Arg1[E1, E2, E3, E4, E5](err)
}
func (e *Error5[E1, E2, E3, E4, E5]) FromFailure2(err E2) {
	e.i = mo.NewEither5Arg2[E1, E2, E3, E4, E5](err)
}
func (e *Error5[E1, E2, E3, E4, E5]) FromFailure3(err E3) {
	e.i = mo.NewEither5Arg3[E1, E2, E3, E4, E5](err)
}
func (e *Error5[E1, E2, E3, E4, E5]) FromFailure4(err E4) {
	e.i = mo.NewEither5Arg4[E1, E2, E3, E4, E5](err)
}
func (e *Error5[E1, E2, E3, E4, E5]) FromFailure5(err E5) {
	e.i = mo.NewEither5Arg5[E1, E2, E3, E4, E5](err)
}
