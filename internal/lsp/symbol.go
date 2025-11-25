package lsp

import (
	"regexp"
	"strings"
)

type SymbolKind int

const (
	FunctionSymbol SymbolKind = iota + 1
	VariableSymbol
)

type Symbol struct {
	Name string
	Kind SymbolKind
}

var symbols = map[string][]Symbol{}

var (
	fnRegex  = regexp.MustCompile(`^(?:pub\s+)?fn\s+([A-Za-z_][A-Za-z0-9_]*)\s*\(`)
	varRegex = regexp.MustCompile(`^(?:val|mut)\s+([A-Za-z_][A-Za-z0-9_]*)`)
)

func updateSymbols(uri, text string) {
	syms := []Symbol{}
	lines := strings.SplitSeq(text, "\n")

	for line := range lines {
		line = strings.TrimSpace(line)
		if fnMatch := fnRegex.FindStringSubmatch(line); fnMatch != nil {
			syms = append(syms, Symbol{Name: fnMatch[1], Kind: FunctionSymbol})
			continue
		}

		if varMatch := varRegex.FindStringSubmatch(line); varMatch != nil {
			syms = append(syms, Symbol{Name: varMatch[1], Kind: VariableSymbol})
		}
	}

	symbols[uri] = syms
}
