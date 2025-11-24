package typechecker

import (
	"fmt"
	"strings"
)

type TypeKind int

type Type struct {
	kind   TypeKind
	Params []*Type
	Ret    *Type
	Elem   *Type
	TElems []*Type
}

const (
	TyError TypeKind = iota
	TyInt
	TyFloat
	TyBool
	TyByte
	TyString
	TyNil
	TyFunc
	TyList
	TyTuple
	TyRange
)

func (t Type) String() string {
	switch t.kind {
	case TyInt:
		return "Int"
	case TyFloat:
		return "Float"
	case TyBool:
		return "Bool"
	case TyString:
		return "String"
	case TyByte:
		return "Byte"
	case TyNil:
		return "Nil"
	case TyList:
		if t.Elem != nil {
			return fmt.Sprintf("List(%s)", t.Elem.String())
		}
		return "List(<unknown>)"
	case TyTuple:
		parts := []string{}
		for _, e := range t.TElems {
			if e == nil {
				parts = append(parts, "<unknown>")
			} else {
				parts = append(parts, e.String())
			}
		}
		return fmt.Sprintf("(%s)", strings.Join(parts, ", "))
	case TyRange:
		if t.Elem != nil {
			return fmt.Sprintf("Range(%s)", t.Elem.String())
		}
		return "Range(Int)"
	case TyFunc:
		parts := []string{}
		for _, p := range t.Params {
			parts = append(parts, p.String())
		}
		return fmt.Sprintf("(%s) -> %s", strings.Join(parts, ", "), t.Ret.String())
	}
	return "<error>"
}

func (t Type) Kind() TypeKind { return t.kind }

func (t *Type) Equal(u *Type) bool {
	if t == nil || u == nil {
		return t == u
	}
	if t.kind != u.kind {
		return false
	}
	switch t.kind {
	case TyFunc:
		if len(t.Params) != len(u.Params) {
			return false
		}
		for i := range t.Params {
			if !t.Params[i].Equal(u.Params[i]) {
				return false
			}
		}
		return t.Ret.Equal(u.Ret)
	case TyList:
		if t.Elem == nil || u.Elem == nil {
			return t.Elem == u.Elem
		}
		return t.Elem.Equal(u.Elem)
	case TyTuple:
		if len(t.TElems) != len(u.TElems) {
			return false
		}
		for i := range t.TElems {
			if !t.TElems[i].Equal(u.TElems[i]) {
				return false
			}
		}
		return true
	case TyRange:
		if t.Elem == nil || u.Elem == nil {
			return t.Elem == u.Elem
		}
		return t.Elem.Equal(u.Elem)
	default:
		return true
	}
}
