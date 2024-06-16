package r

// Error is a type that can be used to represent an error with only one case.
type Error1[E1 error] struct {
	i E1
}

// Return whether the error is in the first case.
func (e Error1[E1]) Failure() (E1, bool) {
	return e.i, true
}

// Overwrite the error into the first case with the given error.
func (e *Error1[E1]) FromFailure(err E1) {
	e.i = err
}
