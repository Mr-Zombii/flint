package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"flint/internal/parser"
)

type CodeGen struct {
	mod      *ir.Module
	strIndex int

	locals map[string]value.Value
	funcs  map[string]*ir.Func
}

func GenerateLLVM(prog *parser.Program) string {
	cg := &CodeGen{
		mod:    ir.NewModule(),
		locals: map[string]value.Value{},
		funcs:  map[string]*ir.Func{},
	}
	for _, e := range prog.Exprs {
		if fn, ok := e.(*parser.FuncDeclExpr); ok {
			name := fn.Name.Lexeme
			ret := cg.resolveType(fn.Ret)
			if name == "main" {
				ret = types.I32
			}
			params := []*ir.Param{}
			for _, p := range fn.Params {
				params = append(params, ir.NewParam(p.Name.Lexeme, cg.resolveType(p.Type)))
			}
			cg.funcs[name] = cg.mod.NewFunc(name, ret, params...)
		}
	}
	for _, e := range prog.Exprs {
		switch n := e.(type) {
		case *parser.FuncDeclExpr:
			cg.emitFunction(n)
		case *parser.IntLiteral, *parser.FloatLiteral, *parser.BoolLiteral,
			*parser.ByteLiteral, *parser.StringLiteral:
			cg.emitTopLiteral(n)
		default:
			panic("unsupported top-level expr")
		}
	}
	return cg.mod.String()
}

func (cg *CodeGen) emitFunction(fn *parser.FuncDeclExpr) {
	name := fn.Name.Lexeme
	mainFn := cg.funcs[name]
	cg.locals = map[string]value.Value{}
	entry := mainFn.NewBlock("entry")
	for _, param := range mainFn.Params {
		alloc := entry.NewAlloca(param.Type())
		entry.NewStore(param, alloc)
		cg.locals[param.Name()] = alloc
	}
	if fn.Body == nil {
		cg.emitDefaultReturn(entry, mainFn.Sig.RetType, name == "main")
		return
	}
	block := fn.Body.(*parser.BlockExpr)
	lastVal := cg.emitBlock(entry, block)
	if lastVal != nil {
		if b := parentBlockOfValue(lastVal); b != nil && b.Term == nil {
			b.NewRet(lastVal)
		} else if entry.Term == nil {
			entry.NewRet(lastVal)
		}
	} else if entry.Term == nil {
		cg.emitDefaultReturn(entry, mainFn.Sig.RetType, name == "main")
	}
}

func (cg *CodeGen) emitBlock(b *ir.Block, blk *parser.BlockExpr) value.Value {
	var last value.Value
	for _, e := range blk.Exprs {
		last = cg.emitExpr(b, e)
	}
	return last
}

func (cg *CodeGen) emitIf(b *ir.Block, i *parser.IfExpr) value.Value {
	cond := cg.emitExpr(b, i.Cond)
	parent := b.Parent
	thenBlock := parent.NewBlock("if.then")
	elseBlock := parent.NewBlock("if.else")
	mergeBlock := parent.NewBlock("if.merge")
	b.NewCondBr(cond, thenBlock, elseBlock)
	thenVal := cg.emitIfBody(thenBlock, i.Then)
	if thenBlock.Term == nil {
		thenBlock.NewBr(mergeBlock)
	}
	elseVal := cg.emitIfBody(elseBlock, i.Else)
	if elseVal == nil {
		panic("if expression requires an else branch")
	}
	if elseBlock.Term == nil {
		elseBlock.NewBr(mergeBlock)
	}
	phi := mergeBlock.NewPhi(
		&ir.Incoming{X: thenVal, Pred: thenBlock},
		&ir.Incoming{X: elseVal, Pred: elseBlock},
	)
	mergeBlock.NewRet(phi)
	return phi
}

func (cg *CodeGen) emitIfBody(b *ir.Block, body parser.Expr) value.Value {
	switch x := body.(type) {
	case *parser.BlockExpr:
		return cg.emitBlock(b, x)
	default:
		return cg.emitExpr(b, body)
	}
}

func (cg *CodeGen) emitString(v *parser.StringLiteral) value.Value {
	label := cg.newStrLabel()
	str := constant.NewCharArrayFromString(v.Value + "\x00")
	global := cg.mod.NewGlobalDef(label, str)
	global.Immutable = true
	global.Align = 1
	zero := constant.NewInt(types.I32, 0)
	return constant.NewGetElementPtr(
		str.Typ,
		global,
		zero,
		zero,
	)
}

func (cg *CodeGen) newStrLabel() string {
	name := ".str." + string(rune('a'+cg.strIndex))
	cg.strIndex++
	return name
}

func (cg *CodeGen) resolveType(t parser.Expr) types.Type {
	if t == nil {
		return types.Void
	}
	ty := t.(*parser.TypeExpr)
	switch ty.Name {
	case "Int":
		return types.I64
	case "Float":
		return types.Double
	case "Bool":
		return types.I1
	case "Byte":
		return types.I8
	case "String":
		return types.I8Ptr
	case "Nil":
		return types.Void
	default:
		return types.I64
	}
}

func (cg *CodeGen) emitCall(b *ir.Block, c *parser.CallExpr) value.Value {
	id, ok := c.Callee.(*parser.Identifier)
	if !ok {
		panic("only simple function calls supported")
	}
	fn := cg.funcs[id.Name]
	if fn == nil {
		panic("undefined function: " + id.Name)
	}
	var args []value.Value
	for _, arg := range c.Args {
		args = append(args, cg.emitExpr(b, arg))
	}
	return b.NewCall(fn, args...)
}

func (cg *CodeGen) emitDefaultReturn(b *ir.Block, ret types.Type, isMain bool) {
	if isMain {
		b.NewRet(constant.NewInt(types.I32, 0))
		return
	}
	switch t := ret.(type) {
	case *types.IntType:
		b.NewRet(constant.NewInt(t, 0))
	case *types.FloatType:
		b.NewRet(constant.NewFloat(t, 0))
	case *types.PointerType:
		b.NewRet(constant.NewNull(t))
	case *types.VoidType:
		b.NewRet(nil)
	default:
		panic("unsupported return type")
	}
}

func (cg *CodeGen) emitTopLiteral(e parser.Expr) {
	fn := cg.mod.NewFunc("main", types.I32)
	b := fn.NewBlock("entry")
	val := cg.emitExpr(b, e)
	if val.Type().Equal(types.I64) {
		val = b.NewTrunc(val, types.I32)
	}
	b.NewRet(val)
}
