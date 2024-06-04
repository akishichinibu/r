package r

type Error1[E1 error] struct {
	i E1
}

func (e Error1[E1]) Failure() (E1, bool) {
	return e.i, true
}
