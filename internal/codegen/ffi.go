package codegen

import (
	"flint/internal/parser"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
)

func (cg *CodeGen) emitExternalFunction(fn *parser.FuncDeclExpr, externDecorationIdx int, name string, mainFn *ir.Func) {
	languageIdentifier := fn.Decorators[externDecorationIdx].Args[0].(*parser.Identifier).Name
	// libraryName := fn.Decorators[externDecorationIdx].Args[1].(*parser.StringLiteral).Value
	functionName := fn.Decorators[externDecorationIdx].Args[2].(*parser.StringLiteral).Value
	if name != functionName {
		panic("Function must have same name as external annotation, " + name + " != " + functionName)
	}
	if languageIdentifier == "c" {
		mainFn.CallingConv = enum.CallingConvC
	}
	mainFn.Linkage = enum.LinkageExternal
	for _, param := range mainFn.Params {
		param.SetName("")
	}
}
