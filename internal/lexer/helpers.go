package lexer

import "unicode"

func StripNumericSeparators(s string) string {
	out := []rune{}
	for _, r := range s {
		if r != '_' {
			out = append(out, r)
		}
	}
	return string(out)
}

func (k TokenKind) Precedence() int {
	if p, ok := precedence[k]; ok {
		return p
	}
	return 0
}

func LookupIdentifier(name string) TokenKind {
	if k, ok := KeywordMap[name]; ok {
		return k
	}
	return Identifier
}

func isIdentifierStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_' || r == '$'
}

func isIdentifierPart(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '$' || r == '\''
}
