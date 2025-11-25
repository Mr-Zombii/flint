package cli

import "fmt"

func compileFile(filename string) {
	fmt.Println("Compiling " + filename)
	_, _ = loadAndParse(filename)
	fmt.Println("Program compiled (no backend yet)")
}
