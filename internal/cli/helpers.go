package cli

import (
	"flint/internal/typechecker"
	"flint/pkg/flint"
	"fmt"
	"os"
)

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func loadAndParse(filename string) ([]flint.Expr, *typechecker.TypeChecker) {
	src, err := os.ReadFile(filename)
	if err != nil {
		fatal(fmt.Sprintf("error reading %s: %v", filename, err))
	}

	tokens, err := flint.Lex(string(src), filename)
	if err != nil {
		fatal(err.Error())
	}

	prog, errs := flint.Parse(tokens)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(1)
	}

	tc := typechecker.New()
	for _, ex := range prog.Exprs {
		if _, err := tc.CheckExpr(ex); err != nil {
			fatal("Type error: " + err.Error())
		}
	}

	return prog.Exprs, tc
}
