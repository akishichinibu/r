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

func ErrorGenericIds(n int) []Code {
	genericIds := make([]Code, 0)
	for i := 1; i <= n; i++ {
		genericIds = append(genericIds, Id(ErrorTypeName(i)))
	}
	return genericIds
}

func FailureMethodName(n int) string {
	return fmt.Sprintf("Failure%d", n)
}

func FromFailureMethodName(n int) string {
	return fmt.Sprintf("FromFailure%d", n)
}

func GenerateError(f *File, n int) {
	f.Type().
		Id(ErrorName(n)).Types(ErrorGenericParams(n)...).
		Struct(
			Id("i").
				Qual(MO_PACKAGE, EitherName(n)).
				Types(ErrorGenericIds(n)...),
		)

	internalRef := Id("e").Dot("i")

	errBranches := make([]Code, n)
	if n == 2 {
		errBranches[0] = internalRef.Dot("Left").Call()
		errBranches[1] = internalRef.Dot("Right").Call()
	} else {
		for i := 1; i <= n; i++ {
			errBranches[i-1] = internalRef.Dot(fmt.Sprintf("Arg%d", i)).Call()
		}
	}

	receiver := Id("e").Id(ErrorName(n)).Types(ErrorGenericIds(n)...)
	mutableReceiver := Id("e").Op("*").Id(ErrorName(n)).Types(ErrorGenericIds(n)...)

	for i := 1; i <= n; i++ {
		f.Func().
			Params(receiver).
			Id(FailureMethodName(i)).
			Params().
			Params(Id(ErrorTypeName(i)), Bool()).
			Block(Return(errBranches[i-1]))
	}

	for i := 1; i <= n; i++ {
		f.Func().
			Params(mutableReceiver).
			Id(FromFailureMethodName(i)).
			Params(Id("err").Id(ErrorTypeName(i))).
			Block(
				internalRef.Op("=").
					Qual(MO_PACKAGE, NewEitherName(n, i)).Types(ErrorGenericIds(n)...).
					Call(Id("err")),
			)
	}
}
