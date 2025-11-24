package typechecker

import "flint/internal/lexer"

type BinOpSig struct {
	Left  Type
	Right Type
	Out   Type
}

type UnaryOpSig struct {
	Arg Type
	Out Type
}

var binOps = map[lexer.TokenKind][]BinOpSig{
	lexer.Plus:    {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyInt}}},
	lexer.Minus:   {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyInt}}},
	lexer.Star:    {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyInt}}},
	lexer.Slash:   {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyInt}}},
	lexer.Percent: {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyInt}}},

	lexer.PlusDot:  {{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyFloat}}},
	lexer.MinusDot: {{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyFloat}}},
	lexer.StarDot:  {{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyFloat}}},
	lexer.SlashDot: {{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyFloat}}},

	lexer.Less:         {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyBool}}},
	lexer.LessEqual:    {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyBool}}},
	lexer.Greater:      {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyBool}}},
	lexer.GreaterEqual: {{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyBool}}},

	lexer.LessDot:         {{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyBool}}},
	lexer.LessEqualDot:    {{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyBool}}},
	lexer.GreaterDot:      {{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyBool}}},
	lexer.GreaterEqualDot: {{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyBool}}},

	lexer.AmperAmper: {{Type{kind: TyBool}, Type{kind: TyBool}, Type{kind: TyBool}}},
	lexer.VbarVbar:   {{Type{kind: TyBool}, Type{kind: TyBool}, Type{kind: TyBool}}},

	lexer.LtGt: {{Type{kind: TyString}, Type{kind: TyString}, Type{kind: TyString}}},

	lexer.EqualEqual: {
		{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyBool}},
		{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyBool}},
		{Type{kind: TyBool}, Type{kind: TyBool}, Type{kind: TyBool}},
		{Type{kind: TyString}, Type{kind: TyString}, Type{kind: TyBool}},
		{Type{kind: TyByte}, Type{kind: TyByte}, Type{kind: TyBool}},
	},
	lexer.NotEqual: {
		{Type{kind: TyInt}, Type{kind: TyInt}, Type{kind: TyBool}},
		{Type{kind: TyFloat}, Type{kind: TyFloat}, Type{kind: TyBool}},
		{Type{kind: TyBool}, Type{kind: TyBool}, Type{kind: TyBool}},
		{Type{kind: TyString}, Type{kind: TyString}, Type{kind: TyBool}},
		{Type{kind: TyByte}, Type{kind: TyByte}, Type{kind: TyBool}},
	},
}

var unaryOps = map[lexer.TokenKind]UnaryOpSig{
	lexer.Minus:    {Type{kind: TyInt}, Type{kind: TyInt}},
	lexer.MinusDot: {Type{kind: TyFloat}, Type{kind: TyFloat}},
	lexer.Bang:     {Type{kind: TyBool}, Type{kind: TyBool}},
}
