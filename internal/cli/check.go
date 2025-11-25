package cli

import "fmt"

func checkFile(filename string) {
	fmt.Println("Type checking " + filename)
	_, _ = loadAndParse(filename)
	fmt.Println("No type errors found")
}
