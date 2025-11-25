package flint

import "flint/internal/lexer"

type (
	Token     = lexer.Token
	TokenKind = lexer.TokenKind
)

func Lex(source, filename string) ([]Token, error) {
	return lexer.Tokenize(source, filename)
}
