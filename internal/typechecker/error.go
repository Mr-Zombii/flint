package typechecker

import (
	"flint/internal/lexer"
	"fmt"
)

func (tc *TypeChecker) error(token lexer.Token, msg string) {
	tc.errors = append(tc.errors,
		fmt.Sprintf("[line %d:%d] Error: %s", token.Line, token.Column, msg))
}
