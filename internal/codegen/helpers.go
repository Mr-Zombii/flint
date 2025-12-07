package codegen

import (
	"flint/internal/typechecker"

	"github.com/llir/llvm/ir/types"
)

func (cg *CodeGen) newStrLabel() string {
	name := ".str." + string(rune('a'+cg.strIndex))
	cg.strIndex++
	return name
}

func (cg *CodeGen) platformIntType() *types.IntType {
	if typechecker.PlatformIntBits == 32 {
		return types.I32
	}
	return types.I64
}

func (cg *CodeGen) platformFloatType() *types.FloatType {
	if typechecker.PlatformIntBits == 32 {
		return types.Float
	}
	return types.Double
}
