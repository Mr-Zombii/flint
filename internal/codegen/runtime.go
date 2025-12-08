package codegen

import (
	"flint/internal/parser"
	"reflect"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/value"
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

func parentBlockOfValue(v value.Value) *ir.Block {
	if v == nil {
		return nil
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return nil
	}
	elem := rv.Elem()
	if elem.Kind() != reflect.Struct {
		return nil
	}
	field := elem.FieldByName("Parent")
	if !field.IsValid() || field.IsNil() {
		return nil
	}
	parent, ok := field.Interface().(*ir.Block)
	if !ok {
		return nil
	}
	return parent
}

func referencesBlock(b *ir.Block, blockToFind *ir.Block) bool {
	if b != nil && b.Term != nil {
		opLen := len(b.Term.Operands())
		for i := range opLen {
			op := *b.Term.Operands()[i]
			v, ok := op.(*ir.Block)
			if ok && v == blockToFind {
				return true
			}
		}
	}
	return false
}
