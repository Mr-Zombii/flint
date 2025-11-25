package flint

import (
	"flint/internal/parser"
)

type (
	Parser  = parser.Parser
	Program = parser.Program
	Expr    = parser.Expr
)

func Parse(tokens []Token) (*Program, []string) {
	return parser.ParseProgram(tokens)
}
