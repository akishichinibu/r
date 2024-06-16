package gen

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func ErrorName(n int) string {
	return fmt.Sprintf("Error%d", n)
}

func ErrorTypeName(i int) string {
	return fmt.Sprintf("E%d", i)
}

func ErrorGenericParams(n int) []Code {
	genericParams := make([]Code, 0)
	for i := 1; i <= n; i++ {
		genericParams = append(genericParams, Id(ErrorTypeName(i)).Error())
	}
	return genericParams
}

func ErrorGenericIDs(n int) []Code {
	genericIDs := make([]Code, 0)
	for i := 1; i <= n; i++ {
		genericIDs = append(genericIDs, Id(ErrorTypeName(i)))
	}
	return genericIDs
}

func FailureMethodName(n int) string {
	return fmt.Sprintf("Failure%d", n)
}

func FromFailureMethodName(n int) string {
	return fmt.Sprintf("FromFailure%d", n)
}

func GenerateError(f *File, n int) {
	f.Commentf("%s is a type that can be used to represent an error with %d cases.", ErrorName(n), n)
	f.Type().
		Id(ErrorName(n)).
		Types(ErrorGenericParams(n)...).
		Struct(
			Id("i").
				Qual(MO_PACKAGE, EitherName(n)).
				Types(ErrorGenericIDs(n)...),
		)

	internalRef := Id("e").Dot("i")

	errBranches := make([]Code, n)
	if n == 2 {
		errBranches[0] = internalRef.Clone().Dot("Left").Call()
		errBranches[1] = internalRef.Clone().Dot("Right").Call()
	} else {
		for i := 1; i <= n; i++ {
			errBranches[i-1] = internalRef.Clone().Dot(fmt.Sprintf("Arg%d", i)).Call()
		}
	}

	receiver := Id("e").Id(ErrorName(n)).Types(ErrorGenericIDs(n)...)
	mutableReceiver := Id("e").Op("*").Id(ErrorName(n)).Types(ErrorGenericIDs(n)...)

	for i := 1; i <= n; i++ {
		f.Commentf("Return whether the error is in the case %d.", i)
		f.Func().
			Params(receiver).
			Id(FailureMethodName(i)).
			Params().
			Params(Id(ErrorTypeName(i)), Bool()).
			Block(Return(errBranches[i-1]))
	}

	for i := 1; i <= n; i++ {
		f.Commentf("Overwrite the error into the case %d with the given error.", i)
		f.Func().
			Params(mutableReceiver).
			Id(FromFailureMethodName(i)).
			Params(Id("err").Id(ErrorTypeName(i))).
			Block(
				internalRef.Clone().
					Op("=").
					Qual(MO_PACKAGE, NewEitherName(n, i)).Types(ErrorGenericIDs(n)...).
					Call(Id("err")),
			)
	}
}
