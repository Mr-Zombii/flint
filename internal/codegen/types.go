package codegen

import (
	"flint/internal/parser"

	"github.com/llir/llvm/ir/types"
)

func (cg *CodeGen) resolveType(t parser.Expr) types.Type {
	if t == nil {
		return types.Void
	}
	ty := t.(*parser.TypeExpr)
	switch ty.Name {
	case "Int":
		return cg.platformIntType()
	case "Float":
		return cg.platformFloatType()
	case "Bool":
		return types.I1
	case "Byte":
		return types.I8
	case "String":
		return types.I8Ptr
	case "Nil":
		return types.Void
	default:
		return nil
	}
}
