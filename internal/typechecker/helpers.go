package typechecker

import "flint/internal/parser"

func (tc *TypeChecker) resolveType(t parser.Expr) *Type {
	switch typ := t.(type) {
	case *parser.TypeExpr:
		switch typ.Name {
		case "Int":
			return &Type{kind: TyInt}
		case "Float":
			return &Type{kind: TyFloat}
		case "Bool":
			return &Type{kind: TyBool}
		case "String":
			return &Type{kind: TyString}
		case "Byte":
			return &Type{kind: TyByte}
		case "Nil":
			return &Type{kind: TyNil}
		case "List":
			elemTy := &Type{kind: TyNil}
			if typ.Generic != nil {
				elemTy = tc.resolveType(typ.Generic)
			}
			return &Type{kind: TyList, Elem: elemTy}
		}
	case *parser.TupleTypeExpr:
		elems := []*Type{}
		for _, te := range typ.Types {
			e := tc.resolveType(te)
			elems = append(elems, e)
		}
		return &Type{kind: TyTuple, TElems: elems}
	}
	return &Type{kind: TyError}
}
