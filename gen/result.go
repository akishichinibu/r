package gen

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func ResultName(n int) string {
	return fmt.Sprintf("Result%d", n)
}

func ValueTypeName() string {
	return "T"
}

func GenerateResult(f *File, n int) {
	tParams := make([]Code, 0)
	tParams = append(tParams, Id(ValueTypeName()).Any())
	tParams = append(tParams, ErrorGenericParams(n)...)

	t := make([]Code, 0)
	t = append(t, Id(ValueTypeName()))
	t = append(t, ErrorGenericIDs(n)...)

	errorType := Id(ErrorName(n)).Types(ErrorGenericIDs(n)...)

	f.Commentf("%s is a type that can be used to represent a result that contains a value in success case and %d error types in failure case.", ResultName(n), n)
	f.Type().
		Id(ResultName(n)).Types(tParams...).
		Struct(
			Id("i").
				Qual(MO_PACKAGE, "Either").
				Types(Id(ValueTypeName()), errorType),
		)

	receiver := Id("r").Id(ResultName(n)).Types(t...)
	mutableReceiver := Id("r").Op("*").Id(ResultName(n)).Types(t...)
	internalRef := Id("r").Dot("i")

	f.Commentf("Return whether the result is the success case.")
	f.Func().
		Params(receiver).
		Id("IsSuccess").
		Params().
		Bool().
		Block(
			Return(internalRef.Clone().Dot("IsLeft").Call()),
		)

	f.Commentf("Return whether the result is the failure case.")
	f.Func().
		Params(receiver).
		Id("IsFailure").
		Params().
		Bool().
		Block(
			Return(internalRef.Clone().Dot("IsRight").Call()),
		)

	f.Commentf("Overwrite the value of the result with the given value into the success case.")
	f.Func().
		Params(mutableReceiver).
		Id("FromSuccess").
		Params(Id("value").Id(ValueTypeName())).
		Block(
			internalRef.Clone().Op("=").Qual(MO_PACKAGE, "Left").
				Types(Id(ValueTypeName()), errorType).
				Call(Id("value")),
		)

	for i := 1; i <= n; i++ {
		f.Commentf("Overwrite the error of the result with the given error into the failure case %d.", i)
		f.Func().
			Params(mutableReceiver).
			Id(FromFailureMethodName(i)).
			Params(Id("err").Id(ErrorTypeName(i))).
			Block(
				Var().Id("e").Id(ErrorName(n)).Types(ErrorGenericIDs(n)...),
				Id("e").Dot(FromFailureMethodName(i)).Call(Id("err")),
				internalRef.Clone().
					Op("=").
					Qual(MO_PACKAGE, "Right").
					Types(Id(ValueTypeName()), Id(ErrorName(n)).Types(ErrorGenericIDs(n)...)).
					Call(Id("e")),
			)
	}

	f.Commentf("Unpack the result into the value, the error, and whether the result is the success case.")
	f.Func().
		Params(receiver).
		Id("Unpack").
		Params().
		Params(Id("T"), errorType, Bool()).
		Block(
			Id("ok").Op(":=").Id("r").Dot("IsSuccess").Call(),
			List(Id("result"), Id("err")).Op(":=").Id("r").Dot("i").Dot("Unpack").Call(),
			Return(Id("result"), Id("err"), Id("ok")),
		)
}
