package typechecker

type Context int

const (
	TopLevel Context = iota
	FunctionBody
	Block
)
