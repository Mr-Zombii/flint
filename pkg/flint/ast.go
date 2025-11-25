package flint

import "flint/internal/parser"

func DumpAst(e Expr) string {
	return parser.DumpExpr(e)
}
