package typechecker

import (
	"flint/internal/parser"
)

var PlatformIntBits int = 0

func (tc *TypeChecker) resolveType(t parser.Expr) *Type {
	switch typ := t.(type) {
	case *parser.TypeExpr:
		switch typ.Name {
		case "Int":
			return &Type{TKind: TyInt}
		case "Float":
			return &Type{TKind: TyFloat}
		case "Bool":
			return &Type{TKind: TyBool}
		case "String":
			return &Type{TKind: TyString}
		case "Byte":
			return &Type{TKind: TyByte}
		case "Nil":
			return &Type{TKind: TyNil}
		case "List":
			elemTy := &Type{TKind: TyNil}
			if typ.Generic != nil {
				elemTy = tc.resolveType(typ.Generic)
			}
			return &Type{TKind: TyList, Elem: elemTy}
		}
	case *parser.TupleTypeExpr:
		elems := []*Type{}
		for _, te := range typ.Types {
			e := tc.resolveType(te)
			elems = append(elems, e)
		}
		return &Type{TKind: TyTuple, TElems: elems}
	}
	return &Type{TKind: TyError}
}
