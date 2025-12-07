package codegen

import (
	"reflect"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

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
