package gen

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
)

func EitherName(n int) string {
	if n == 2 {
		return "Either"
	}
	return fmt.Sprintf("Either%d", n)
}

func NewEitherName(n int, i int) string {
	if n == 2 {
		if i == 1 {
			return "Left"
		}
		return "Right"
	}
	return fmt.Sprintf("NewEither%dArg%d", n, i)
}

func Branches(n int) []j.Code {
	branches := make([]j.Code, n)
	if n == 2 {
		branches[0] = j.Id("e").Dot("i").Dot("Left").Call()
		branches[1] = j.Id("e").Dot("i").Dot("Right").Call()
	} else {
		for i := 1; i <= n; i++ {
			branches[i-1] = j.Id("e").Dot("i").Dot(fmt.Sprintf("Arg%d", i)).Call()
		}
	}
	return branches
}

const MO_PACKAGE = "github.com/samber/mo"
