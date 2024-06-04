package main

import (
	"github.com/akishichinibu/r/gen"
	. "github.com/dave/jennifer/jen"
)

func genError() {
	f := NewFile("r")
	for i := 2; i <= 5; i++ {
		gen.GenerateError(f, i)
	}
	f.Save("error.gen.go")
}

func genResult() {
	f := NewFile("r")
	for i := 2; i <= 5; i++ {
		gen.GenerateResult(f, i)
	}
	f.Save("result.gen.go")
}

func main() {
	genError()
	genResult()
}
